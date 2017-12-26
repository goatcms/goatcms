package articlectrl

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcms/cmsapp/services"
)

const (
	// InsertURL is url to insert page
	InsertURL = "/article/add"
	// ListURL is url to articles list page
	ListURL = "/article"
	// ViewURL is url to single article page
	ViewURL = "/article/{id:[0-9]+}"
)

// InitDependencies initialize the article controllers
func InitDependencies(a app.App) error {
	var deps struct {
		Router services.Router `dependency:"RouterService"`
	}
	dp := a.DependencyProvider()
	if err := dp.InjectTo(&deps); err != nil {
		return err
	}
	insertCtrl, err := NewInsertCtrl(dp)
	if err != nil {
		return err
	}
	listCtrl, err := NewListCtrl(dp)
	if err != nil {
		return err
	}
	viewCtrl, err := NewViewCtrl(dp)
	if err != nil {
		return err
	}
	deps.Router.OnGet(InsertURL, insertCtrl.Get)
	deps.Router.OnPost(InsertURL, insertCtrl.Post)
	deps.Router.OnGet(ListURL, listCtrl.Get)
	deps.Router.OnGet(ViewURL, viewCtrl.Get)
	return nil
}
