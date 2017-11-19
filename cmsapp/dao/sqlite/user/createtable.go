package userdao

import (
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	sqlitebase "github.com/goatcms/goatcms/cmsapp/dao/sqlite"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/varutil"
)

// UserCreateTable is a Data Access Object for user entity
type UserCreateTable struct {
	deps struct {
		DB *sql.DB `dependency:"sqlitedb"`
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

func (dao UserCreateTable) CreateTable(scope app.Scope) error {
	var (
		err error
		tx  *sql.Tx
	)
	if tx, err = sqlitebase.TX(scope, dao.deps.DB); err != nil {
		return err
	}
	tx.MustExec(dao.SQL())
	return nil
}

func (dao UserCreateTable) SQL() string {
	return `CREATE TABLE User (Lastname TEXT, Firstname TEXT)`
}
