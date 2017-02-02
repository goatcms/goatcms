package reqsession

import (
	"github.com/goatcms/goat-core/app"
	"github.com/goatcms/goatcms/cmsapp/services"
)

// InitDependencies is init callback to inject dependencies inside module
func InitDependencies(a app.App) error {
	routerIns, err := a.DependencyProvider().Get(services.RouterService)
	if err != nil {
		return err
	}
	router := routerIns.(services.Router)
	if err := router.AddFactory(services.SessionManagerService, SessionManagerFactory); err != nil {
		return err
	}
	return nil
}
