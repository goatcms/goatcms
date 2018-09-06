package fragments

import (
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

// RegisterDependencies is init callback to register module dependencies
func RegisterDependencies(dp dependency.Provider) error {
	if err := dp.AddDefaultFactory("FragmentStorage", StorageFactory); err != nil {
		return err
	}
	return dp.AddDefaultFactory("FragmentTemplateHelper", TemplateHelperFactory)
}

// InitDependencies is init callback to inject dependencies inside module
func InitDependencies(a app.App) (err error) {
	var deps struct {
		Template       services.Template               `dependency:"TemplateService"`
		TemplateHelper services.FragmentTemplateHelper `dependency:"FragmentTemplateHelper"`
	}
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	return deps.Template.AddFunc("Fragment", deps.TemplateHelper.RenderFragment)
}
