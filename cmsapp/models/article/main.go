package user

import (
	"github.com/goatcms/goat-core/db"
	"github.com/goatcms/goat-core/db/orm"
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/simpletype"
)

const (
	ArticleTable = "Articles"
)

func RegisterDependencies(dp dependency.Provider, dsql db.DSQL) error {
	table := orm.NewTable(ArticleTable, map[string]types.CustomType{
		"Title":   simpletype.NewTitleType(map[string]string{types.Required: "true"}),
		"Content": simpletype.NewContentType(map[string]string{types.Required: "true"}),
	})

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
