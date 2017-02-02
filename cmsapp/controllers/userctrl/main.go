package userctrl

import (
	"github.com/goatcms/goat-core/app"
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
	register, err := NewUserRegisterController(a.DependencyProvider())
	if err != nil {
		return err
	}
	login, err := NewUserLoginController(a.DependencyProvider())
	if err != nil {
		return err
	}
	logout, err := NewUserLogoutController(a.DependencyProvider())
	if err != nil {
		return err
	}
	deps.Router.OnGet("/register", register.Get)
	deps.Router.OnPost("/register", register.Post)
	deps.Router.OnGet("/login", login.Get)
	deps.Router.OnPost("/login", login.Post)
	deps.Router.OnGet("/logout", logout.All)
	deps.Router.OnPost("/logout", logout.All)
	return nil
}
