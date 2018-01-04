package fragment

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcms/cmsapp/services"
)

// InitDependencies initialize the Fragment controllers
func InitDependencies(a app.App) error {
	var deps struct {
		Router services.Router `dependency:"RouterService"`
	}
	dp := a.DependencyProvider()
	if err := dp.InjectTo(&deps); err != nil {
		return err
	}
	// add list controller
	list, err := NewList(dp)
	if err != nil {
		return err
	}
	deps.Router.OnGet("/admin/fragment", list.Get)
	return nil
}