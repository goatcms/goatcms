package article

import (
	"reflect"

	"github.com/goatcms/goatcms/cmsapp/models"
	"github.com/goatcms/goatcore/db"
	"github.com/goatcms/goatcore/db/orm"
	"github.com/goatcms/goatcore/dependency"
)

const (
	ArticleTable = "Articles"
)

func RegisterDependencies(dp dependency.Provider, dsql db.DSQL) error {
	var entityPtr *models.Article
	table := orm.NewTable(ArticleTable, reflect.TypeOf(entityPtr).Elem())

	createTable, err := orm.NewCreateTable(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("ArticleCreateTable", createTable)

	insert, err := orm.NewInsert(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("ArticleInsert", insert)

	insertWithID, err := orm.NewInsertWithID(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("ArticleInsertWithID", insertWithID)

	update, err := orm.NewUpdate(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("ArticleUpdate", update)

	delete, err := orm.NewDelete(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("ArticleDelete", delete)

	findAll, err := orm.NewFindAll(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("ArticleFindAll", findAll)

	findByID, err := orm.NewFindByID(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("ArticleFindByID", findByID)

	return nil
}
