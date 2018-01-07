package user

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcms/cmsapp/services"
)

// InitDependencies initialize the User controllers
func InitDependencies(a app.App) (err error) {
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
	deps.Router.OnGet("/admin/user", list.Get)
	// add insert controller
	insert, err := NewInsert(dp)
	if err != nil {
		return err
	}
	deps.Router.OnGet("/admin/user/insert", insert.Get)
	deps.Router.OnPost("/admin/user/insert", insert.Post)
	return nil
}