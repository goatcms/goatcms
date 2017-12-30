package cmsapp

import (
	"github.com/goatcms/goatcms/cmsapp/controllers/adminctrl"
	"github.com/goatcms/goatcms/cmsapp/controllers/homectrl"
	"github.com/goatcms/goatcore/app"
)

func InitControllers(a app.App) error {
	if err := adminctrl.InitDependencies(a); err != nil {
		return err
	}
	if err := homectrl.InitDependencies(a); err != nil {
		return err
	}
	return nil
}
