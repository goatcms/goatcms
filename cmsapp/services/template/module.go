package template

import (
	"github.com/goatcms/goat-core/app"
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/cmsapp/services"
)

// RegisterDependencies is init callback to register module dependencies
func RegisterDependencies(dp dependency.Provider) error {
	if err := dp.AddDefaultFactory(services.TemplateService, TemplateProviderFactory); err != nil {
		return err
	}
	return nil
}

// InitDependencies is init callback to inject dependencies inside module
func InitDependencies(a app.App) error {
	ins, err := a.DependencyProvider().Get(services.TemplateService)
	if err != nil {
		return err
	}
	templateProvider := ins.(services.Template)
	AddDefaultTemplateFunctions(templateProvider)
	return nil
}
