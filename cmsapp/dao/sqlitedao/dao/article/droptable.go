package dao

import (
	"database/sql"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

// ArticleDropTable is a Data Access Object for article entity
type ArticleDropTable struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
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

func (dao ArticleDropTable) DropTable(scope app.Scope) (err error) {
	var tx *sql.Tx
	if tx, err = helpers.TX(scope, dao.deps.DB); err != nil {
		return err
	}
	_, err = tx.Exec(dao.SQL())
	return err
}

func (dao ArticleDropTable) SQL() string {
	return `DROP TABLE IF EXISTS Article `
}
