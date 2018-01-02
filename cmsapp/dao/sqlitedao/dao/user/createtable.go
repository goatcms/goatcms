package dao

import (
	"database/sql"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

// UserCreateTable is a Data Access Object for user entity
type UserCreateTable struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewUserCreateTable(di dependency.Injector) (*UserCreateTable, error) {
	instance := &UserCreateTable{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func UserCreateTableFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewUserCreateTable(dp)
	if err != nil {
		return nil, err
	}
	return maindef.CreateTable(instance), nil
}

func (dao UserCreateTable) CreateTable(scope app.Scope) (err error) {
	var tx *sql.Tx
	if tx, err = helpers.TX(scope, dao.deps.DB); err != nil {
		return err
	}
	_, err = tx.Exec(dao.SQL())
	return err
}

func (dao UserCreateTable) SQL() string {
	return `CREATE TABLE IF NOT EXISTS User (ID INTEGER PRIMARY KEY, Password TEXT, Firstname TEXT, Login TEXT, Email TEXT)`
}
