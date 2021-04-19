package workqueue

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	txtypes "github.com/cosmos/cosmos-sdk/types/tx"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/public-awesome/stargazer/client"
	"github.com/public-awesome/stargazer/models"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"

	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/gofrs/uuid"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

// Worker is the queue processor.
type Worker struct {
	q             <-chan int64
	cdc           codec.Marshaler
	amino         *codec.LegacyAmino
	db            *sql.DB
	cp            *client.Proxy
	genesisHeight int64
}

// NewWorker returns an intialized worker
func NewWorker(cdc codec.Marshaler, amino *codec.LegacyAmino, queue <-chan int64, db *sql.DB, cp *client.Proxy) *Worker {
	return &Worker{
		q:             queue,
		cdc:           cdc,
		amino:         amino,
		db:            db,
		cp:            cp,
		genesisHeight: 1,
	}
}

// WithGenesisHeight sets the first block height
func (w *Worker) WithGenesisHeight(height int64) {
	w.genesisHeight = height
}

// Start runs the listener that process blocks.
func (w *Worker) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Info().Msg("closing worker listener")
			return
		case blockHeight := <-w.q:
			err := w.process(ctx, blockHeight)
			if err != nil {
				log.Error().Err(err).Int64("height", blockHeight).Msg("error processing block")
				continue
			}
			log.Info().Int64("height", blockHeight).Msg("block synced")
		}
	}
}

func (w *Worker) processGenesisBlock(ctx context.Context, height int64) error {
	exists, err := models.BlockExists(ctx, w.db, height)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	block, err := w.cp.Block(ctx, height)
	if err != nil {
		return fmt.Errorf("failed to fetch block: %w", err)
	}

	vals, err := w.cp.Validators(ctx, block.Block.LastCommit.GetHeight())
	if err != nil {
		return fmt.Errorf("failed to fetch validators for block: %w", err)
	}
	appValidators, err := w.appValidators(ctx, block.Block.LastCommit.GetHeight())
	if err != nil {
		log.Err(err).Msg("error retrieving app validators")
	}
	err = ExportBlockSignatures(ctx, block.Block.LastCommit, vals, appValidators, w.db)
	if err != nil {
		return fmt.Errorf("failed to export block signatures %w", err)
	}
	sl, err := models.FindSyncLog(ctx, w.db, height)
	if err != nil {
		return fmt.Errorf("error finding sync log: %w", err)
	}
	sl.Processed = true
	sl.SyncedAt = null.NewTime(time.Now(), true)

	_, err = sl.Update(ctx, w.db, boil.Infer())
	return err
}
func (w *Worker) appValidators(ctx context.Context, height int64) (map[string]stakingtypes.Validator, error) {
	appValidators, err := w.cp.AppValidators(ctx, height)
	validators := make(map[string]stakingtypes.Validator)
	if err != nil {
		return validators, fmt.Errorf("failed to fetch app validators for block: %w", err)
	}

	for _, validator := range appValidators {
		v := validator
		err = w.cdc.UnpackAny(v.ConsensusPubkey, new(cryptotypes.PubKey))
		if err != nil {
			log.Err(err).Msg("failed to unpack val consensus pub key")
			continue
		}
		consAddr, err := v.GetConsAddr()
		if err != nil {
			log.Err(err).Msg("failed retrieving cons addr")
			continue
		}
		validators[consAddr.String()] = v
	}

	return validators, nil
}
func (w *Worker) process(ctx context.Context, height int64) error {
	// skip genesis height
	if height == w.genesisHeight {
		return w.processGenesisBlock(ctx, height)
	}
	exists, err := models.BlockExists(ctx, w.db, height)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	block, err := w.cp.Block(ctx, height)
	if err != nil {
		return fmt.Errorf("failed to fetch block: %w", err)
	}

	blockResults, err := w.cp.BlockResults(ctx, height)
	if err != nil {
		return fmt.Errorf("failed to fetch block result: %w", err)
	}
	err = w.processBlockEvents(ctx, blockResults, height)
	if err != nil {
		return fmt.Errorf("failed to process block events: %w", err)
	}

	txs, err := w.cp.Txs(block)
	if err != nil {
		return fmt.Errorf("failed to fetch transactions for block: %w", err)
	}

	vals, err := w.cp.Validators(ctx, block.Block.LastCommit.GetHeight())
	if err != nil {
		return fmt.Errorf("failed to fetch validators for block: %w", err)
	}
	appValidators, err := w.appValidators(ctx, block.Block.LastCommit.GetHeight())
	if err != nil {
		log.Err(err).Msg("error retrieving app validators")
	}
	err = ExportBlockSignatures(ctx, block.Block.LastCommit, vals, appValidators, w.db)
	if err != nil {
		return fmt.Errorf("failed to export block signatures %w", err)
	}
	err = w.ExportBlock(ctx, block, txs, vals, w.db)
	if err != nil {
		return fmt.Errorf("failed to export block: %w", err)
	}

	sl, err := models.FindSyncLog(ctx, w.db, height)
	if err != nil {
		return fmt.Errorf("error finding sync log: %w", err)
	}
	sl.Processed = true
	sl.SyncedAt = null.NewTime(time.Now(), true)

	_, err = sl.Update(ctx, w.db, boil.Infer())
	return err
}

