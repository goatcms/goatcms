package articledao

import (
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlite/helpers"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/jmoiron/sqlx"
)

// ArticleCreateTable is a Data Access Object for article entity
type ArticleCreateTable struct {
	deps struct {
		DB *sqlx.DB `dependency:"sqlitedb"`
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

func (dao ArticleCreateTable) CreateTable(scope app.Scope) error {
	var (
		err error
		tx  *sqlx.Tx
	)
	if tx, err = helpers.TX(scope, dao.deps.DB); err != nil {
		return err
	}
	tx.MustExec(dao.SQL())
	return nil
}

func (dao ArticleCreateTable) SQL() string {
	return `CREATE TABLE Article (ID INTEGER PRIMARY KEY, Content TEXT, Title TEXT)`
}
