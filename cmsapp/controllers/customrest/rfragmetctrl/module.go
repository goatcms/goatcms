package ruserctrl

import (
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
)

// InitDependencies init all dependency modules
func InitDependencies(a app.App) (err error) {
	var (
		deps struct {
			Router services.Router `dependency:"RouterService"`
		}
		setCtrl *SetCtrl
	)
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	// status
	if status, err = NewStatus(a.DependencyProvider()); err != nil {
		return err
	}
	deps.Router.OnGet("/rest/auth/status", status.DO)
	return nil
}
