package fragments

import (
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

// RegisterDependencies is init callback to register module dependencies
func RegisterDependencies(dp dependency.Provider) error {
	if err := dp.AddDefaultFactory("FragmentCache", CacheFactory); err != nil {
		return err
	}
	return nil
}

// InitDependencies is init callback to inject dependencies inside module
func InitDependencies(a app.App) (err error) {
	var deps struct {
		Template      services.Template      `dependency:"TemplateService"`
		FragmentCache services.FragmentCache `dependency:"FragmentCache"`
	}
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	return deps.Template.AddFunc("Fragment", deps.FragmentCache.RenderFragment)
}
