package dao

import (
	"github.com/goatcms/goatcore/app/scope"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestDeleteEntityStory(t *testing.T) {
	t.Parallel()
	doDeleteEntityStory(t)
}

func doDeleteEntityStory(t *testing.T) (bool, *sqlx.DB) {
	ok, db, entity := doInsertStory(t)
	if !ok {
		return false, db
	}
	deleteSeervice := TranslationDelete{}
	deleteSeervice.deps.DB = db
	s := scope.NewScope("testtag")
	if err := deleteSeervice.Delete(s, entity.ID); err != nil {
		t.Error(err)
		return false, db
	}
	return true, db
}
