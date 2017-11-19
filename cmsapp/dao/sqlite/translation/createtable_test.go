package translationdao

import (
	sqlitebase "github.com/goatcms/goatcms/cmsapp/dao/sqlite"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestCreateTable(t *testing.T) {
	t.Parallel()
	var (
		db  *sqlx.DB
		err error
	)
	if db, err = sqlitebase.NewMemoryDB(); err != nil {
		t.Error(err)
		return
	}
	createTable := TranslationCreateTable{}
	createTable.deps.DB = db
	s := scope.NewScope("testtag")
	if err = createTable.CreateTable(s); err != nil {
		t.Error(err)
		return
	}
}
