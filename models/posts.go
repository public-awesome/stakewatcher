// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Post is an object representing the database table.
type Post struct {
	ID                      string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Height                  int64     `boil:"height" json:"height" toml:"height" yaml:"height"`
	VendorID                int       `boil:"vendor_id" json:"vendorID" toml:"vendorID" yaml:"vendorID"`
	PostID                  string    `boil:"post_id" json:"postID" toml:"postID" yaml:"postID"`
	Creator                 string    `boil:"creator" json:"creator" toml:"creator" yaml:"creator"`
	RewardAddress           string    `boil:"reward_address" json:"rewardAddress" toml:"rewardAddress" yaml:"rewardAddress"`
	Timestamp               time.Time `boil:"timestamp" json:"timestamp" toml:"timestamp" yaml:"timestamp"`
	CurationEndTime         time.Time `boil:"curation_end_time" json:"curationEndTime" toml:"curationEndTime" yaml:"curationEndTime"`
	Body                    string    `boil:"body" json:"body" toml:"body" yaml:"body"`
	TotalVotes              int       `boil:"total_votes" json:"totalVotes" toml:"totalVotes" yaml:"totalVotes"`
	TotalVotesAmount        int64     `boil:"total_votes_amount" json:"totalVotesAmount" toml:"totalVotesAmount" yaml:"totalVotesAmount"`
	TotalVotesDenom         string    `boil:"total_votes_denom" json:"totalVotesDenom" toml:"totalVotesDenom" yaml:"totalVotesDenom"`
	CreatedAt               time.Time `boil:"created_at" json:"createdAt" toml:"createdAt" yaml:"createdAt"`
	UpdatedAt               time.Time `boil:"updated_at" json:"updatedAt" toml:"updatedAt" yaml:"updatedAt"`
	DeletedAt               null.Time `boil:"deleted_at" json:"deletedAt,omitempty" toml:"deletedAt" yaml:"deletedAt,omitempty"`
	TotalVoterCount         int       `boil:"total_voter_count" json:"totalVoterCount" toml:"totalVoterCount" yaml:"totalVoterCount"`
	TotalUpvoteRewardAmount int64     `boil:"total_upvote_reward_amount" json:"totalUpvoteRewardAmount" toml:"totalUpvoteRewardAmount" yaml:"totalUpvoteRewardAmount"`
	TotalUpvoteRewardDenom  string    `boil:"total_upvote_reward_denom" json:"totalUpvoteRewardDenom" toml:"totalUpvoteRewardDenom" yaml:"totalUpvoteRewardDenom"`
	TotalStakedAmount       int64     `boil:"total_staked_amount" json:"totalStakedAmount" toml:"totalStakedAmount" yaml:"totalStakedAmount"`
	BodyHash                string    `boil:"body_hash" json:"bodyHash" toml:"bodyHash" yaml:"bodyHash"`
	ChainID                 string    `boil:"chain_id" json:"chainID" toml:"chainID" yaml:"chainID"`
	Owner                   string    `boil:"owner" json:"owner" toml:"owner" yaml:"owner"`
	ContractAddress         string    `boil:"contract_address" json:"contractAddress" toml:"contractAddress" yaml:"contractAddress"`
	Metadata                string    `boil:"metadata" json:"metadata" toml:"metadata" yaml:"metadata"`
	Locked                  bool      `boil:"locked" json:"locked" toml:"locked" yaml:"locked"`
	ParentID                string    `boil:"parent_id" json:"parentID" toml:"parentID" yaml:"parentID"`
	TX                      string    `boil:"tx" json:"tx" toml:"tx" yaml:"tx"`

	R *postR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L postL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PostColumns = struct {
	ID                      string
	Height                  string
	VendorID                string
	PostID                  string
	Creator                 string
	RewardAddress           string
	Timestamp               string
	CurationEndTime         string
	Body                    string
	TotalVotes              string
	TotalVotesAmount        string
	TotalVotesDenom         string
	CreatedAt               string
	UpdatedAt               string
	DeletedAt               string
	TotalVoterCount         string
	TotalUpvoteRewardAmount string
	TotalUpvoteRewardDenom  string
	TotalStakedAmount       string
	BodyHash                string
	ChainID                 string
	Owner                   string
	ContractAddress         string
	Metadata                string
	Locked                  string
	ParentID                string
	TX                      string
}{
	ID:                      "id",
	Height:                  "height",
	VendorID:                "vendor_id",
	PostID:                  "post_id",
	Creator:                 "creator",
	RewardAddress:           "reward_address",
	Timestamp:               "timestamp",
	CurationEndTime:         "curation_end_time",
	Body:                    "body",
	TotalVotes:              "total_votes",
	TotalVotesAmount:        "total_votes_amount",
	TotalVotesDenom:         "total_votes_denom",
	CreatedAt:               "created_at",
	UpdatedAt:               "updated_at",
	DeletedAt:               "deleted_at",
	TotalVoterCount:         "total_voter_count",
	TotalUpvoteRewardAmount: "total_upvote_reward_amount",
	TotalUpvoteRewardDenom:  "total_upvote_reward_denom",
	TotalStakedAmount:       "total_staked_amount",
	BodyHash:                "body_hash",
	ChainID:                 "chain_id",
	Owner:                   "owner",
	ContractAddress:         "contract_address",
	Metadata:                "metadata",
	Locked:                  "locked",
	ParentID:                "parent_id",
	TX:                      "tx",
}

// Generated where

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var PostWhere = struct {
	ID                      whereHelperstring
	Height                  whereHelperint64
	VendorID                whereHelperint
	PostID                  whereHelperstring
	Creator                 whereHelperstring
	RewardAddress           whereHelperstring
	Timestamp               whereHelpertime_Time
	CurationEndTime         whereHelpertime_Time
	Body                    whereHelperstring
	TotalVotes              whereHelperint
	TotalVotesAmount        whereHelperint64
	TotalVotesDenom         whereHelperstring
	CreatedAt               whereHelpertime_Time
	UpdatedAt               whereHelpertime_Time
	DeletedAt               whereHelpernull_Time
	TotalVoterCount         whereHelperint
	TotalUpvoteRewardAmount whereHelperint64
	TotalUpvoteRewardDenom  whereHelperstring
	TotalStakedAmount       whereHelperint64
	BodyHash                whereHelperstring
	ChainID                 whereHelperstring
	Owner                   whereHelperstring
	ContractAddress         whereHelperstring
	Metadata                whereHelperstring
	Locked                  whereHelperbool
	ParentID                whereHelperstring
	TX                      whereHelperstring
}{
	ID:                      whereHelperstring{field: "\"posts\".\"id\""},
	Height:                  whereHelperint64{field: "\"posts\".\"height\""},
	VendorID:                whereHelperint{field: "\"posts\".\"vendor_id\""},
	PostID:                  whereHelperstring{field: "\"posts\".\"post_id\""},
	Creator:                 whereHelperstring{field: "\"posts\".\"creator\""},
	RewardAddress:           whereHelperstring{field: "\"posts\".\"reward_address\""},
	Timestamp:               whereHelpertime_Time{field: "\"posts\".\"timestamp\""},
	CurationEndTime:         whereHelpertime_Time{field: "\"posts\".\"curation_end_time\""},
	Body:                    whereHelperstring{field: "\"posts\".\"body\""},
	TotalVotes:              whereHelperint{field: "\"posts\".\"total_votes\""},
	TotalVotesAmount:        whereHelperint64{field: "\"posts\".\"total_votes_amount\""},
	TotalVotesDenom:         whereHelperstring{field: "\"posts\".\"total_votes_denom\""},
	CreatedAt:               whereHelpertime_Time{field: "\"posts\".\"created_at\""},
	UpdatedAt:               whereHelpertime_Time{field: "\"posts\".\"updated_at\""},
	DeletedAt:               whereHelpernull_Time{field: "\"posts\".\"deleted_at\""},
	TotalVoterCount:         whereHelperint{field: "\"posts\".\"total_voter_count\""},
	TotalUpvoteRewardAmount: whereHelperint64{field: "\"posts\".\"total_upvote_reward_amount\""},
	TotalUpvoteRewardDenom:  whereHelperstring{field: "\"posts\".\"total_upvote_reward_denom\""},
	TotalStakedAmount:       whereHelperint64{field: "\"posts\".\"total_staked_amount\""},
	BodyHash:                whereHelperstring{field: "\"posts\".\"body_hash\""},
	ChainID:                 whereHelperstring{field: "\"posts\".\"chain_id\""},
	Owner:                   whereHelperstring{field: "\"posts\".\"owner\""},
	ContractAddress:         whereHelperstring{field: "\"posts\".\"contract_address\""},
	Metadata:                whereHelperstring{field: "\"posts\".\"metadata\""},
	Locked:                  whereHelperbool{field: "\"posts\".\"locked\""},
	ParentID:                whereHelperstring{field: "\"posts\".\"parent_id\""},
	TX:                      whereHelperstring{field: "\"posts\".\"tx\""},
}

// PostRels is where relationship names are stored.
var PostRels = struct {
}{}

// postR is where relationships are stored.
type postR struct {
}

// NewStruct creates a new relationship struct
func (*postR) NewStruct() *postR {
	return &postR{}
}

// postL is where Load methods for each relationship are stored.
type postL struct{}

var (
	postAllColumns            = []string{"id", "height", "vendor_id", "post_id", "creator", "reward_address", "timestamp", "curation_end_time", "body", "total_votes", "total_votes_amount", "total_votes_denom", "created_at", "updated_at", "deleted_at", "total_voter_count", "total_upvote_reward_amount", "total_upvote_reward_denom", "total_staked_amount", "body_hash", "chain_id", "owner", "contract_address", "metadata", "locked", "parent_id", "tx"}
	postColumnsWithoutDefault = []string{"id", "height", "vendor_id", "post_id", "creator", "reward_address", "timestamp", "curation_end_time", "body", "total_votes_denom", "deleted_at"}
	postColumnsWithDefault    = []string{"total_votes", "total_votes_amount", "created_at", "updated_at", "total_voter_count", "total_upvote_reward_amount", "total_upvote_reward_denom", "total_staked_amount", "body_hash", "chain_id", "owner", "contract_address", "metadata", "locked", "parent_id", "tx"}
	postPrimaryKeyColumns     = []string{"id"}
)

type (
	// PostSlice is an alias for a slice of pointers to Post.
	// This should generally be used opposed to []Post.
	PostSlice []*Post

	postQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	postType                 = reflect.TypeOf(&Post{})
	postMapping              = queries.MakeStructMapping(postType)
	postPrimaryKeyMapping, _ = queries.BindMapping(postType, postMapping, postPrimaryKeyColumns)
	postInsertCacheMut       sync.RWMutex
	postInsertCache          = make(map[string]insertCache)
	postUpdateCacheMut       sync.RWMutex
	postUpdateCache          = make(map[string]updateCache)
	postUpsertCacheMut       sync.RWMutex
	postUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single post record from the query.
func (q postQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Post, error) {
	o := &Post{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for posts")
	}

	return o, nil
}

// All returns all Post records from the query.
func (q postQuery) All(ctx context.Context, exec boil.ContextExecutor) (PostSlice, error) {
	var o []*Post

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Post slice")
	}

	return o, nil
}

// Count returns the count of all Post records in the query.
func (q postQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count posts rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q postQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if posts exists")
	}

	return count > 0, nil
}

