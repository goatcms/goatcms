package dao

import (
	"database/sql"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

// SessionDropTable is a Data Access Object for session entity
type SessionDropTable struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewSessionDropTable(di dependency.Injector) (*SessionDropTable, error) {
	instance := &SessionDropTable{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func SessionDropTableFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewSessionDropTable(dp)
	if err != nil {
		return nil, err
	}
	return maindef.DropTable(instance), nil
}

func (dao SessionDropTable) DropTable(scope app.Scope) (err error) {
	var tx *sql.Tx
	if tx, err = helpers.TX(scope, dao.deps.DB); err != nil {
		return err
	}
	_, err = tx.Exec(dao.SQL())
	return err
}

func (dao SessionDropTable) SQL() string {
	return `DROP TABLE IF EXISTS Session `
}
