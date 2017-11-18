package userdao

import (
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	sqlitebase "github.com/goatcms/goatcms/cmsapp/dao/sqlite"
	"github.com/goatcms/goatcore/app"
)

// UserDropTable is a Data Access Object for user entity
type UserDropTable struct {
	deps struct {
		DB *sql.DB `dependency:"sqlitedb"`
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

func (dao UserDropTable) DropTable(scope app.Scope) error {
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

func (dao UserDropTable) SQL() string {
	return `DROP TABLE IF EXISTS User `
}
