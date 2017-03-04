package user

import (
	"reflect"

	"github.com/goatcms/goatcms/cmsapp/models"
	"github.com/goatcms/goatcore/db"
	"github.com/goatcms/goatcore/db/orm"
	"github.com/goatcms/goatcore/dependency"
)

const (
	UserTable = "Users"
)

func RegisterDependencies(dp dependency.Provider, dsql db.DSQL) error {
	var entityPtr *models.User
	table := orm.NewTable(UserTable, reflect.TypeOf(entityPtr).Elem())

	//extended queries
	dp.AddDefaultFactory("UserRegister", RegisterFactory)

	login, err := NewLogin(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("UserLogin", login)

	createTable, err := orm.NewCreateTable(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("UserCreateTable", createTable)

	insert, err := orm.NewInsert(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("UserInsert", insert)

	insertWithID, err := orm.NewInsertWithID(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("UserInsertWithID", insertWithID)

	update, err := orm.NewUpdate(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("UserUpdate", update)

	delete, err := orm.NewDelete(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("UserDelete", delete)

	findAll, err := orm.NewFindAll(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("UserFindAll", findAll)

	findByID, err := orm.NewFindByID(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("UserFindByID", findByID)

	return nil
}
