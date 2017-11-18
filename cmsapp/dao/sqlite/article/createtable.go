package articledao

import (
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	sqlitebase "github.com/goatcms/goatcms/cmsapp/dao/sqlite"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/varutil"
)

// ArticleCreateTable is a Data Access Object for article entity
type ArticleCreateTable struct {
	deps struct {
		DB *sql.DB `dependency:"sqlitedb"`
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
		tx  *sql.Tx
	)
	if tx, err = sqlitebase.TX(scope, dao.deps.DB); err != nil {
		return err
	}
	tx.MustExec(dao.SQL())
	return nil
}

func (dao ArticleCreateTable) SQL() string {
	return `CREATE TABLE Article (Title TEXT, Content TEXT)`
}
