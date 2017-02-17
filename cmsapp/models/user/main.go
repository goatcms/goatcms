package user

import (
	"reflect"

	"github.com/goatcms/goat-core/db"
	"github.com/goatcms/goat-core/db/orm"
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/cmsapp/models"
)

const (
	UserTable = "Users"
)

func RegisterDependencies(dp dependency.Provider, dsql db.DSQL) error {
	var entityPtr *models.User
	table := orm.NewTable(UserTable, reflect.TypeOf(entityPtr).Elem())

	loginQuery, err := NewLogin(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("db.query.user.LoginQuery", loginQuery)

	createTableQuery, err := orm.NewCreateTable(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("db.query.user.CreateTable", createTableQuery)

	insertQuery, err := orm.NewInsert(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("db.query.user.Insert", insertQuery)

	insertWithIDQuery, err := orm.NewInsertWithID(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("db.query.user.InsertWithID", insertWithIDQuery)

	updateQuery, err := orm.NewUpdate(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("db.query.user.Update", updateQuery)

	deleteQuery, err := orm.NewDelete(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("db.query.user.Delete", deleteQuery)

	findAllQuery, err := orm.NewFindAll(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("db.query.user.FindAll", findAllQuery)

	findByIDQuery, err := orm.NewFindByID(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("db.query.user.FindByID", findByIDQuery)

	//extended queries
	dp.AddDefaultFactory("db.query.user.RegisterQuery", RegisterQueryFactory)

	return nil
}
