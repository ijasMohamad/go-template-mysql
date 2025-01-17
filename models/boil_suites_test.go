// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Articles", testArticles)
	t.Run("Authors", testAuthors)
	t.Run("GorpMigrations", testGorpMigrations)
}

func TestDelete(t *testing.T) {
	t.Run("Articles", testArticlesDelete)
	t.Run("Authors", testAuthorsDelete)
	t.Run("GorpMigrations", testGorpMigrationsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Articles", testArticlesQueryDeleteAll)
	t.Run("Authors", testAuthorsQueryDeleteAll)
	t.Run("GorpMigrations", testGorpMigrationsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Articles", testArticlesSliceDeleteAll)
	t.Run("Authors", testAuthorsSliceDeleteAll)
	t.Run("GorpMigrations", testGorpMigrationsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Articles", testArticlesExists)
	t.Run("Authors", testAuthorsExists)
	t.Run("GorpMigrations", testGorpMigrationsExists)
}

func TestFind(t *testing.T) {
	t.Run("Articles", testArticlesFind)
	t.Run("Authors", testAuthorsFind)
	t.Run("GorpMigrations", testGorpMigrationsFind)
}

func TestBind(t *testing.T) {
	t.Run("Articles", testArticlesBind)
	t.Run("Authors", testAuthorsBind)
	t.Run("GorpMigrations", testGorpMigrationsBind)
}

func TestOne(t *testing.T) {
	t.Run("Articles", testArticlesOne)
	t.Run("Authors", testAuthorsOne)
	t.Run("GorpMigrations", testGorpMigrationsOne)
}

func TestAll(t *testing.T) {
	t.Run("Articles", testArticlesAll)
	t.Run("Authors", testAuthorsAll)
	t.Run("GorpMigrations", testGorpMigrationsAll)
}

func TestCount(t *testing.T) {
	t.Run("Articles", testArticlesCount)
	t.Run("Authors", testAuthorsCount)
	t.Run("GorpMigrations", testGorpMigrationsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Articles", testArticlesHooks)
	t.Run("Authors", testAuthorsHooks)
	t.Run("GorpMigrations", testGorpMigrationsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Articles", testArticlesInsert)
	t.Run("Articles", testArticlesInsertWhitelist)
	t.Run("Authors", testAuthorsInsert)
	t.Run("Authors", testAuthorsInsertWhitelist)
	t.Run("GorpMigrations", testGorpMigrationsInsert)
	t.Run("GorpMigrations", testGorpMigrationsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("ArticleToAuthorUsingAuthor", testArticleToOneAuthorUsingAuthor)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("AuthorToArticles", testAuthorToManyArticles)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("ArticleToAuthorUsingArticles", testArticleToOneSetOpAuthorUsingAuthor)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("ArticleToAuthorUsingArticles", testArticleToOneRemoveOpAuthorUsingAuthor)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("AuthorToArticles", testAuthorToManyAddOpArticles)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("AuthorToArticles", testAuthorToManySetOpArticles)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("AuthorToArticles", testAuthorToManyRemoveOpArticles)
}

func TestReload(t *testing.T) {
	t.Run("Articles", testArticlesReload)
	t.Run("Authors", testAuthorsReload)
	t.Run("GorpMigrations", testGorpMigrationsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Articles", testArticlesReloadAll)
	t.Run("Authors", testAuthorsReloadAll)
	t.Run("GorpMigrations", testGorpMigrationsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Articles", testArticlesSelect)
	t.Run("Authors", testAuthorsSelect)
	t.Run("GorpMigrations", testGorpMigrationsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Articles", testArticlesUpdate)
	t.Run("Authors", testAuthorsUpdate)
	t.Run("GorpMigrations", testGorpMigrationsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Articles", testArticlesSliceUpdateAll)
	t.Run("Authors", testAuthorsSliceUpdateAll)
	t.Run("GorpMigrations", testGorpMigrationsSliceUpdateAll)
}
