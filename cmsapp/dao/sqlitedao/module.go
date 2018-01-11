package sqlitedao

import (
	fragment "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/dao/fragment"
	fragmentq "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/dao/fragment/queries"
	session "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/dao/session"
	sessionq "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/dao/session/queries"
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
	if err := fragment.RegisterDependencies(dp); err != nil {
		return err
	}
	if err := fragmentq.RegisterDependencies(dp); err != nil {
		return err
	}
	if err := session.RegisterDependencies(dp); err != nil {
		return err
	}
	if err := sessionq.RegisterDependencies(dp); err != nil {
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
