// Code generated by SQLBoiler 4.1.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// BlockSignature is an object representing the database table.
type BlockSignature struct {
	ID               int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Height           int64     `boil:"height" json:"height" toml:"height" yaml:"height"`
	Round            int       `boil:"round" json:"round" toml:"round" yaml:"round"`
	ValidatorAddress string    `boil:"validator_address" json:"validatorAddress" toml:"validatorAddress" yaml:"validatorAddress"`
	Flag             int       `boil:"flag" json:"flag" toml:"flag" yaml:"flag"`
	Timestamp        time.Time `boil:"timestamp" json:"timestamp" toml:"timestamp" yaml:"timestamp"`
	Hash             string    `boil:"hash" json:"hash" toml:"hash" yaml:"hash"`
	VotingPower      int       `boil:"voting_power" json:"votingPower" toml:"votingPower" yaml:"votingPower"`
	ProposerPriority int       `boil:"proposer_priority" json:"proposerPriority" toml:"proposerPriority" yaml:"proposerPriority"`
	CreatedAt        time.Time `boil:"created_at" json:"createdAt" toml:"createdAt" yaml:"createdAt"`
	UpdatedAt        time.Time `boil:"updated_at" json:"updatedAt" toml:"updatedAt" yaml:"updatedAt"`
	DeletedAt        null.Time `boil:"deleted_at" json:"deletedAt,omitempty" toml:"deletedAt" yaml:"deletedAt,omitempty"`

	R *blockSignatureR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L blockSignatureL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BlockSignatureColumns = struct {
	ID               string
	Height           string
	Round            string
	ValidatorAddress string
	Flag             string
	Timestamp        string
	Hash             string
	VotingPower      string
	ProposerPriority string
	CreatedAt        string
	UpdatedAt        string
	DeletedAt        string
}{
	ID:               "id",
	Height:           "height",
	Round:            "round",
	ValidatorAddress: "validator_address",
	Flag:             "flag",
	Timestamp:        "timestamp",
	Hash:             "hash",
	VotingPower:      "voting_power",
	ProposerPriority: "proposer_priority",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	DeletedAt:        "deleted_at",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var BlockSignatureWhere = struct {
	ID               whereHelperint
	Height           whereHelperint64
	Round            whereHelperint
	ValidatorAddress whereHelperstring
	Flag             whereHelperint
	Timestamp        whereHelpertime_Time
	Hash             whereHelperstring
	VotingPower      whereHelperint
	ProposerPriority whereHelperint
	CreatedAt        whereHelpertime_Time
	UpdatedAt        whereHelpertime_Time
	DeletedAt        whereHelpernull_Time
}{
	ID:               whereHelperint{field: "\"block_signatures\".\"id\""},
	Height:           whereHelperint64{field: "\"block_signatures\".\"height\""},
	Round:            whereHelperint{field: "\"block_signatures\".\"round\""},
	ValidatorAddress: whereHelperstring{field: "\"block_signatures\".\"validator_address\""},
	Flag:             whereHelperint{field: "\"block_signatures\".\"flag\""},
	Timestamp:        whereHelpertime_Time{field: "\"block_signatures\".\"timestamp\""},
	Hash:             whereHelperstring{field: "\"block_signatures\".\"hash\""},
	VotingPower:      whereHelperint{field: "\"block_signatures\".\"voting_power\""},
	ProposerPriority: whereHelperint{field: "\"block_signatures\".\"proposer_priority\""},
	CreatedAt:        whereHelpertime_Time{field: "\"block_signatures\".\"created_at\""},
	UpdatedAt:        whereHelpertime_Time{field: "\"block_signatures\".\"updated_at\""},
	DeletedAt:        whereHelpernull_Time{field: "\"block_signatures\".\"deleted_at\""},
}

// BlockSignatureRels is where relationship names are stored.
var BlockSignatureRels = struct {
	Validator string
}{
	Validator: "Validator",
}

// blockSignatureR is where relationships are stored.
type blockSignatureR struct {
	Validator *Validator `boil:"Validator" json:"Validator" toml:"Validator" yaml:"Validator"`
}

// NewStruct creates a new relationship struct
func (*blockSignatureR) NewStruct() *blockSignatureR {
	return &blockSignatureR{}
}

// blockSignatureL is where Load methods for each relationship are stored.
type blockSignatureL struct{}

var (
	blockSignatureAllColumns            = []string{"id", "height", "round", "validator_address", "flag", "timestamp", "hash", "voting_power", "proposer_priority", "created_at", "updated_at", "deleted_at"}
	blockSignatureColumnsWithoutDefault = []string{"height", "round", "validator_address", "timestamp", "hash", "voting_power", "proposer_priority", "deleted_at"}
	blockSignatureColumnsWithDefault    = []string{"id", "flag", "created_at", "updated_at"}
	blockSignaturePrimaryKeyColumns     = []string{"id"}
)

type (
	// BlockSignatureSlice is an alias for a slice of pointers to BlockSignature.
	// This should generally be used opposed to []BlockSignature.
	BlockSignatureSlice []*BlockSignature

	blockSignatureQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	blockSignatureType                 = reflect.TypeOf(&BlockSignature{})
	blockSignatureMapping              = queries.MakeStructMapping(blockSignatureType)
	blockSignaturePrimaryKeyMapping, _ = queries.BindMapping(blockSignatureType, blockSignatureMapping, blockSignaturePrimaryKeyColumns)
	blockSignatureInsertCacheMut       sync.RWMutex
	blockSignatureInsertCache          = make(map[string]insertCache)
	blockSignatureUpdateCacheMut       sync.RWMutex
	blockSignatureUpdateCache          = make(map[string]updateCache)
	blockSignatureUpsertCacheMut       sync.RWMutex
	blockSignatureUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single blockSignature record from the query.
func (q blockSignatureQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BlockSignature, error) {
	o := &BlockSignature{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for block_signatures")
	}

	return o, nil
}

// All returns all BlockSignature records from the query.
func (q blockSignatureQuery) All(ctx context.Context, exec boil.ContextExecutor) (BlockSignatureSlice, error) {
	var o []*BlockSignature

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BlockSignature slice")
	}

	return o, nil
}

// Count returns the count of all BlockSignature records in the query.
func (q blockSignatureQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count block_signatures rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q blockSignatureQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if block_signatures exists")
	}

	return count > 0, nil
}

// Validator pointed to by the foreign key.
func (o *BlockSignature) Validator(mods ...qm.QueryMod) validatorQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"address\" = ?", o.ValidatorAddress),
	}

	queryMods = append(queryMods, mods...)

	query := Validators(queryMods...)
	queries.SetFrom(query.Query, "\"validators\"")

	return query
}

