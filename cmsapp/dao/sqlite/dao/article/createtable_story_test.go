package articledao

import (
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlite/helpers"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestCreateTable(t *testing.T) {
	t.Parallel()
	doCreateTable(t)
}

func doCreateTable(t *testing.T) (bool, *sqlx.DB) {
	var (
		db  *sqlx.DB
		err error
	)
	if db, err = helpers.NewMemoryDB(); err != nil {
		t.Error(err)
		return false, db
	}
	createTable := ArticleCreateTable{}
	createTable.deps.DB = db
	s := scope.NewScope("testtag")
	if err := createTable.CreateTable(s); err != nil {
		t.Error(err)
		return false, db
	}
	if _, err = helpers.Commit(s); err != nil {
		t.Error(err)
		return false, db
	}
	return true, db
}
