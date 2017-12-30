package adminctrl

import (
	"github.com/goatcms/goatcore/app"
	user "github.com/goatcms/goatcms/cmsapp/controllers/adminctrl/model/user"
	translation "github.com/goatcms/goatcms/cmsapp/controllers/adminctrl/model/translation"
	article "github.com/goatcms/goatcms/cmsapp/controllers/adminctrl/model/article")

func InitDependencies(a app.App) error {
	if err := translation.InitDependencies(a); err != nil {
		return err
	}
	if err := user.InitDependencies(a); err != nil {
		return err
	}
	if err := article.InitDependencies(a); err != nil {
		return err
	}
	return nil
}