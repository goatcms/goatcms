package dao

import (
	"database/sql"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

// TranslationDropTable is a Data Access Object for translation entity
type TranslationDropTable struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewTranslationDropTable(di dependency.Injector) (*TranslationDropTable, error) {
	instance := &TranslationDropTable{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func TranslationDropTableFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewTranslationDropTable(dp)
	if err != nil {
		return nil, err
	}
	return maindef.DropTable(instance), nil
}

func (dao TranslationDropTable) DropTable(scope app.Scope) (err error) {
	var tx *sql.Tx
	if tx, err = helpers.TX(scope, dao.deps.DB); err != nil {
		return err
	}
	_, err = tx.Exec(dao.SQL())
	return err
}

func (dao TranslationDropTable) SQL() string {
	return `DROP TABLE IF EXISTS Translation `
}
