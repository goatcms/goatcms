package reqauth

import (
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/app"
)

// InitDependencies is init callback to register module dependencies
func InitDependencies(a app.App) error {
	var deps struct {
		Router services.Router `dependency:"RouterService"`
	}
	if err := a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	if err := deps.Router.AddFactory(requestdep.AuthService, AuthFactory); err != nil {
		return err
	}
	return nil
}
