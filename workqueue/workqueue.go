package workqueue

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/public-awesome/stakewatcher/client"
	"github.com/public-awesome/stakewatcher/models"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gofrs/uuid"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

// Worker is the queue processor.
type Worker struct {
	q        <-chan int64
	cdc      *codec.Codec
	appCodec *std.Codec
	db       *sql.DB
	cp       *client.Proxy
}

// NewWorker returns an intialized worker
func NewWorker(cdc *codec.Codec, appCodec *std.Codec, queue <-chan int64, db *sql.DB, cp *client.Proxy) *Worker {
	return &Worker{
		q:        queue,
		cdc:      cdc,
		appCodec: appCodec,
		db:       db,
		cp:       cp,
	}
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
				// TODO: add retry
				continue
			}
			log.Info().Int64("height", blockHeight).Msg("block synced")
		}
	}
}

func (w *Worker) process(ctx context.Context, height int64) error {
	// skip block 1
	if height == 1 {
		return nil
	}
	exists, err := models.BlockExists(ctx, w.db, height)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	block, err := w.cp.Block(height)
	if err != nil {
		return fmt.Errorf("failed to fetch block: %w", err)
	}

	txs, err := w.cp.Txs(block)
	if err != nil {
		return fmt.Errorf("failed to fetch transactions for block: %w", err)
	}

	vals, err := w.cp.Validators(block.Block.LastCommit.GetHeight())
	if err != nil {
		return fmt.Errorf("failed to fetch validators for block: %w", err)
	}

	err = ExportBlockSignatures(ctx, block.Block.LastCommit, vals, w.db)
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

// ExportBlock exports a block by processing it.
func (w *Worker) ExportBlock(ctx context.Context, b *tmctypes.ResultBlock, txs []sdk.TxResponse, validators *tmctypes.ResultValidators, db *sql.DB) error {
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

	for _, t := range txs {
		stdTx, ok := t.Tx.(auth.StdTx)
		if !ok {
			log.Warn().Msgf("unsupported tx type: %T", t.Tx)
			continue
		}
		msgsBz, err := w.cdc.MarshalJSON(stdTx.GetMsgs())
		if err != nil {
			return fmt.Errorf("failed to JSON encode tx messages: %w", err)
		}
		evtsBz, err := w.cdc.MarshalJSON(t.Logs)
		if err != nil {
			return fmt.Errorf("failed to JSON encode tx logs: %w", err)
		}
		tx := &models.Transaction{
			Height:    b.Block.Height,
			Hash:      t.TxHash,
			GasWanted: int(t.GasWanted),
			GasUsed:   int(t.GasUsed),
			Messages:  msgsBz,
			Events:    evtsBz,
		}
		err = tx.Insert(ctx, db, boil.Infer())
		if err != nil {
			return fmt.Errorf("error inserting tx %w", err)
		}
		err = parseLogs(ctx, db, b.Block.Height, b.Block.Time, t.Logs)
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
func handlePost(ctx context.Context, db *sql.DB, attributes []sdk.Attribute, height int64, ts time.Time) error {
	attrs := parseAttributes(attributes)
	vendorID, err := strconv.Atoi(attrs["vendor_id"])
	if err != nil {
		return err
	}
	deposit, err := sdk.ParseCoin(attrs["deposit"])
	if err != nil {
		return err
	}
	endTime, err := time.Parse(time.RFC3339, attrs["curation_end_time"])
	if err != nil {
		return err
	}
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	p := &models.Post{
		ID:              id.String(),
		VendorID:        vendorID,
		Height:          height,
		PostID:          attrs["post_id"],
		Body:            attrs["body"],
		Creator:         attrs["creator"],
		RewardAddress:   attrs["reward_account"],
		DepositAmount:   deposit.Amount.Int64(),
		DepositDenom:    deposit.Denom,
		CurationEndTime: endTime,
		Timestamp:       ts,
	}
	return p.Insert(ctx, db, boil.Infer())
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
	deposit, err := sdk.ParseCoin(attrs["deposit"])
	if err != nil {
		return err
	}
	voteAmount, err := sdk.ParseCoin(attrs["vote_amount"])
	if err != nil {
		return err
	}
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	p := &models.Upvote{
		ID:            id.String(),
		VendorID:      vendorID,
		Height:        height,
		PostID:        attrs["post_id"],
		Creator:       attrs["curator"],
		RewardAddress: attrs["reward_account"],
		DepositAmount: deposit.Amount.Int64(),
		DepositDenom:  deposit.Denom,
		VoteAmount:    voteAmount.Amount.Int64(),
		VoteDenom:     voteAmount.Denom,
		VoteNumber:    voteNum,
		Timestamp:     ts,
	}
	return p.Insert(ctx, db, boil.Infer())
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
			}
		}
	}
	return nil
}

// sumGasTxs returns the total gas consumed by a set of transactions.
func sumGasTxs(txs []sdk.TxResponse) uint64 {
	var totalGas uint64
	for _, tx := range txs {
		totalGas += uint64(tx.GasUsed)
	}
	return totalGas
}

// ExportBlockSignatures ...
func ExportBlockSignatures(ctx context.Context, commit *tmtypes.Commit, validators *tmctypes.ResultValidators, db *sql.DB) error {
	for _, sig := range commit.Signatures {
		if sig.Signature == nil {
			continue
		}
		addr := sdk.ConsAddress(sig.ValidatorAddress).String()
		val := findValidatorByAddr(addr, validators)
		if val == nil {
			return fmt.Errorf("failed to find validator by address %s for block %d", addr, commit.GetHeight())
		}
		err := ExportValidator(ctx, val, db)
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
		Round:            commit.GetRound(),
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
func ExportValidator(ctx context.Context, val *tmtypes.Validator, db *sql.DB) error {
	address := sdk.ConsAddress(val.Address).String()
	consPubKey, err := sdk.Bech32ifyPubKey(sdk.Bech32PubKeyTypeConsPub, val.PubKey)

	if err != nil {
		return fmt.Errorf("failed to convert validator public key %w", err)
	}
	validator := &models.Validator{
		Address: address,
		PubKey:  consPubKey,
	}
	return validator.Upsert(ctx, db, false, []string{}, boil.Columns{}, boil.Infer())
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
