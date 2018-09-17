package cmsapp

import (
	"github.com/goatcms/goatcms/cmsapp/controllers/adminctrl"
	"github.com/goatcms/goatcms/cmsapp/controllers/custom/oauthctrl"
	"github.com/goatcms/goatcms/cmsapp/controllers/custom/userctrl"
	"github.com/goatcms/goatcms/cmsapp/controllers/customrest/ruserctrl"
	"github.com/goatcms/goatcms/cmsapp/controllers/restctrl"
	"github.com/goatcms/goatcore/app"
)

// InitControllers add cmsapp controllers to an application
func InitControllers(a app.App) (err error) {
	// webpages
	if err = adminctrl.InitDependencies(a); err != nil {
		return err
	}
	if err = userctrl.InitDependencies(a); err != nil {
		return err
	}
	if err = oauthctrl.InitDependencies(a); err != nil {
		return err
	}
	// restapi
	if err = restctrl.InitDependencies(a); err != nil {
		return err
	}
	if err = ruserctrl.InitDependencies(a); err != nil {
		return err
	}
	return nil
}
