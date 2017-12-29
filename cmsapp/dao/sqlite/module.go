package dao

import (
	article "github.com/goatcms/goatcms/cmsapp/dao/sqlite/dao/article"
	articleq "github.com/goatcms/goatcms/cmsapp/dao/sqlite/dao/article/queries"
	translation "github.com/goatcms/goatcms/cmsapp/dao/sqlite/dao/translation"
	translationq "github.com/goatcms/goatcms/cmsapp/dao/sqlite/dao/translation/queries"
	user "github.com/goatcms/goatcms/cmsapp/dao/sqlite/dao/user"
	userq "github.com/goatcms/goatcms/cmsapp/dao/sqlite/dao/user/queries"
	database "github.com/goatcms/goatcms/cmsapp/dao/sqlite/database"
	"github.com/goatcms/goatcore/dependency"
)

func RegisterDependencies(dp dependency.Provider) error {
	if err := dp.AddDefaultFactory("sqlitedb0", database.Factory); err != nil {
		return err
	}
	if err := translation.RegisterDependencies(dp); err != nil {
		return err
	}
	if err := translationq.RegisterDependencies(dp); err != nil {
		return err
	}
	if err := article.RegisterDependencies(dp); err != nil {
		return err
	}
	if err := articleq.RegisterDependencies(dp); err != nil {
		return err
	}
	if err := user.RegisterDependencies(dp); err != nil {
		return err
	}
	if err := userq.RegisterDependencies(dp); err != nil {
		return err
	}
	return nil
}
