package sqlitedao

import (
	article "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/dao/article"
	articleq "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/dao/article/queries"
	translation "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/dao/translation"
	translationq "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/dao/translation/queries"
	user "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/dao/user"
	userq "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/dao/user/queries"
	database "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/database"
	"github.com/goatcms/goatcore/dependency"
)

func RegisterDependencies(dp dependency.Provider) error {
	if err := dp.AddDefaultFactory("db0.engine", database.EngineFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("db0", database.Factory); err != nil {
		return err
	}
	if err := article.RegisterDependencies(dp); err != nil {
		return err
	}
	if err := articleq.RegisterDependencies(dp); err != nil {
		return err
	}
	if err := translation.RegisterDependencies(dp); err != nil {
		return err
	}
	if err := translationq.RegisterDependencies(dp); err != nil {
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
