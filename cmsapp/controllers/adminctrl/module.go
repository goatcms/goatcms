package adminctrl

import (
	"github.com/goatcms/goatcore/app"
	fragment "github.com/goatcms/goatcms/cmsapp/controllers/adminctrl/model/fragment"
	user "github.com/goatcms/goatcms/cmsapp/controllers/adminctrl/model/user")

func InitDependencies(a app.App) error {
	if err := fragment.InitDependencies(a); err != nil {
		return err
	}
	if err := user.InitDependencies(a); err != nil {
		return err
	}
	return nil
}