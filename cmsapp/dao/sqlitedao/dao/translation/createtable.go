package dao

import (
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/jmoiron/sqlx"
)

// TranslationCreateTable is a Data Access Object for translation entity
type TranslationCreateTable struct {
	deps struct {
		DB *sqlx.DB `dependency:"db0.engine"`
	}
}

func NewTranslationCreateTable(di dependency.Injector) (*TranslationCreateTable, error) {
	instance := &TranslationCreateTable{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func TranslationCreateTableFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewTranslationCreateTable(dp)
	if err != nil {
		return nil, err
	}
	return maindef.CreateTable(instance), nil
}

func (dao TranslationCreateTable) CreateTable(scope app.Scope) error {
	var (
		err error
		tx  *sqlx.Tx
	)
	if tx, err = helpers.TX(scope, dao.deps.DB); err != nil {
		return err
	}
	_, err = tx.Exec(dao.SQL())
	return err
}

func (dao TranslationCreateTable) SQL() string {
	return `CREATE TABLE Translation (ID INTEGER PRIMARY KEY, Value TEXT, Key TEXT)`
}
