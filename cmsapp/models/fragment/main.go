package user

import (
	"reflect"

	"github.com/goatcms/goat-core/db"
	"github.com/goatcms/goat-core/db/orm"
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/cmsapp/models"
)

const (
	FragmentTable = "Fragments"
)

func RegisterDependencies(dp dependency.Provider, dsql db.DSQL) error {
	var entityPtr *models.User
	table := orm.NewTable(FragmentTable, reflect.TypeOf(entityPtr).Elem())

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
