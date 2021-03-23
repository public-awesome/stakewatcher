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

// Validator is an object representing the database table.
type Validator struct {
	Address         string    `boil:"address" json:"address" toml:"address" yaml:"address"`
	PubKey          string    `boil:"pub_key" json:"pubKey" toml:"pubKey" yaml:"pubKey"`
	CreatedAt       time.Time `boil:"created_at" json:"createdAt" toml:"createdAt" yaml:"createdAt"`
	UpdatedAt       time.Time `boil:"updated_at" json:"updatedAt" toml:"updatedAt" yaml:"updatedAt"`
	DeletedAt       null.Time `boil:"deleted_at" json:"deletedAt,omitempty" toml:"deletedAt" yaml:"deletedAt,omitempty"`
	OperatorAddress string    `boil:"operator_address" json:"operatorAddress" toml:"operatorAddress" yaml:"operatorAddress"`
	Moniker         string    `boil:"moniker" json:"moniker" toml:"moniker" yaml:"moniker"`

	R *validatorR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L validatorL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ValidatorColumns = struct {
	Address         string
	PubKey          string
	CreatedAt       string
	UpdatedAt       string
	DeletedAt       string
	OperatorAddress string
	Moniker         string
}{
	Address:         "address",
	PubKey:          "pub_key",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	DeletedAt:       "deleted_at",
	OperatorAddress: "operator_address",
	Moniker:         "moniker",
}

// Generated where

var ValidatorWhere = struct {
	Address         whereHelperstring
	PubKey          whereHelperstring
	CreatedAt       whereHelpertime_Time
	UpdatedAt       whereHelpertime_Time
	DeletedAt       whereHelpernull_Time
	OperatorAddress whereHelperstring
	Moniker         whereHelperstring
}{
	Address:         whereHelperstring{field: "\"validators\".\"address\""},
	PubKey:          whereHelperstring{field: "\"validators\".\"pub_key\""},
	CreatedAt:       whereHelpertime_Time{field: "\"validators\".\"created_at\""},
	UpdatedAt:       whereHelpertime_Time{field: "\"validators\".\"updated_at\""},
	DeletedAt:       whereHelpernull_Time{field: "\"validators\".\"deleted_at\""},
	OperatorAddress: whereHelperstring{field: "\"validators\".\"operator_address\""},
	Moniker:         whereHelperstring{field: "\"validators\".\"moniker\""},
}

// ValidatorRels is where relationship names are stored.
var ValidatorRels = struct {
	ValidatorAddressBlockSignatures string
	ProposerAddressBlocks           string
	ValidatorAddressSlashingEvents  string
}{
	ValidatorAddressBlockSignatures: "ValidatorAddressBlockSignatures",
	ProposerAddressBlocks:           "ProposerAddressBlocks",
	ValidatorAddressSlashingEvents:  "ValidatorAddressSlashingEvents",
}

// validatorR is where relationships are stored.
type validatorR struct {
	ValidatorAddressBlockSignatures BlockSignatureSlice `boil:"ValidatorAddressBlockSignatures" json:"ValidatorAddressBlockSignatures" toml:"ValidatorAddressBlockSignatures" yaml:"ValidatorAddressBlockSignatures"`
	ProposerAddressBlocks           BlockSlice          `boil:"ProposerAddressBlocks" json:"ProposerAddressBlocks" toml:"ProposerAddressBlocks" yaml:"ProposerAddressBlocks"`
	ValidatorAddressSlashingEvents  SlashingEventSlice  `boil:"ValidatorAddressSlashingEvents" json:"ValidatorAddressSlashingEvents" toml:"ValidatorAddressSlashingEvents" yaml:"ValidatorAddressSlashingEvents"`
}

// NewStruct creates a new relationship struct
func (*validatorR) NewStruct() *validatorR {
	return &validatorR{}
}

// validatorL is where Load methods for each relationship are stored.
type validatorL struct{}

var (
	validatorAllColumns            = []string{"address", "pub_key", "created_at", "updated_at", "deleted_at", "operator_address", "moniker"}
	validatorColumnsWithoutDefault = []string{"address", "pub_key", "deleted_at"}
	validatorColumnsWithDefault    = []string{"created_at", "updated_at", "operator_address", "moniker"}
	validatorPrimaryKeyColumns     = []string{"address"}
)

type (
	// ValidatorSlice is an alias for a slice of pointers to Validator.
	// This should generally be used opposed to []Validator.
	ValidatorSlice []*Validator

	validatorQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	validatorType                 = reflect.TypeOf(&Validator{})
	validatorMapping              = queries.MakeStructMapping(validatorType)
	validatorPrimaryKeyMapping, _ = queries.BindMapping(validatorType, validatorMapping, validatorPrimaryKeyColumns)
	validatorInsertCacheMut       sync.RWMutex
	validatorInsertCache          = make(map[string]insertCache)
	validatorUpdateCacheMut       sync.RWMutex
	validatorUpdateCache          = make(map[string]updateCache)
	validatorUpsertCacheMut       sync.RWMutex
	validatorUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single validator record from the query.
func (q validatorQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Validator, error) {
	o := &Validator{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for validators")
	}

	return o, nil
}

// All returns all Validator records from the query.
func (q validatorQuery) All(ctx context.Context, exec boil.ContextExecutor) (ValidatorSlice, error) {
	var o []*Validator

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Validator slice")
	}

	return o, nil
}

// Count returns the count of all Validator records in the query.
func (q validatorQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count validators rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q validatorQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if validators exists")
	}

	return count > 0, nil
}

