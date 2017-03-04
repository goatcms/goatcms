package reqsession

import (
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/app"
)

// InitDependencies is init callback to inject dependencies inside module
func InitDependencies(a app.App) error {
	routerIns, err := a.DependencyProvider().Get(services.RouterService)
	if err != nil {
		return err
	}
	router := routerIns.(services.Router)
	if err := router.AddFactory(requestdep.SessionService, SessionFactory); err != nil {
		return err
	}
	return nil
}
