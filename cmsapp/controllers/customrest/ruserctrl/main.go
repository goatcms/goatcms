package ruserctrl

import (
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
)

// InitDependencies init all dependency modules
func InitDependencies(a app.App) error {
	var deps struct {
		Router services.Router `dependency:"RouterService"`
	}
	if err := a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	// status
	status, err := NewStatus(a.DependencyProvider())
	if err != nil {
		return err
	}
	deps.Router.OnGet("/rest/auth/status", status.DO)
	// signup
	signup, err := NewSignup(a.DependencyProvider())
	if err != nil {
		return err
	}
	deps.Router.OnPost("/rest/auth/signup", signup.DO)
	// signin
	signin, err := NewSignin(a.DependencyProvider())
	if err != nil {
		return err
	}
	deps.Router.OnPost("/rest/auth/signin", signin.DO)
	// signout
	signout, err := NewSignout(a.DependencyProvider())
	if err != nil {
		return err
	}
	deps.Router.OnPost("/rest/auth/signout", signout.DO)
	return nil
}
