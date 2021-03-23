// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testSlashingEvents(t *testing.T) {
	t.Parallel()

	query := SlashingEvents()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSlashingEventsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SlashingEvent{}
	if err = randomize.Struct(seed, o, slashingEventDBTypes, true, slashingEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SlashingEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SlashingEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSlashingEventsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SlashingEvent{}
	if err = randomize.Struct(seed, o, slashingEventDBTypes, true, slashingEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SlashingEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := SlashingEvents().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SlashingEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSlashingEventsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SlashingEvent{}
	if err = randomize.Struct(seed, o, slashingEventDBTypes, true, slashingEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SlashingEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SlashingEventSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SlashingEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSlashingEventsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SlashingEvent{}
	if err = randomize.Struct(seed, o, slashingEventDBTypes, true, slashingEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SlashingEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SlashingEventExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if SlashingEvent exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SlashingEventExists to return true, but got false.")
	}
}

func testSlashingEventsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SlashingEvent{}
	if err = randomize.Struct(seed, o, slashingEventDBTypes, true, slashingEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SlashingEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slashingEventFound, err := FindSlashingEvent(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if slashingEventFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSlashingEventsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SlashingEvent{}
	if err = randomize.Struct(seed, o, slashingEventDBTypes, true, slashingEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SlashingEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = SlashingEvents().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSlashingEventsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SlashingEvent{}
	if err = randomize.Struct(seed, o, slashingEventDBTypes, true, slashingEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SlashingEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := SlashingEvents().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSlashingEventsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	slashingEventOne := &SlashingEvent{}
	slashingEventTwo := &SlashingEvent{}
	if err = randomize.Struct(seed, slashingEventOne, slashingEventDBTypes, false, slashingEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SlashingEvent struct: %s", err)
	}
	if err = randomize.Struct(seed, slashingEventTwo, slashingEventDBTypes, false, slashingEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SlashingEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = slashingEventOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = slashingEventTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := SlashingEvents().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSlashingEventsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slashingEventOne := &SlashingEvent{}
	slashingEventTwo := &SlashingEvent{}
	if err = randomize.Struct(seed, slashingEventOne, slashingEventDBTypes, false, slashingEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SlashingEvent struct: %s", err)
	}
	if err = randomize.Struct(seed, slashingEventTwo, slashingEventDBTypes, false, slashingEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SlashingEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = slashingEventOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = slashingEventTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SlashingEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testSlashingEventsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SlashingEvent{}
	if err = randomize.Struct(seed, o, slashingEventDBTypes, true, slashingEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SlashingEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SlashingEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSlashingEventsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SlashingEvent{}
	if err = randomize.Struct(seed, o, slashingEventDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SlashingEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(slashingEventColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := SlashingEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSlashingEventToOneValidatorUsingValidatorAddressValidator(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local SlashingEvent
	var foreign Validator

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, slashingEventDBTypes, false, slashingEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SlashingEvent struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, validatorDBTypes, false, validatorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Validator struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ValidatorAddress = foreign.Address
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.ValidatorAddressValidator().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.Address != foreign.Address {
		t.Errorf("want: %v, got %v", foreign.Address, check.Address)
	}

	slice := SlashingEventSlice{&local}
	if err = local.L.LoadValidatorAddressValidator(ctx, tx, false, (*[]*SlashingEvent)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ValidatorAddressValidator == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ValidatorAddressValidator = nil
	if err = local.L.LoadValidatorAddressValidator(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ValidatorAddressValidator == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testSlashingEventToOneSetOpValidatorUsingValidatorAddressValidator(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a SlashingEvent
	var b, c Validator

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, slashingEventDBTypes, false, strmangle.SetComplement(slashingEventPrimaryKeyColumns, slashingEventColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, validatorDBTypes, false, strmangle.SetComplement(validatorPrimaryKeyColumns, validatorColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, validatorDBTypes, false, strmangle.SetComplement(validatorPrimaryKeyColumns, validatorColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Validator{&b, &c} {
		err = a.SetValidatorAddressValidator(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ValidatorAddressValidator != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ValidatorAddressSlashingEvents[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ValidatorAddress != x.Address {
			t.Error("foreign key was wrong value", a.ValidatorAddress)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ValidatorAddress))
		reflect.Indirect(reflect.ValueOf(&a.ValidatorAddress)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ValidatorAddress != x.Address {
			t.Error("foreign key was wrong value", a.ValidatorAddress, x.Address)
		}
	}
}

func testSlashingEventsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SlashingEvent{}
	if err = randomize.Struct(seed, o, slashingEventDBTypes, true, slashingEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SlashingEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSlashingEventsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SlashingEvent{}
	if err = randomize.Struct(seed, o, slashingEventDBTypes, true, slashingEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SlashingEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SlashingEventSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSlashingEventsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SlashingEvent{}
	if err = randomize.Struct(seed, o, slashingEventDBTypes, true, slashingEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SlashingEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := SlashingEvents().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	slashingEventDBTypes = map[string]string{`ID`: `bigint`, `Height`: `bigint`, `ValidatorAddress`: `text`, `EventType`: `text`, `Counter`: `bigint`, `Reason`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `DeletedAt`: `timestamp without time zone`}
	_                    = bytes.MinRead
)

func testSlashingEventsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(slashingEventPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(slashingEventAllColumns) == len(slashingEventPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &SlashingEvent{}
	if err = randomize.Struct(seed, o, slashingEventDBTypes, true, slashingEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SlashingEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SlashingEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, slashingEventDBTypes, true, slashingEventPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SlashingEvent struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSlashingEventsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(slashingEventAllColumns) == len(slashingEventPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &SlashingEvent{}
	if err = randomize.Struct(seed, o, slashingEventDBTypes, true, slashingEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SlashingEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SlashingEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, slashingEventDBTypes, true, slashingEventPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SlashingEvent struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(slashingEventAllColumns, slashingEventPrimaryKeyColumns) {
		fields = slashingEventAllColumns
	} else {
		fields = strmangle.SetComplement(
			slashingEventAllColumns,
			slashingEventPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := SlashingEventSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSlashingEventsUpsert(t *testing.T) {
	t.Parallel()

	if len(slashingEventAllColumns) == len(slashingEventPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := SlashingEvent{}
	if err = randomize.Struct(seed, &o, slashingEventDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SlashingEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert SlashingEvent: %s", err)
	}

	count, err := SlashingEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, slashingEventDBTypes, false, slashingEventPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SlashingEvent struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert SlashingEvent: %s", err)
	}

	count, err = SlashingEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
