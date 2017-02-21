package servec

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcms/cmsapp/services"
)

// Run run command
func Run(app app.App) error {
	var deps struct {
		Router services.Router `dependency:"RouterService"`
	}
	if err := app.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	return deps.Router.Start()
}
