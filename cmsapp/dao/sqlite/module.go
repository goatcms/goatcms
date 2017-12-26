package sqlitedao

import (
	article "github.com/goatcms/goatcms/cmsapp/dao/sqlite/dao/article"
	translation "github.com/goatcms/goatcms/cmsapp/dao/sqlite/dao/translation"
	user "github.com/goatcms/goatcms/cmsapp/dao/sqlite/dao/user"
	"github.com/goatcms/goatcore/dependency"
)

func RegisterDependencies(dp dependency.Provider) error {
	if err := article.RegisterDependencies(dp); err != nil {
		return err
	}
	if err := user.RegisterDependencies(dp); err != nil {
		return err
	}
	if err := translation.RegisterDependencies(dp); err != nil {
		return err
	}
	return nil
}