func (w *Worker) processBlockEvents(ctx context.Context, br *tmctypes.ResultBlockResults, height int64) error {
	for _, evt := range br.EndBlockEvents {
		switch evt.Type {
		case "curation_complete":
			err := handleCurationComplete(ctx, w.db, evt.Attributes, height)
			if err != nil {
				return err
			}
		case "protocol_reward":
			err := handleProtocolReward(ctx, w.db, evt.Attributes, height)
			if err != nil {
				return err
			}
		}
	}
	for _, evt := range br.BeginBlockEvents {
		switch evt.Type {
		case slashingtypes.EventTypeLiveness:
			err := handleLiveness(ctx, w.db, evt.Attributes, height)
			if err != nil {
				return err
			}
		case slashingtypes.EventTypeSlash:
			err := handleSlash(ctx, w.db, evt.Attributes, height)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (w *Worker) MarshalMsgs(msgs []sdk.Msg) ([]byte, error) {
	resp := bytes.NewBuffer(make([]byte, 0))
	resp.Write([]byte("["))
	total := len(msgs)
	for i, msg := range msgs {
		msgBz, err := w.cdc.MarshalJSON(msg)
		if err != nil {
			return nil, err
		}
		_, err = resp.Write(msgBz)
		if err != nil {
			return nil, err
		}
		if i != total-1 {
			resp.Write([]byte(","))
		}
	}
	resp.Write([]byte("]"))
	return resp.Bytes(), nil
}

// ExportBlock exports a block by processing it.
func (w *Worker) ExportBlock(ctx context.Context, b *tmctypes.ResultBlock, txs []*sdk.TxResponse, validators *tmctypes.ResultValidators, db *sql.DB) error {
	totalGas := sumGasTxs(txs)
	signatures := len(b.Block.LastCommit.Signatures)

	proposerAddress := sdk.ConsAddress(b.Block.ProposerAddress).String()
	val := findValidatorByAddr(proposerAddress, validators)
	if val == nil {
		return fmt.Errorf("failed to find validator by address %s for block %d", proposerAddress, b.Block.Height)
	}
	block := &models.Block{
		Height:          b.Block.Height,
		NumTXS:          len(txs),
		TotalGas:        int64(totalGas),
		Signatures:      signatures,
		ProposerAddress: proposerAddress,
		Hash:            b.BlockID.Hash.String(),
		BlockTimestamp:  b.Block.Time,
	}

	err := block.Insert(ctx, db, boil.Infer())
	if err != nil {
		return fmt.Errorf("failed to insert block %w", err)
	}

	for _, txResponse := range txs {
		var tx txtypes.Tx
		err := w.cdc.UnmarshalBinaryBare(txResponse.Tx.Value, &tx)
		if err != nil {
			return fmt.Errorf("failed to JSON encode tx messages: %w", err)
		}

		msgs := tx.GetMsgs()
		msgsBz, err := w.MarshalMsgs(msgs)
		if err != nil {
			return fmt.Errorf("failed to JSON encode tx messages: %w", err)
		}
		evtsBz, err := json.Marshal(txResponse.Logs)
		if err != nil {
			return fmt.Errorf("failed to JSON encode tx logs: %w", err)
		}
		t := &models.Transaction{
			Height:    b.Block.Height,
			Hash:      txResponse.TxHash,
			GasWanted: int(txResponse.GasWanted),
			GasUsed:   int(txResponse.GasUsed),
			Messages:  msgsBz,
			Events:    evtsBz,
		}
		err = t.Insert(ctx, db, boil.Infer())
		if err != nil {
			return fmt.Errorf("error inserting tx %w", err)
		}
		err = parseLogs(ctx, db, b.Block.Height, b.Block.Time, txResponse.Logs)
		if err != nil {
			return fmt.Errorf("error parsing logs %w", err)
		}
	}
	return nil
}

func parseAttributes(attributes []sdk.Attribute) map[string]string {
	attrs := make(map[string]string)
	for _, a := range attributes {
		attrs[a.Key] = a.Value
	}
	return attrs
}

func parseEventAttributes(attributes []abcitypes.EventAttribute) map[string]string {
	attrs := make(map[string]string)
	for _, a := range attributes {
		attrs[string(a.Key)] = string(a.Value)
	}
	return attrs
}

func handlePost(ctx context.Context, db *sql.DB, attributes []sdk.Attribute, height int64, ts time.Time) error {
	attrs := parseAttributes(attributes)
	vendorID, err := strconv.Atoi(attrs["vendor_id"])
	if err != nil {
		return err
	}

	locked := attrs["locked"] == "true"

	endTime, err := time.Parse(time.RFC3339, attrs["curation_end_time"])
	if err != nil {
		return err
	}
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	p := &models.Post{
		ID:               id.String(),
		ChainID:          attrs["chain_id"],
		Height:           height,
		VendorID:         vendorID,
		PostID:           attrs["post_id"],
		Body:             attrs["body"],
		BodyHash:         attrs["body_hash"],
		ContractAddress:  attrs["contract_address"],
		Creator:          attrs["creator"],
		Locked:           locked,
		Metadata:         attrs["metadata"],
		Owner:            attrs["owner"],
		RewardAddress:    attrs["reward_account"],
		ParentID:         attrs["parent_id"],
		TotalVotes:       0,
		TotalVotesAmount: 0,
		TotalVoterCount:  0,
		TotalVotesDenom:  attrs["vote_denom"],
		CurationEndTime:  endTime,
		Timestamp:        ts,
	}
	return p.Insert(ctx, db, boil.Infer())
}

func handleCurationComplete(ctx context.Context, db *sql.DB, attributes []abcitypes.EventAttribute, height int64) error {
	attrs := parseEventAttributes(attributes)
	if attrs["vendor_id"] == "" || attrs["reward_amount"] == "" || attrs["post_id"] == "" {
		return nil
	}
	vendorID, err := strconv.Atoi(attrs["vendor_id"])
	if err != nil {
		return fmt.Errorf("curation_complete: error parsing vendor id %s %w", attrs["vendor_id"], err)
	}
	postID := attrs["post_id"]
	amount, err := sdk.ParseCoinNormalized(attrs["reward_amount"])
	if err != nil {
		return fmt.Errorf("curation_complete: error parsing reward_amount %s %w", attrs["reward_amount"], err)
	}

	p, err := models.Posts(
		models.PostWhere.VendorID.EQ(vendorID),
		models.PostWhere.PostID.EQ(postID),
	).One(ctx, db)
	if err != nil {
		return fmt.Errorf("curation_complete: error retrieving post %s %w", postID, err)
	}

	p.TotalUpvoteRewardAmount = amount.Amount.Int64()
	p.TotalUpvoteRewardDenom = amount.Denom
	_, err = p.Update(
		ctx,
		db,
		boil.Whitelist(
			models.PostColumns.TotalUpvoteRewardAmount,
			models.PostColumns.TotalUpvoteRewardDenom,
			models.PostColumns.UpdatedAt),
	)

	return err
}

func handleProtocolReward(ctx context.Context, db *sql.DB, attributes []abcitypes.EventAttribute, height int64) error {
	attrs := parseEventAttributes(attributes)
	if attrs["vendor_id"] == "" || attrs["reward_account"] == "" || attrs["post_id"] == "" || attrs["reward_amount"] == "" {
		return nil
	}
	vendorID, err := strconv.Atoi(attrs["vendor_id"])
	if err != nil {
		return fmt.Errorf("protocol_reward: error parsing vendor id %s %w", attrs["vendor_id"], err)
	}
	postID := attrs["post_id"]
	rewardAddress := attrs["reward_account"]
	amount, err := sdk.ParseCoinNormalized(attrs["reward_amount"])
	if err != nil {
		return fmt.Errorf("protocol_reward: error parsing reward_amount %s %w", attrs["reward_amount"], err)
	}

	ur := &models.UpvoteReward{
		Height:        height,
		VendorID:      vendorID,
		PostID:        postID,
		RewardAddress: rewardAddress,
		RewardAmount:  amount.Amount.Int64(),
		RewardDenom:   amount.Denom,
	}

	err = ur.Insert(ctx, db, boil.Infer())
	if err != nil {
		return fmt.Errorf("protocol_reward: error inserting upvote reward (%s) %w", postID, err)
	}
	return nil
}

func handleLiveness(ctx context.Context, db *sql.DB, attributes []abcitypes.EventAttribute, height int64) error {
	attrs := parseEventAttributes(attributes)
	if attrs[slashingtypes.AttributeKeyAddress] == "" || attrs[slashingtypes.AttributeKeyMissedBlocks] == "" {
		return nil
	}
	addr := attrs[slashingtypes.AttributeKeyAddress]

	blocksCounter, err := strconv.Atoi(attrs[slashingtypes.AttributeKeyMissedBlocks])
	if err != nil {
		log.Warn().Msg(fmt.Sprintf("can not parse block counter [%s] [%s]", addr, slashingtypes.AttributeKeyMissedBlocks))
	}
	se := &models.SlashingEvent{
		Height:           height,
		ValidatorAddress: addr,
		EventType:        slashingtypes.EventTypeLiveness,
		Counter:          int64(blocksCounter),
	}
	err = se.Insert(ctx, db, boil.Infer())
	if err != nil {
		return fmt.Errorf("slashing_event: error inserting liveness %s %w", addr, err)
	}
	return nil
}

func handleSlash(ctx context.Context, db *sql.DB, attributes []abcitypes.EventAttribute, height int64) error {
	attrs := parseEventAttributes(attributes)
	if attrs[slashingtypes.AttributeKeyAddress] == "" {
		return nil
	}
	addr := attrs[slashingtypes.AttributeKeyAddress]
	reason := attrs[slashingtypes.AttributeKeyReason]

	se := &models.SlashingEvent{
		Height:           height,
		ValidatorAddress: addr,
		EventType:        slashingtypes.EventTypeSlash,
		Counter:          0,
		Reason:           reason,
	}
	err := se.Insert(ctx, db, boil.Infer())
	if err != nil {
		return fmt.Errorf("slashing_event: error inserting liveness %s %w", addr, err)
	}
	return nil
}

func handleUpvote(ctx context.Context, db *sql.DB, attributes []sdk.Attribute, height int64, ts time.Time) error {
	attrs := parseAttributes(attributes)
	vendorID, err := strconv.Atoi(attrs["vendor_id"])
	if err != nil {
		return err
	}
	voteNum, err := strconv.Atoi(attrs["vote_number"])
	if err != nil {
		return err
	}
	voteAmount, err := sdk.ParseCoinNormalized(attrs["vote_amount"])
	if err != nil {
		return err
	}
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	updatedVoteAmount := voteAmount.Amount.Int64()
	updatedVoteNumber := voteNum
	p := &models.Upvote{
		ID:            id.String(),
		VendorID:      vendorID,
		Height:        height,
		PostID:        attrs["post_id"],
		Creator:       attrs["curator"],
		RewardAddress: attrs["reward_account"],
		VoteAmount:    updatedVoteAmount,
		VoteDenom:     voteAmount.Denom,
		VoteNumber:    updatedVoteNumber,
		Timestamp:     ts,
	}
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	upvote, err := models.Upvotes(
		models.UpvoteWhere.VendorID.EQ(vendorID),
		models.UpvoteWhere.PostID.EQ(p.PostID),
		models.UpvoteWhere.Creator.EQ(p.Creator),
	).One(ctx, tx)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if err == sql.ErrNoRows {
		err = p.Insert(ctx, tx, boil.Infer())
		if err != nil {
			_ = tx.Rollback()
			return err
		}

		updatePostQuery := `UPDATE posts SET total_votes = total_votes + $1,
	        total_votes_amount = total_votes_amount + $2,
	        total_voter_count = total_voter_count + 1,
	        updated_at = $3
            WHERE vendor_id=$4 and post_id=$5`
		_, err = queries.
			Raw(updatePostQuery, updatedVoteNumber, updatedVoteAmount, time.Now().UTC(), p.VendorID, p.PostID).
			ExecContext(ctx, tx)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		return tx.Commit()
	} else {
		updatedVoteAmount = upvote.VoteAmount + p.VoteAmount
		updatedVoteNumber = upvote.VoteNumber + p.VoteNumber
		upvote.VoteAmount = updatedVoteAmount
		upvote.VoteNumber = updatedVoteNumber
		_, err = upvote.Update(ctx, tx, boil.Infer())
		if err != nil {
			_ = tx.Rollback()
			return err
		}

		updatePostQuery := `UPDATE posts SET total_votes = total_votes + $1,
	        total_votes_amount = total_votes_amount + $2,
	        updated_at = $3
            WHERE vendor_id=$4 and post_id=$5`
		_, err = queries.
			Raw(updatePostQuery, p.VoteNumber, p.VoteAmount, time.Now().UTC(), p.VendorID, p.PostID).
			ExecContext(ctx, tx)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		return tx.Commit()
	}

}

// This works for both stake and unstake commands since they both send only the total amount, which is upserted.
func handleStake(ctx context.Context, db *sql.DB, attributes []sdk.Attribute, height int64, ts time.Time) error {
	attrs := parseAttributes(attributes)
	vendorID, err := strconv.Atoi(attrs["vendor_id"])
	if err != nil {
		return err
	}
	amount, err := strconv.ParseInt(attrs["amount"], 10, 64)
	if err != nil {
		return err
	}
	postID := attrs["post_id"]
	delegator := attrs["delegator"]
	validator := attrs["validator"]

	if amount == 0 {
		s, err := models.Stakes(
			models.StakeWhere.VendorID.EQ(vendorID),
			models.StakeWhere.PostID.EQ(postID),
			models.StakeWhere.Delegator.EQ(delegator),
		).One(ctx, db)
		if err != nil {
			return err
		}

		rowAff, err := s.Delete(ctx, db)
		if (rowAff == 0) || (err != nil) {
			return err
		}

		return nil
	}

	model := &models.Stake{
		Height:    height,
		VendorID:  vendorID,
		PostID:    postID,
		Delegator: delegator,
		Validator: validator,
		Amount:    amount,
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = model.Upsert(
		ctx,
		tx,
		true,
		[]string{"vendor_id", "post_id", "delegator", "validator"},
		boil.Whitelist("amount", "updated_at"),
		boil.Infer(),
	)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	post, err := models.Posts(
		models.PostWhere.VendorID.EQ(vendorID),
		models.PostWhere.PostID.EQ(postID),
	).One(ctx, db)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	post.TotalStakedAmount = amount
	_, err = post.Update(
		ctx,
		tx,
		boil.Whitelist(
			models.PostColumns.TotalStakedAmount,
			models.PostColumns.UpdatedAt),
	)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

func handleBuyCreatorCoin(ctx context.Context, db *sql.DB, attributes []sdk.Attribute, height int64, ts time.Time) error {
	attrs := parseAttributes(attributes)

	amount, err := strconv.ParseInt(attrs["amount"], 10, 64)
	if err != nil {
		return err
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	buyer := attrs["buyer"]
	creator := attrs["creator"]
	username := attrs["username"]
	validator := attrs["validator"]

	cc, err := models.SocialGraphs(
		models.SocialGraphWhere.BuyerAddress.EQ(buyer),
		models.SocialGraphWhere.CreatorAddress.EQ(creator),
		models.SocialGraphWhere.Username.EQ(username),
		models.SocialGraphWhere.ValidatorAddress.EQ(validator),
	).One(ctx, db)

	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if err == sql.ErrNoRows {
		cc := &models.SocialGraph{
			Amount:           amount,
			BuyerAddress:     buyer,
			CreatorAddress:   creator,
			Height:           height,
			Username:         username,
			ValidatorAddress: validator,
		}

		err = cc.Insert(ctx, db, boil.Infer())
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	} else {
		cc.Amount = cc.Amount + amount
		_, err = cc.Update(ctx, db, boil.Infer())
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func parseLogs(ctx context.Context, db *sql.DB, height int64, ts time.Time, logs sdk.ABCIMessageLogs) error {
	for _, l := range logs {
		for _, evt := range l.Events {
			switch evt.Type {
			case "post":
				err := handlePost(ctx, db, evt.Attributes, height, ts)
				if err != nil {
					return err
				}
			case "upvote":
				err := handleUpvote(ctx, db, evt.Attributes, height, ts)
				if err != nil {
					return err
				}
			case "stake":
				err := handleStake(ctx, db, evt.Attributes, height, ts)
				if err != nil {
					return err
				}
			case "unstake":
				err := handleStake(ctx, db, evt.Attributes, height, ts)
				if err != nil {
					return err
				}
			case "buy_creator_coin":
				err := handleBuyCreatorCoin(ctx, db, evt.Attributes, height, ts)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// sumGasTxs returns the total gas consumed by a set of transactions.
func sumGasTxs(txs []*sdk.TxResponse) uint64 {
	var totalGas uint64
	for _, tx := range txs {
		totalGas += uint64(tx.GasUsed)
	}
	return totalGas
}

// ExportBlockSignatures ...
func ExportBlockSignatures(ctx context.Context, commit *tmtypes.Commit, validators *tmctypes.ResultValidators, appValidators map[string]stakingtypes.Validator, db *sql.DB) error {
	for _, sig := range commit.Signatures {
		if sig.Signature == nil {
			continue
		}
		addr := sdk.ConsAddress(sig.ValidatorAddress).String()
		val := findValidatorByAddr(addr, validators)
		if val == nil {
			return fmt.Errorf("failed to find validator by address %s for block %d", addr, commit.GetHeight())
		}
		err := ExportValidator(ctx, val, appValidators, db)
		if err != nil {
			return fmt.Errorf("failed to export validator %w", err)
		}

		err = SetBlockSignature(ctx, commit, sig, val.VotingPower, val.ProposerPriority, db)
		if err != nil {
			return err
		}
	}
	return nil
}

// SetBlockSignature stores a block prevote
func SetBlockSignature(ctx context.Context, commit *tmtypes.Commit, sig tmtypes.CommitSig, vp, pp int64, db *sql.DB) error {
	s := &models.BlockSignature{
		Height:           commit.GetHeight(),
		Round:            int(commit.GetRound()),
		ValidatorAddress: sdk.ConsAddress(sig.ValidatorAddress).String(),
		VotingPower:      int(vp),
		ProposerPriority: int(pp),
		Flag:             int(sig.BlockIDFlag),
		Timestamp:        sig.Timestamp,
		Hash:             sig.BlockID(commit.BlockID).Hash.String(),
	}
	return s.Upsert(ctx, db, false, nil, boil.Columns{}, boil.Infer())
}

// ExportValidator exports validator
func ExportValidator(ctx context.Context, val *tmtypes.Validator, appValidators map[string]stakingtypes.Validator, db *sql.DB) error {
	address := sdk.ConsAddress(val.Address).String()
	pk, err := cryptocodec.FromTmPubKeyInterface(val.PubKey)
	if err != nil {
		return err
	}
	consPubKey, err := sdk.Bech32ifyPubKey(sdk.Bech32PubKeyTypeConsPub, pk)

	if err != nil {
		return fmt.Errorf("failed to convert validator public key %w", err)
	}
	validator := &models.Validator{
		Address: address,
		PubKey:  consPubKey,
	}
	appValidator, ok := appValidators[address]
	if ok {
		validator.Moniker = appValidator.Description.Moniker
		validator.OperatorAddress = appValidator.OperatorAddress
	}
	return validator.Upsert(ctx, db, true, []string{}, boil.Whitelist("operator_address", "moniker"), boil.Infer())
}

// findValidatorByAddr finds a validator by address given a set of
// Tendermint validators for a particular block. If no validator is found, nil
// is returned.
func findValidatorByAddr(address string, vals *tmctypes.ResultValidators) *tmtypes.Validator {
	for _, val := range vals.Validators {
		if strings.EqualFold(address, sdk.ConsAddress(val.Address).String()) {
			return val
		}
	}
	return nil
}
