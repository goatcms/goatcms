package cmsapp

import (
	"github.com/goatcms/goatcms/cmsapp/controllers/adminctrl"
	"github.com/goatcms/goatcms/cmsapp/controllers/custom/homectrl"
	"github.com/goatcms/goatcms/cmsapp/controllers/custom/userctrl"
	"github.com/goatcms/goatcore/app"
)

func InitControllers(a app.App) error {
	if err := adminctrl.InitDependencies(a); err != nil {
		return err
	}
	if err := userctrl.InitDependencies(a); err != nil {
		return err
	}
	if err := homectrl.InitDependencies(a); err != nil {
		return err
	}
	return nil
}