// LoadValidator allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (blockSignatureL) LoadValidator(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBlockSignature interface{}, mods queries.Applicator) error {
	var slice []*BlockSignature
	var object *BlockSignature

	if singular {
		object = maybeBlockSignature.(*BlockSignature)
	} else {
		slice = *maybeBlockSignature.(*[]*BlockSignature)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &blockSignatureR{}
		}
		args = append(args, object.ValidatorAddress)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &blockSignatureR{}
			}

			for _, a := range args {
				if a == obj.ValidatorAddress {
					continue Outer
				}
			}

			args = append(args, obj.ValidatorAddress)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`validators`),
		qm.WhereIn(`validators.address in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Validator")
	}

	var resultSlice []*Validator
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Validator")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for validators")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for validators")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Validator = foreign
		if foreign.R == nil {
			foreign.R = &validatorR{}
		}
		foreign.R.ValidatorAddressBlockSignatures = append(foreign.R.ValidatorAddressBlockSignatures, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ValidatorAddress == foreign.Address {
				local.R.Validator = foreign
				if foreign.R == nil {
					foreign.R = &validatorR{}
				}
				foreign.R.ValidatorAddressBlockSignatures = append(foreign.R.ValidatorAddressBlockSignatures, local)
				break
			}
		}
	}

	return nil
}

// SetValidator of the blockSignature to the related item.
// Sets o.R.Validator to related.
// Adds o to related.R.ValidatorAddressBlockSignatures.
func (o *BlockSignature) SetValidator(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Validator) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"block_signatures\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"validator_address"}),
		strmangle.WhereClause("\"", "\"", 2, blockSignaturePrimaryKeyColumns),
	)
	values := []interface{}{related.Address, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ValidatorAddress = related.Address
	if o.R == nil {
		o.R = &blockSignatureR{
			Validator: related,
		}
	} else {
		o.R.Validator = related
	}

	if related.R == nil {
		related.R = &validatorR{
			ValidatorAddressBlockSignatures: BlockSignatureSlice{o},
		}
	} else {
		related.R.ValidatorAddressBlockSignatures = append(related.R.ValidatorAddressBlockSignatures, o)
	}

	return nil
}

// BlockSignatures retrieves all the records using an executor.
func BlockSignatures(mods ...qm.QueryMod) blockSignatureQuery {
	mods = append(mods, qm.From("\"block_signatures\""))
	return blockSignatureQuery{NewQuery(mods...)}
}

// FindBlockSignature retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBlockSignature(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*BlockSignature, error) {
	blockSignatureObj := &BlockSignature{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"block_signatures\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, blockSignatureObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from block_signatures")
	}

	return blockSignatureObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *BlockSignature) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no block_signatures provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(blockSignatureColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	blockSignatureInsertCacheMut.RLock()
	cache, cached := blockSignatureInsertCache[key]
	blockSignatureInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			blockSignatureAllColumns,
			blockSignatureColumnsWithDefault,
			blockSignatureColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(blockSignatureType, blockSignatureMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(blockSignatureType, blockSignatureMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"block_signatures\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"block_signatures\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into block_signatures")
	}

	if !cached {
		blockSignatureInsertCacheMut.Lock()
		blockSignatureInsertCache[key] = cache
		blockSignatureInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the BlockSignature.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *BlockSignature) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	blockSignatureUpdateCacheMut.RLock()
	cache, cached := blockSignatureUpdateCache[key]
	blockSignatureUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			blockSignatureAllColumns,
			blockSignaturePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update block_signatures, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"block_signatures\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, blockSignaturePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(blockSignatureType, blockSignatureMapping, append(wl, blockSignaturePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update block_signatures row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for block_signatures")
	}

	if !cached {
		blockSignatureUpdateCacheMut.Lock()
		blockSignatureUpdateCache[key] = cache
		blockSignatureUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q blockSignatureQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for block_signatures")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for block_signatures")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BlockSignatureSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), blockSignaturePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"block_signatures\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, blockSignaturePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in blockSignature slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all blockSignature")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *BlockSignature) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no block_signatures provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(blockSignatureColumnsWithDefault, o)

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

	blockSignatureUpsertCacheMut.RLock()
	cache, cached := blockSignatureUpsertCache[key]
	blockSignatureUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			blockSignatureAllColumns,
			blockSignatureColumnsWithDefault,
			blockSignatureColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			blockSignatureAllColumns,
			blockSignaturePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert block_signatures, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(blockSignaturePrimaryKeyColumns))
			copy(conflict, blockSignaturePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"block_signatures\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(blockSignatureType, blockSignatureMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(blockSignatureType, blockSignatureMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert block_signatures")
	}

	if !cached {
		blockSignatureUpsertCacheMut.Lock()
		blockSignatureUpsertCache[key] = cache
		blockSignatureUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single BlockSignature record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BlockSignature) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no BlockSignature provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), blockSignaturePrimaryKeyMapping)
	sql := "DELETE FROM \"block_signatures\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from block_signatures")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for block_signatures")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q blockSignatureQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no blockSignatureQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from block_signatures")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for block_signatures")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BlockSignatureSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), blockSignaturePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"block_signatures\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, blockSignaturePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from blockSignature slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for block_signatures")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *BlockSignature) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBlockSignature(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BlockSignatureSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BlockSignatureSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), blockSignaturePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"block_signatures\".* FROM \"block_signatures\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, blockSignaturePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BlockSignatureSlice")
	}

	*o = slice

	return nil
}

// BlockSignatureExists checks if the BlockSignature row exists.
func BlockSignatureExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"block_signatures\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if block_signatures exists")
	}

	return exists, nil
}
