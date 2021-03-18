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

func testValidators(t *testing.T) {
	t.Parallel()

	query := Validators()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testValidatorsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Validator{}
	if err = randomize.Struct(seed, o, validatorDBTypes, true, validatorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Validator struct: %s", err)
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

	count, err := Validators().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testValidatorsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Validator{}
	if err = randomize.Struct(seed, o, validatorDBTypes, true, validatorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Validator struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Validators().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Validators().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testValidatorsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Validator{}
	if err = randomize.Struct(seed, o, validatorDBTypes, true, validatorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Validator struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ValidatorSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Validators().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testValidatorsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Validator{}
	if err = randomize.Struct(seed, o, validatorDBTypes, true, validatorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Validator struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ValidatorExists(ctx, tx, o.Address)
	if err != nil {
		t.Errorf("Unable to check if Validator exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ValidatorExists to return true, but got false.")
	}
}

func testValidatorsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Validator{}
	if err = randomize.Struct(seed, o, validatorDBTypes, true, validatorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Validator struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	validatorFound, err := FindValidator(ctx, tx, o.Address)
	if err != nil {
		t.Error(err)
	}

	if validatorFound == nil {
		t.Error("want a record, got nil")
	}
}

func testValidatorsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Validator{}
	if err = randomize.Struct(seed, o, validatorDBTypes, true, validatorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Validator struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Validators().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testValidatorsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Validator{}
	if err = randomize.Struct(seed, o, validatorDBTypes, true, validatorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Validator struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Validators().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testValidatorsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	validatorOne := &Validator{}
	validatorTwo := &Validator{}
	if err = randomize.Struct(seed, validatorOne, validatorDBTypes, false, validatorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Validator struct: %s", err)
	}
	if err = randomize.Struct(seed, validatorTwo, validatorDBTypes, false, validatorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Validator struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = validatorOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = validatorTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Validators().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testValidatorsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	validatorOne := &Validator{}
	validatorTwo := &Validator{}
	if err = randomize.Struct(seed, validatorOne, validatorDBTypes, false, validatorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Validator struct: %s", err)
	}
	if err = randomize.Struct(seed, validatorTwo, validatorDBTypes, false, validatorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Validator struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = validatorOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = validatorTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Validators().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testValidatorsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Validator{}
	if err = randomize.Struct(seed, o, validatorDBTypes, true, validatorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Validator struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Validators().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testValidatorsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Validator{}
	if err = randomize.Struct(seed, o, validatorDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Validator struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(validatorColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Validators().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testValidatorToManyValidatorAddressBlockSignatures(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Validator
	var b, c BlockSignature

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, validatorDBTypes, true, validatorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Validator struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, blockSignatureDBTypes, false, blockSignatureColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, blockSignatureDBTypes, false, blockSignatureColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ValidatorAddress = a.Address
	c.ValidatorAddress = a.Address

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.ValidatorAddressBlockSignatures().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ValidatorAddress == b.ValidatorAddress {
			bFound = true
		}
		if v.ValidatorAddress == c.ValidatorAddress {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ValidatorSlice{&a}
	if err = a.L.LoadValidatorAddressBlockSignatures(ctx, tx, false, (*[]*Validator)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ValidatorAddressBlockSignatures); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ValidatorAddressBlockSignatures = nil
	if err = a.L.LoadValidatorAddressBlockSignatures(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ValidatorAddressBlockSignatures); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testValidatorToManyProposerAddressBlocks(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Validator
	var b, c Block

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, validatorDBTypes, true, validatorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Validator struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, blockDBTypes, false, blockColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, blockDBTypes, false, blockColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ProposerAddress = a.Address
	c.ProposerAddress = a.Address

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.ProposerAddressBlocks().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ProposerAddress == b.ProposerAddress {
			bFound = true
		}
		if v.ProposerAddress == c.ProposerAddress {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ValidatorSlice{&a}
	if err = a.L.LoadProposerAddressBlocks(ctx, tx, false, (*[]*Validator)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ProposerAddressBlocks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ProposerAddressBlocks = nil
	if err = a.L.LoadProposerAddressBlocks(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ProposerAddressBlocks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testValidatorToManyAddOpValidatorAddressBlockSignatures(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Validator
	var b, c, d, e BlockSignature

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, validatorDBTypes, false, strmangle.SetComplement(validatorPrimaryKeyColumns, validatorColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BlockSignature{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, blockSignatureDBTypes, false, strmangle.SetComplement(blockSignaturePrimaryKeyColumns, blockSignatureColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*BlockSignature{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddValidatorAddressBlockSignatures(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.Address != first.ValidatorAddress {
			t.Error("foreign key was wrong value", a.Address, first.ValidatorAddress)
		}
		if a.Address != second.ValidatorAddress {
			t.Error("foreign key was wrong value", a.Address, second.ValidatorAddress)
		}

		if first.R.Validator != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Validator != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ValidatorAddressBlockSignatures[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ValidatorAddressBlockSignatures[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ValidatorAddressBlockSignatures().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testValidatorToManyAddOpProposerAddressBlocks(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Validator
	var b, c, d, e Block

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, validatorDBTypes, false, strmangle.SetComplement(validatorPrimaryKeyColumns, validatorColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Block{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Block{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddProposerAddressBlocks(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.Address != first.ProposerAddress {
			t.Error("foreign key was wrong value", a.Address, first.ProposerAddress)
		}
		if a.Address != second.ProposerAddress {
			t.Error("foreign key was wrong value", a.Address, second.ProposerAddress)
		}

		if first.R.Proposer != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Proposer != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ProposerAddressBlocks[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ProposerAddressBlocks[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ProposerAddressBlocks().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testValidatorsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Validator{}
	if err = randomize.Struct(seed, o, validatorDBTypes, true, validatorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Validator struct: %s", err)
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

func testValidatorsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Validator{}
	if err = randomize.Struct(seed, o, validatorDBTypes, true, validatorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Validator struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ValidatorSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testValidatorsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Validator{}
	if err = randomize.Struct(seed, o, validatorDBTypes, true, validatorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Validator struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Validators().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	validatorDBTypes = map[string]string{`Address`: `text`, `PubKey`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `DeletedAt`: `timestamp without time zone`}
	_                = bytes.MinRead
)

func testValidatorsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(validatorPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(validatorAllColumns) == len(validatorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Validator{}
	if err = randomize.Struct(seed, o, validatorDBTypes, true, validatorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Validator struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Validators().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, validatorDBTypes, true, validatorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Validator struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testValidatorsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(validatorAllColumns) == len(validatorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Validator{}
	if err = randomize.Struct(seed, o, validatorDBTypes, true, validatorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Validator struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Validators().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, validatorDBTypes, true, validatorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Validator struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(validatorAllColumns, validatorPrimaryKeyColumns) {
		fields = validatorAllColumns
	} else {
		fields = strmangle.SetComplement(
			validatorAllColumns,
			validatorPrimaryKeyColumns,
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

	slice := ValidatorSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testValidatorsUpsert(t *testing.T) {
	t.Parallel()

	if len(validatorAllColumns) == len(validatorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Validator{}
	if err = randomize.Struct(seed, &o, validatorDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Validator struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Validator: %s", err)
	}

	count, err := Validators().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, validatorDBTypes, false, validatorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Validator struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Validator: %s", err)
	}

	count, err = Validators().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
