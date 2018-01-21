package homectrl

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcms/cmsapp/controllers"
	"github.com/goatcms/goatcms/cmsapp/services"
)

// InitDependencies init all dependency modules
func InitDependencies(a app.App) error {
	var deps struct {
		Router services.Router `dependency:"RouterService"`
	}
	if err := a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	home, err := NewHome(a.DependencyProvider())
	if err != nil {
		return err
	}
	deps.Router.OnGet(controllers.HomeUrl, home.Get)
	return nil
}
