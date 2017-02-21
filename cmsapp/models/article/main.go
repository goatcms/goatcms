package article

import (
	"reflect"

	"github.com/goatcms/goatcore/db"
	"github.com/goatcms/goatcore/db/orm"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcms/cmsapp/models"
)

const (
	ArticleTable = "Articles"
)

func RegisterDependencies(dp dependency.Provider, dsql db.DSQL) error {
	var entityPtr *models.Article
	table := orm.NewTable(ArticleTable, reflect.TypeOf(entityPtr).Elem())

	createTableQuery, err := orm.NewCreateTable(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("db.query.article.CreateTable", createTableQuery)

	insertQuery, err := orm.NewInsert(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("db.query.article.Insert", insertQuery)

	insertWithIDQuery, err := orm.NewInsertWithID(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("db.query.article.InsertWithID", insertWithIDQuery)

	updateQuery, err := orm.NewUpdate(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("db.query.article.Update", updateQuery)

	deleteQuery, err := orm.NewDelete(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("db.query.article.Delete", deleteQuery)

	findAllQuery, err := orm.NewFindAll(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("db.query.article.FindAll", findAllQuery)

	findByIDQuery, err := orm.NewFindByID(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("db.query.article.FindByID", findByIDQuery)

	return nil
}
