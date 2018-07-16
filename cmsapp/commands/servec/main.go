package servec

import (
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
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