// Posts retrieves all the records using an executor.
func Posts(mods ...qm.QueryMod) postQuery {
	mods = append(mods, qm.From("\"posts\""))
	return postQuery{NewQuery(mods...)}
}

// FindPost retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPost(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Post, error) {
	postObj := &Post{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"posts\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, postObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from posts")
	}

	return postObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Post) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no posts provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(postColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	postInsertCacheMut.RLock()
	cache, cached := postInsertCache[key]
	postInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			postAllColumns,
			postColumnsWithDefault,
			postColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(postType, postMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(postType, postMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"posts\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"posts\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into posts")
	}

	if !cached {
		postInsertCacheMut.Lock()
		postInsertCache[key] = cache
		postInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Post.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Post) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	postUpdateCacheMut.RLock()
	cache, cached := postUpdateCache[key]
	postUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			postAllColumns,
			postPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update posts, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"posts\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, postPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(postType, postMapping, append(wl, postPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update posts row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for posts")
	}

	if !cached {
		postUpdateCacheMut.Lock()
		postUpdateCache[key] = cache
		postUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q postQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for posts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for posts")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PostSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), postPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"posts\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, postPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in post slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all post")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Post) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no posts provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(postColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	postUpsertCacheMut.RLock()
	cache, cached := postUpsertCache[key]
	postUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			postAllColumns,
			postColumnsWithDefault,
			postColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			postAllColumns,
			postPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert posts, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(postPrimaryKeyColumns))
			copy(conflict, postPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"posts\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(postType, postMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(postType, postMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert posts")
	}

	if !cached {
		postUpsertCacheMut.Lock()
		postUpsertCache[key] = cache
		postUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Post record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Post) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Post provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), postPrimaryKeyMapping)
	sql := "DELETE FROM \"posts\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from posts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for posts")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q postQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no postQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from posts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for posts")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PostSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), postPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"posts\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, postPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from post slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for posts")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Post) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPost(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PostSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PostSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), postPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"posts\".* FROM \"posts\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, postPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PostSlice")
	}

	*o = slice

	return nil
}

// PostExists checks if the Post row exists.
func PostExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"posts\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if posts exists")
	}

	return exists, nil
}
