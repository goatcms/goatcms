package dao

import (
	"database/sql"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	"github.com/goatcms/goatcore/app/scope"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestCreateTable(t *testing.T) {
	t.Parallel()
	doCreateTable(t)
}

func doCreateTable(t *testing.T) (bool, *sql.DB) {
	var (
		db  *sql.DB
		err error
	)
	if db, err = helpers.NewMemoryDB(); err != nil {
		t.Error(err)
		return false, db
	}
	createTable := FragmentCreateTable{}
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
