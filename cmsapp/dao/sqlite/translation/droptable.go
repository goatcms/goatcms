package translationdao

import (
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	sqlitebase "github.com/goatcms/goatcms/cmsapp/dao/sqlite"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/jmoiron/sqlx"
)

// TranslationDropTable is a Data Access Object for translation entity
type TranslationDropTable struct {
	deps struct {
		DB *sqlx.DB `dependency:"sqlitedb"`
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

func (dao TranslationDropTable) DropTable(scope app.Scope) error {
	var (
		err error
		tx  *sqlx.Tx
	)
	if tx, err = sqlitebase.TX(scope, dao.deps.DB); err != nil {
		return err
	}
	tx.MustExec(dao.SQL())
	return nil
}

func (dao TranslationDropTable) SQL() string {
	return `DROP TABLE IF EXISTS Translation `
}
