package dao

import (
	"database/sql"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

// FragmentDropTable is a Data Access Object for fragment entity
type FragmentDropTable struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewFragmentDropTable(di dependency.Injector) (*FragmentDropTable, error) {
	instance := &FragmentDropTable{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func FragmentDropTableFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewFragmentDropTable(dp)
	if err != nil {
		return nil, err
	}
	return maindef.DropTable(instance), nil
}

func (dao FragmentDropTable) DropTable(scope app.Scope) (err error) {
	var tx *sql.Tx
	if tx, err = helpers.TX(scope, dao.deps.DB); err != nil {
		return err
	}
	_, err = tx.Exec(dao.SQL())
	return err
}

func (dao FragmentDropTable) SQL() string {
	return `DROP TABLE IF EXISTS Fragment `
}
