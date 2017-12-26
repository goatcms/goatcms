package articledao

import (
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlite/helpers"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/jmoiron/sqlx"
)

// ArticleDropTable is a Data Access Object for article entity
type ArticleDropTable struct {
	deps struct {
		DB *sqlx.DB `dependency:"sqlitedb"`
	}
}

func NewArticleDropTable(di dependency.Injector) (*ArticleDropTable, error) {
	instance := &ArticleDropTable{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func ArticleDropTableFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewArticleDropTable(dp)
	if err != nil {
		return nil, err
	}
	return maindef.DropTable(instance), nil
}

func (dao ArticleDropTable) DropTable(scope app.Scope) error {
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

func (dao ArticleDropTable) SQL() string {
	return `DROP TABLE IF EXISTS Article `
}
