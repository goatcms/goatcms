package dao

import (
	"database/sql"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	"github.com/goatcms/goatcore/app/scope"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestDropTable(t *testing.T) {
	t.Parallel()
	doDropTable(t)
}

func doDropTable(t *testing.T) (bool, *sql.DB) {
	var (
		db  *sql.DB
		err error
		ok  bool
	)
	if ok, db = doCreateTable(t); !ok {
		return false, nil
	}
	dropTable := FragmentDropTable{}
	dropTable.deps.DB = db
	s := scope.NewScope("testtag")
	if err := dropTable.DropTable(s); err != nil {
		t.Error(err)
		return false, db
	}
	if _, err = helpers.Commit(s); err != nil {
		t.Error(err)
		return false, db
	}
	return true, db
}
