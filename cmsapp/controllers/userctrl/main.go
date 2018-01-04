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
	// register
	register, err := NewRegister(a.DependencyProvider())
	if err != nil {
		return err
	}
	deps.Router.OnGet("/user/register", register.Get)
	deps.Router.OnPost("/user/register", register.Post)
	// login
	login, err := NewLogin(a.DependencyProvider())
	if err != nil {
		return err
	}
	deps.Router.OnGet("/user/login", login.Get)
	deps.Router.OnPost("/user/login", login.Post)
	// logout
	logout, err := NewLogout(a.DependencyProvider())
	if err != nil {
		return err
	}
	deps.Router.OnGet("/user/logout", logout.Do)
	deps.Router.OnPost("/user/logout", logout.Do)
	return nil
}