// ValidatorAddressBlockSignatures retrieves all the block_signature's BlockSignatures with an executor via validator_address column.
func (o *Validator) ValidatorAddressBlockSignatures(mods ...qm.QueryMod) blockSignatureQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"block_signatures\".\"validator_address\"=?", o.Address),
	)

	query := BlockSignatures(queryMods...)
	queries.SetFrom(query.Query, "\"block_signatures\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"block_signatures\".*"})
	}

	return query
}

// ProposerAddressBlocks retrieves all the block's Blocks with an executor via proposer_address column.
func (o *Validator) ProposerAddressBlocks(mods ...qm.QueryMod) blockQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"blocks\".\"proposer_address\"=?", o.Address),
	)

	query := Blocks(queryMods...)
	queries.SetFrom(query.Query, "\"blocks\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"blocks\".*"})
	}

	return query
}

// ValidatorAddressSlashingEvents retrieves all the slashing_event's SlashingEvents with an executor via validator_address column.
func (o *Validator) ValidatorAddressSlashingEvents(mods ...qm.QueryMod) slashingEventQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"slashing_events\".\"validator_address\"=?", o.Address),
	)

	query := SlashingEvents(queryMods...)
	queries.SetFrom(query.Query, "\"slashing_events\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"slashing_events\".*"})
	}

	return query
}

// LoadValidatorAddressBlockSignatures allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (validatorL) LoadValidatorAddressBlockSignatures(ctx context.Context, e boil.ContextExecutor, singular bool, maybeValidator interface{}, mods queries.Applicator) error {
	var slice []*Validator
	var object *Validator

	if singular {
		object = maybeValidator.(*Validator)
	} else {
		slice = *maybeValidator.(*[]*Validator)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &validatorR{}
		}
		args = append(args, object.Address)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &validatorR{}
			}

			for _, a := range args {
				if a == obj.Address {
					continue Outer
				}
			}

			args = append(args, obj.Address)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`block_signatures`),
		qm.WhereIn(`block_signatures.validator_address in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load block_signatures")
	}

	var resultSlice []*BlockSignature
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice block_signatures")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on block_signatures")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for block_signatures")
	}

	if singular {
		object.R.ValidatorAddressBlockSignatures = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &blockSignatureR{}
			}
			foreign.R.Validator = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.Address == foreign.ValidatorAddress {
				local.R.ValidatorAddressBlockSignatures = append(local.R.ValidatorAddressBlockSignatures, foreign)
				if foreign.R == nil {
					foreign.R = &blockSignatureR{}
				}
				foreign.R.Validator = local
				break
			}
		}
	}

	return nil
}

// LoadProposerAddressBlocks allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (validatorL) LoadProposerAddressBlocks(ctx context.Context, e boil.ContextExecutor, singular bool, maybeValidator interface{}, mods queries.Applicator) error {
	var slice []*Validator
	var object *Validator

	if singular {
		object = maybeValidator.(*Validator)
	} else {
		slice = *maybeValidator.(*[]*Validator)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &validatorR{}
		}
		args = append(args, object.Address)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &validatorR{}
			}

			for _, a := range args {
				if a == obj.Address {
					continue Outer
				}
			}

			args = append(args, obj.Address)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`blocks`),
		qm.WhereIn(`blocks.proposer_address in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load blocks")
	}

	var resultSlice []*Block
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice blocks")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on blocks")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for blocks")
	}

	if singular {
		object.R.ProposerAddressBlocks = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &blockR{}
			}
			foreign.R.Proposer = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.Address == foreign.ProposerAddress {
				local.R.ProposerAddressBlocks = append(local.R.ProposerAddressBlocks, foreign)
				if foreign.R == nil {
					foreign.R = &blockR{}
				}
				foreign.R.Proposer = local
				break
			}
		}
	}

	return nil
}

