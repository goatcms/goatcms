package dao

import (
	"database/sql"
	"github.com/goatcms/goatcore/app/scope"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestDeleteEntityStory(t *testing.T) {
	t.Parallel()
	doDeleteEntityStory(t)
}

func doDeleteEntityStory(t *testing.T) (bool, *sql.DB) {
	ok, db, entity := doInsertStory(t)
	if !ok {
		return false, db
	}
	deleteService := TranslationDelete{}
	deleteService.deps.DB = db
	s := scope.NewScope("testtag")
	if err := deleteService.Delete(s, *entity.ID); err != nil {
		t.Error(err)
		return false, db
	}
	return true, db
}
