// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testAuthors(t *testing.T) {
	t.Parallel()

	query := Authors()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testAuthorsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Author{}
	if err = randomize.Struct(seed, o, authorDBTypes, true, authorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Author struct: %s", err)
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

	count, err := Authors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthorsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Author{}
	if err = randomize.Struct(seed, o, authorDBTypes, true, authorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Author struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Authors().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Authors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthorsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Author{}
	if err = randomize.Struct(seed, o, authorDBTypes, true, authorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Author struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AuthorSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Authors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthorsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Author{}
	if err = randomize.Struct(seed, o, authorDBTypes, true, authorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Author struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := AuthorExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Author exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AuthorExists to return true, but got false.")
	}
}

func testAuthorsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Author{}
	if err = randomize.Struct(seed, o, authorDBTypes, true, authorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Author struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	authorFound, err := FindAuthor(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if authorFound == nil {
		t.Error("want a record, got nil")
	}
}

func testAuthorsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Author{}
	if err = randomize.Struct(seed, o, authorDBTypes, true, authorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Author struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Authors().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testAuthorsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Author{}
	if err = randomize.Struct(seed, o, authorDBTypes, true, authorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Author struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Authors().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAuthorsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authorOne := &Author{}
	authorTwo := &Author{}
	if err = randomize.Struct(seed, authorOne, authorDBTypes, false, authorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Author struct: %s", err)
	}
	if err = randomize.Struct(seed, authorTwo, authorDBTypes, false, authorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Author struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = authorOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = authorTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Authors().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAuthorsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	authorOne := &Author{}
	authorTwo := &Author{}
	if err = randomize.Struct(seed, authorOne, authorDBTypes, false, authorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Author struct: %s", err)
	}
	if err = randomize.Struct(seed, authorTwo, authorDBTypes, false, authorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Author struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = authorOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = authorTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Authors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func authorBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Author) error {
	*o = Author{}
	return nil
}

func authorAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Author) error {
	*o = Author{}
	return nil
}

func authorAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Author) error {
	*o = Author{}
	return nil
}

func authorBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Author) error {
	*o = Author{}
	return nil
}

func authorAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Author) error {
	*o = Author{}
	return nil
}

func authorBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Author) error {
	*o = Author{}
	return nil
}

func authorAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Author) error {
	*o = Author{}
	return nil
}

func authorBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Author) error {
	*o = Author{}
	return nil
}

func authorAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Author) error {
	*o = Author{}
	return nil
}

func testAuthorsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Author{}
	o := &Author{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, authorDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Author object: %s", err)
	}

	AddAuthorHook(boil.BeforeInsertHook, authorBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	authorBeforeInsertHooks = []AuthorHook{}

	AddAuthorHook(boil.AfterInsertHook, authorAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	authorAfterInsertHooks = []AuthorHook{}

	AddAuthorHook(boil.AfterSelectHook, authorAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	authorAfterSelectHooks = []AuthorHook{}

	AddAuthorHook(boil.BeforeUpdateHook, authorBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	authorBeforeUpdateHooks = []AuthorHook{}

	AddAuthorHook(boil.AfterUpdateHook, authorAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	authorAfterUpdateHooks = []AuthorHook{}

	AddAuthorHook(boil.BeforeDeleteHook, authorBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	authorBeforeDeleteHooks = []AuthorHook{}

	AddAuthorHook(boil.AfterDeleteHook, authorAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	authorAfterDeleteHooks = []AuthorHook{}

	AddAuthorHook(boil.BeforeUpsertHook, authorBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	authorBeforeUpsertHooks = []AuthorHook{}

	AddAuthorHook(boil.AfterUpsertHook, authorAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	authorAfterUpsertHooks = []AuthorHook{}
}

func testAuthorsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Author{}
	if err = randomize.Struct(seed, o, authorDBTypes, true, authorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Author struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Authors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthorsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Author{}
	if err = randomize.Struct(seed, o, authorDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Author struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(authorColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Authors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthorToManyArticles(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Author
	var b, c Article

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authorDBTypes, true, authorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Author struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, articleDBTypes, false, articleColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, articleDBTypes, false, articleColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.AuthorID, a.ID)
	queries.Assign(&c.AuthorID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Articles().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.AuthorID, b.AuthorID) {
			bFound = true
		}
		if queries.Equal(v.AuthorID, c.AuthorID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AuthorSlice{&a}
	if err = a.L.LoadArticles(ctx, tx, false, (*[]*Author)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Articles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Articles = nil
	if err = a.L.LoadArticles(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Articles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testAuthorToManyAddOpArticles(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Author
	var b, c, d, e Article

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authorDBTypes, false, strmangle.SetComplement(authorPrimaryKeyColumns, authorColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Article{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, articleDBTypes, false, strmangle.SetComplement(articlePrimaryKeyColumns, articleColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Article{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddArticles(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.AuthorID) {
			t.Error("foreign key was wrong value", a.ID, first.AuthorID)
		}
		if !queries.Equal(a.ID, second.AuthorID) {
			t.Error("foreign key was wrong value", a.ID, second.AuthorID)
		}

		if first.R.Author != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Author != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Articles[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Articles[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Articles().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testAuthorToManySetOpArticles(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Author
	var b, c, d, e Article

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authorDBTypes, false, strmangle.SetComplement(authorPrimaryKeyColumns, authorColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Article{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, articleDBTypes, false, strmangle.SetComplement(articlePrimaryKeyColumns, articleColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetArticles(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Articles().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetArticles(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Articles().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.AuthorID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.AuthorID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.AuthorID) {
		t.Error("foreign key was wrong value", a.ID, d.AuthorID)
	}
	if !queries.Equal(a.ID, e.AuthorID) {
		t.Error("foreign key was wrong value", a.ID, e.AuthorID)
	}

	if b.R.Author != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Author != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Author != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Author != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Articles[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Articles[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testAuthorToManyRemoveOpArticles(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Author
	var b, c, d, e Article

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authorDBTypes, false, strmangle.SetComplement(authorPrimaryKeyColumns, authorColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Article{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, articleDBTypes, false, strmangle.SetComplement(articlePrimaryKeyColumns, articleColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddArticles(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Articles().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveArticles(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Articles().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.AuthorID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.AuthorID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Author != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Author != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Author != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Author != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Articles) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Articles[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Articles[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testAuthorsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Author{}
	if err = randomize.Struct(seed, o, authorDBTypes, true, authorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Author struct: %s", err)
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

func testAuthorsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Author{}
	if err = randomize.Struct(seed, o, authorDBTypes, true, authorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Author struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AuthorSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAuthorsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Author{}
	if err = randomize.Struct(seed, o, authorDBTypes, true, authorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Author struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Authors().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	authorDBTypes = map[string]string{`ID`: `int`, `FirstName`: `text`, `LastName`: `text`, `Username`: `varchar`, `Password`: `text`, `Active`: `tinyint`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `DeletedAt`: `timestamp`}
	_             = bytes.MinRead
)

func testAuthorsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(authorPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(authorAllColumns) == len(authorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Author{}
	if err = randomize.Struct(seed, o, authorDBTypes, true, authorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Author struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Authors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, authorDBTypes, true, authorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Author struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testAuthorsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(authorAllColumns) == len(authorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Author{}
	if err = randomize.Struct(seed, o, authorDBTypes, true, authorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Author struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Authors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, authorDBTypes, true, authorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Author struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(authorAllColumns, authorPrimaryKeyColumns) {
		fields = authorAllColumns
	} else {
		fields = strmangle.SetComplement(
			authorAllColumns,
			authorPrimaryKeyColumns,
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

	slice := AuthorSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testAuthorsUpsert(t *testing.T) {
	t.Parallel()

	if len(authorAllColumns) == len(authorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLAuthorUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Author{}
	if err = randomize.Struct(seed, &o, authorDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Author struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Author: %s", err)
	}

	count, err := Authors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, authorDBTypes, false, authorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Author struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Author: %s", err)
	}

	count, err = Authors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
