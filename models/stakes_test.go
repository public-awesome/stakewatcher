// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testStakes(t *testing.T) {
	t.Parallel()

	query := Stakes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testStakesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stake{}
	if err = randomize.Struct(seed, o, stakeDBTypes, true, stakeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stake struct: %s", err)
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

	count, err := Stakes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStakesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stake{}
	if err = randomize.Struct(seed, o, stakeDBTypes, true, stakeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stake struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Stakes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Stakes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStakesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stake{}
	if err = randomize.Struct(seed, o, stakeDBTypes, true, stakeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stake struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := StakeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Stakes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStakesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stake{}
	if err = randomize.Struct(seed, o, stakeDBTypes, true, stakeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stake struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := StakeExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Stake exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StakeExists to return true, but got false.")
	}
}

func testStakesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stake{}
	if err = randomize.Struct(seed, o, stakeDBTypes, true, stakeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stake struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	stakeFound, err := FindStake(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if stakeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testStakesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stake{}
	if err = randomize.Struct(seed, o, stakeDBTypes, true, stakeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stake struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Stakes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testStakesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stake{}
	if err = randomize.Struct(seed, o, stakeDBTypes, true, stakeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stake struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Stakes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStakesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stakeOne := &Stake{}
	stakeTwo := &Stake{}
	if err = randomize.Struct(seed, stakeOne, stakeDBTypes, false, stakeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stake struct: %s", err)
	}
	if err = randomize.Struct(seed, stakeTwo, stakeDBTypes, false, stakeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stake struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = stakeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = stakeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Stakes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStakesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	stakeOne := &Stake{}
	stakeTwo := &Stake{}
	if err = randomize.Struct(seed, stakeOne, stakeDBTypes, false, stakeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stake struct: %s", err)
	}
	if err = randomize.Struct(seed, stakeTwo, stakeDBTypes, false, stakeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stake struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = stakeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = stakeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Stakes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testStakesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stake{}
	if err = randomize.Struct(seed, o, stakeDBTypes, true, stakeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stake struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Stakes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStakesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stake{}
	if err = randomize.Struct(seed, o, stakeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stake struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(stakeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Stakes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStakesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stake{}
	if err = randomize.Struct(seed, o, stakeDBTypes, true, stakeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stake struct: %s", err)
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

func testStakesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stake{}
	if err = randomize.Struct(seed, o, stakeDBTypes, true, stakeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stake struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := StakeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testStakesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stake{}
	if err = randomize.Struct(seed, o, stakeDBTypes, true, stakeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stake struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Stakes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	stakeDBTypes = map[string]string{`ID`: `integer`, `Height`: `bigint`, `VendorID`: `integer`, `PostID`: `text`, `Delegator`: `text`, `Validator`: `text`, `Amount`: `bigint`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `DeletedAt`: `timestamp without time zone`}
	_            = bytes.MinRead
)

func testStakesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(stakePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(stakeAllColumns) == len(stakePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Stake{}
	if err = randomize.Struct(seed, o, stakeDBTypes, true, stakeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stake struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Stakes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, stakeDBTypes, true, stakePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Stake struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testStakesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(stakeAllColumns) == len(stakePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Stake{}
	if err = randomize.Struct(seed, o, stakeDBTypes, true, stakeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stake struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Stakes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, stakeDBTypes, true, stakePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Stake struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(stakeAllColumns, stakePrimaryKeyColumns) {
		fields = stakeAllColumns
	} else {
		fields = strmangle.SetComplement(
			stakeAllColumns,
			stakePrimaryKeyColumns,
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

	slice := StakeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testStakesUpsert(t *testing.T) {
	t.Parallel()

	if len(stakeAllColumns) == len(stakePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Stake{}
	if err = randomize.Struct(seed, &o, stakeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stake struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Stake: %s", err)
	}

	count, err := Stakes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, stakeDBTypes, false, stakePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Stake struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Stake: %s", err)
	}

	count, err = Stakes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
