package user

import (
	"github.com/goatcms/goat-core/db"
	"github.com/goatcms/goat-core/db/orm"
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/simpletype"
)

const (
	FragmentTable = "Fragments"
)

func RegisterDependencies(dp dependency.Provider, dsql db.DSQL) error {
	table := orm.NewTable(FragmentTable, map[string]types.CustomType{
		"Key":   simpletype.NewTitleType(map[string]string{types.Required: "true"}),
		"Value": simpletype.NewContentType(map[string]string{types.Required: "true"}),
	})

	createTableQuery, err := orm.NewCreateTable(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("db.query.fragment.CreateTable", createTableQuery)

	insertQuery, err := orm.NewInsert(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("db.query.fragment.Insert", insertQuery)

	insertWithIDQuery, err := orm.NewInsertWithID(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("db.query.fragment.InsertWithID", insertWithIDQuery)

	updateQuery, err := orm.NewUpdate(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("db.query.fragment.Update", updateQuery)

	deleteQuery, err := orm.NewDelete(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("db.query.fragment.Delete", deleteQuery)

	findAllQuery, err := orm.NewFindAll(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("db.query.fragment.FindAll", findAllQuery)

	findByIDQuery, err := orm.NewFindByID(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("db.query.fragment.FindByID", findByIDQuery)

	return nil
}