// LoadValidatorAddressSlashingEvents allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (validatorL) LoadValidatorAddressSlashingEvents(ctx context.Context, e boil.ContextExecutor, singular bool, maybeValidator interface{}, mods queries.Applicator) error {
	var slice []*Validator
	var object *Validator

	if singular {
		object = maybeValidator.(*Validator)
	} else {
		slice = *maybeValidator.(*[]*Validator)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &validatorR{}
		}
		args = append(args, object.Address)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &validatorR{}
			}

			for _, a := range args {
				if a == obj.Address {
					continue Outer
				}
			}

			args = append(args, obj.Address)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`slashing_events`),
		qm.WhereIn(`slashing_events.validator_address in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load slashing_events")
	}

	var resultSlice []*SlashingEvent
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice slashing_events")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on slashing_events")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for slashing_events")
	}

	if singular {
		object.R.ValidatorAddressSlashingEvents = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &slashingEventR{}
			}
			foreign.R.ValidatorAddressValidator = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.Address == foreign.ValidatorAddress {
				local.R.ValidatorAddressSlashingEvents = append(local.R.ValidatorAddressSlashingEvents, foreign)
				if foreign.R == nil {
					foreign.R = &slashingEventR{}
				}
				foreign.R.ValidatorAddressValidator = local
				break
			}
		}
	}

	return nil
}

// AddValidatorAddressBlockSignatures adds the given related objects to the existing relationships
// of the validator, optionally inserting them as new records.
// Appends related to o.R.ValidatorAddressBlockSignatures.
// Sets related.R.Validator appropriately.
func (o *Validator) AddValidatorAddressBlockSignatures(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*BlockSignature) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ValidatorAddress = o.Address
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"block_signatures\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"validator_address"}),
				strmangle.WhereClause("\"", "\"", 2, blockSignaturePrimaryKeyColumns),
			)
			values := []interface{}{o.Address, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ValidatorAddress = o.Address
		}
	}

	if o.R == nil {
		o.R = &validatorR{
			ValidatorAddressBlockSignatures: related,
		}
	} else {
		o.R.ValidatorAddressBlockSignatures = append(o.R.ValidatorAddressBlockSignatures, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &blockSignatureR{
				Validator: o,
			}
		} else {
			rel.R.Validator = o
		}
	}
	return nil
}

// AddProposerAddressBlocks adds the given related objects to the existing relationships
// of the validator, optionally inserting them as new records.
// Appends related to o.R.ProposerAddressBlocks.
// Sets related.R.Proposer appropriately.
func (o *Validator) AddProposerAddressBlocks(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Block) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ProposerAddress = o.Address
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"blocks\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"proposer_address"}),
				strmangle.WhereClause("\"", "\"", 2, blockPrimaryKeyColumns),
			)
			values := []interface{}{o.Address, rel.Height}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ProposerAddress = o.Address
		}
	}

	if o.R == nil {
		o.R = &validatorR{
			ProposerAddressBlocks: related,
		}
	} else {
		o.R.ProposerAddressBlocks = append(o.R.ProposerAddressBlocks, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &blockR{
				Proposer: o,
			}
		} else {
			rel.R.Proposer = o
		}
	}
	return nil
}

// AddValidatorAddressSlashingEvents adds the given related objects to the existing relationships
// of the validator, optionally inserting them as new records.
// Appends related to o.R.ValidatorAddressSlashingEvents.
// Sets related.R.ValidatorAddressValidator appropriately.
func (o *Validator) AddValidatorAddressSlashingEvents(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*SlashingEvent) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ValidatorAddress = o.Address
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"slashing_events\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"validator_address"}),
				strmangle.WhereClause("\"", "\"", 2, slashingEventPrimaryKeyColumns),
			)
			values := []interface{}{o.Address, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ValidatorAddress = o.Address
		}
	}

	if o.R == nil {
		o.R = &validatorR{
			ValidatorAddressSlashingEvents: related,
		}
	} else {
		o.R.ValidatorAddressSlashingEvents = append(o.R.ValidatorAddressSlashingEvents, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &slashingEventR{
				ValidatorAddressValidator: o,
			}
		} else {
			rel.R.ValidatorAddressValidator = o
		}
	}
	return nil
}

// Validators retrieves all the records using an executor.
func Validators(mods ...qm.QueryMod) validatorQuery {
	mods = append(mods, qm.From("\"validators\""))
	return validatorQuery{NewQuery(mods...)}
}

// FindValidator retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindValidator(ctx context.Context, exec boil.ContextExecutor, address string, selectCols ...string) (*Validator, error) {
	validatorObj := &Validator{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"validators\" where \"address\"=$1", sel,
	)

	q := queries.Raw(query, address)

	err := q.Bind(ctx, exec, validatorObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from validators")
	}

	return validatorObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Validator) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no validators provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(validatorColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	validatorInsertCacheMut.RLock()
	cache, cached := validatorInsertCache[key]
	validatorInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			validatorAllColumns,
			validatorColumnsWithDefault,
			validatorColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(validatorType, validatorMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(validatorType, validatorMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"validators\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"validators\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into validators")
	}

	if !cached {
		validatorInsertCacheMut.Lock()
		validatorInsertCache[key] = cache
		validatorInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Validator.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Validator) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	validatorUpdateCacheMut.RLock()
	cache, cached := validatorUpdateCache[key]
	validatorUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			validatorAllColumns,
			validatorPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update validators, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"validators\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, validatorPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(validatorType, validatorMapping, append(wl, validatorPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update validators row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for validators")
	}

	if !cached {
		validatorUpdateCacheMut.Lock()
		validatorUpdateCache[key] = cache
		validatorUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q validatorQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for validators")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for validators")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ValidatorSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), validatorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"validators\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, validatorPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in validator slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all validator")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Validator) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no validators provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(validatorColumnsWithDefault, o)

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

	validatorUpsertCacheMut.RLock()
	cache, cached := validatorUpsertCache[key]
	validatorUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			validatorAllColumns,
			validatorColumnsWithDefault,
			validatorColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			validatorAllColumns,
			validatorPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert validators, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(validatorPrimaryKeyColumns))
			copy(conflict, validatorPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"validators\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(validatorType, validatorMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(validatorType, validatorMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert validators")
	}

	if !cached {
		validatorUpsertCacheMut.Lock()
		validatorUpsertCache[key] = cache
		validatorUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Validator record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Validator) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Validator provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), validatorPrimaryKeyMapping)
	sql := "DELETE FROM \"validators\" WHERE \"address\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from validators")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for validators")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q validatorQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no validatorQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from validators")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for validators")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ValidatorSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), validatorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"validators\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, validatorPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from validator slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for validators")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Validator) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindValidator(ctx, exec, o.Address)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ValidatorSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ValidatorSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), validatorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"validators\".* FROM \"validators\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, validatorPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ValidatorSlice")
	}

	*o = slice

	return nil
}

// ValidatorExists checks if the Validator row exists.
func ValidatorExists(ctx context.Context, exec boil.ContextExecutor, address string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"validators\" where \"address\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, address)
	}
	row := exec.QueryRowContext(ctx, sql, address)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if validators exists")
	}

	return exists, nil
}
