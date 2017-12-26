package userdao

import (
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlite/helpers"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestDropTable(t *testing.T) {
	t.Parallel()
	doDropTable(t)
}

func doDropTable(t *testing.T) (bool, *sqlx.DB) {
	var (
		db  *sqlx.DB
		err error
		ok  bool
	)
	if ok, db = doCreateTable(t); !ok {
		return false, nil
	}
	dropTable := UserDropTable{}
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
