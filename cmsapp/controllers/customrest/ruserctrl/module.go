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
		status  *Status
		signup  *Signup
		signin  *Signin
		signout *Signout
	)
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	// status
	if status, err = NewStatus(a.DependencyProvider()); err != nil {
		return err
	}
	deps.Router.OnGet("/rest/auth/status", status.DO)
	// signup
	if signup, err = NewSignup(a.DependencyProvider()); err != nil {
		return err
	}
	deps.Router.OnPost("/rest/auth/signup", signup.DO)
	// signin
	if signin, err = NewSignin(a.DependencyProvider()); err != nil {
		return err
	}
	deps.Router.OnPost("/rest/auth/signin", signin.DO)
	// signout
	if signout, err = NewSignout(a.DependencyProvider()); err != nil {
		return err
	}
	deps.Router.OnPost("/rest/auth/signout", signout.DO)
	return nil
}
