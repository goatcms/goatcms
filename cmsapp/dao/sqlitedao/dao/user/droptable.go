package dao

import (
	"database/sql"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

// UserDropTable is a Data Access Object for user entity
type UserDropTable struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewUserDropTable(di dependency.Injector) (*UserDropTable, error) {
	instance := &UserDropTable{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func UserDropTableFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewUserDropTable(dp)
	if err != nil {
		return nil, err
	}
	return maindef.DropTable(instance), nil
}

func (dao UserDropTable) DropTable(scope app.Scope) (err error) {
	var tx *sql.Tx
	if tx, err = helpers.TX(scope, dao.deps.DB); err != nil {
		return err
	}
	_, err = tx.Exec(dao.SQL())
	return err
}

func (dao UserDropTable) SQL() string {
	return `DROP TABLE IF EXISTS User `
}
