package articledao

import (
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	sqlitebase "github.com/goatcms/goatcms/cmsapp/dao/sqlite"
	"github.com/goatcms/goatcore/app"
)

// ArticleDropTable is a Data Access Object for article entity
type ArticleDropTable struct {
	deps struct {
		DB *sql.DB `dependency:"sqlitedb"`
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
		tx  *sql.Tx
	)
	if tx, err = sqlitebase.TX(scope, dao.deps.DB); err != nil {
		return err
	}
	tx.MustExec(dao.SQL())
	return nil
}

func (dao ArticleDropTable) SQL() string {
	return `DROP TABLE IF EXISTS Article `
}
