package userctrl

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
	// signup
	signup, err := NewSignup(a.DependencyProvider())
	if err != nil {
		return err
	}
	deps.Router.OnGet("/user/signup", signup.Get)
	deps.Router.OnPost("/user/signup", signup.Post)
	deps.Router.OnGet("/user/register", signup.Get)
	deps.Router.OnPost("/user/register", signup.Post)
	// signin
	signin, err := NewSignin(a.DependencyProvider())
	if err != nil {
		return err
	}
	deps.Router.OnGet("/user/signin", signin.Get)
	deps.Router.OnPost("/user/signin", signin.Post)
	deps.Router.OnGet("/user/login", signin.Get)
	deps.Router.OnPost("/user/login", signin.Post)
	// signout
	signout, err := NewSignout(a.DependencyProvider())
	if err != nil {
		return err
	}
	deps.Router.OnGet("/user/signout", signout.Do)
	deps.Router.OnPost("/user/signout", signout.Do)
	deps.Router.OnGet("/user/logout", signout.Do)
	deps.Router.OnPost("/user/logout", signout.Do)
	return nil
}
