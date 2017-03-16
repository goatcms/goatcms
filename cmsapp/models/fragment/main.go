package user

import (
	"reflect"

	"github.com/goatcms/goatcms/cmsapp/models"
	"github.com/goatcms/goatcore/db"
	"github.com/goatcms/goatcore/db/orm"
	"github.com/goatcms/goatcore/dependency"
)

const (
	FragmentTable = "Fragments"
)

func RegisterDependencies(dp dependency.Provider, dsql db.DSQL) error {
	var entityPtr *models.User
	table := orm.NewTable(FragmentTable, reflect.TypeOf(entityPtr).Elem())
	dp.Set("FragmentTable", table)

	insert, err := orm.NewInsert(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("FragmentInsert", insert)

	insertWithID, err := orm.NewInsertWithID(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("FragmentInsertWithID", insertWithID)

	update, err := orm.NewUpdate(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("FragmentUpdate", update)

	delete, err := orm.NewDelete(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("FragmentDelete", delete)

	findAll, err := orm.NewFindAll(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("FragmentFindAll", findAll)

	findByID, err := orm.NewFindByID(table, dsql)
	if err != nil {
		return err
	}
	dp.SetDefault("FragmentFindByID", findByID)

	return nil
}
