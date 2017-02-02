package reqauth

import (
	"github.com/goatcms/goat-core/app"
	"github.com/goatcms/goatcms/cmsapp/services"
)

// InitDependencies is init callback to register module dependencies
func InitDependencies(a app.App) error {
	var deps struct {
		Router services.Router `dependency:"RouterService"`
	}
	if err := a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	if err := deps.Router.AddFactory(services.RequestAuthService, RequestAuthFactory); err != nil {
		return err
	}
	return nil
}
