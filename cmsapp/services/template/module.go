package template

import (
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcms/cmsapp/services/template/templatex"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

// RegisterDependencies is init callback to register module dependencies
func RegisterDependencies(dp dependency.Provider) error {
	if err := dp.AddDefaultFactory(services.TemplateService, TemplateProviderFactory); err != nil {
		return err
	}
	return nil
}

// InitDependencies is init callback to inject dependencies inside module
func InitDependencies(a app.App) (err error) {
	var xmodule *templatex.Module
	if xmodule, err = templatex.NewModule(a.DependencyProvider()); err != nil {
		return err
	}
	if err = xmodule.Register(); err != nil {
		return err
	}
	return nil
}
