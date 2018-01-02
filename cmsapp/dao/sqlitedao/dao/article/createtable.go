package dao

import (
	"database/sql"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

// ArticleCreateTable is a Data Access Object for article entity
type ArticleCreateTable struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewArticleCreateTable(di dependency.Injector) (*ArticleCreateTable, error) {
	instance := &ArticleCreateTable{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func ArticleCreateTableFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewArticleCreateTable(dp)
	if err != nil {
		return nil, err
	}
	return maindef.CreateTable(instance), nil
}

func (dao ArticleCreateTable) CreateTable(scope app.Scope) (err error) {
	var tx *sql.Tx
	if tx, err = helpers.TX(scope, dao.deps.DB); err != nil {
		return err
	}
	_, err = tx.Exec(dao.SQL())
	return err
}

func (dao ArticleCreateTable) SQL() string {
	return `CREATE TABLE IF NOT EXISTS Article (ID INTEGER PRIMARY KEY, Title TEXT, Content TEXT)`
}
