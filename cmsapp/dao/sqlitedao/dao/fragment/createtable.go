package dao

import (
	"database/sql"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

// FragmentCreateTable is a Data Access Object for fragment entity
type FragmentCreateTable struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewFragmentCreateTable(di dependency.Injector) (*FragmentCreateTable, error) {
	instance := &FragmentCreateTable{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func FragmentCreateTableFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewFragmentCreateTable(dp)
	if err != nil {
		return nil, err
	}
	return maindef.CreateTable(instance), nil
}

func (dao FragmentCreateTable) CreateTable(scope app.Scope) (err error) {
	var tx *sql.Tx
	if tx, err = helpers.TX(scope, dao.deps.DB); err != nil {
		return err
	}
	_, err = tx.Exec(dao.SQL())
	return err
}

func (dao FragmentCreateTable) SQL() string {
	return `CREATE TABLE IF NOT EXISTS Fragment (ID INTEGER PRIMARY KEY, Content TEXT NOT NULL, Name TEXT NOT NULL, Lang TEXT NOT NULL);`
}
