package translationctrl

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/goatcms/goatcms/cmsapp/models"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/db"
	"github.com/goatcms/goatcore/db/entityChan"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/goathtml"
)

// List is a controler to show a list of article
type List struct {
	deps struct {
		Template services.Template `dependency:"TemplateService"`
		FindAll  db.FindAll        `dependency:"TranslationFindAll"`
	}
	view *template.Template
}

// NewList create instance of a list translation controller
func NewList(dp dependency.Provider) (*List, error) {
	var err error
	ctrl := &List{}
	if err = dp.InjectTo(&ctrl.deps); err != nil {
		return nil, err
	}
	if ctrl.view, err = ctrl.deps.Template.View(goathtml.DefaultLayout, "admin/translation/list", nil); err != nil {
		return nil, err
	}
	return ctrl, nil
}

// Get is handler to serve template where one can add new article
func (c *ListCtrl) Get(requestScope app.Scope) {
	var (
		rows        db.Rows
		err         error
		requestDeps struct {
			RequestError requestdep.Error     `request:"ErrorService"`
			Responser    requestdep.Responser `request:"ResponserService"`
			Request      *http.Request        `request:"Request"`
		}
	)
	if err = requestScope.InjectTo(&requestDeps); err != nil {
		fmt.Println(err)
		return
	}
	if rows, err = c.deps.FindAll(requestScope); err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
	requestScope.On(app.ErrorEvent, func(erri interface{}) error {
		scopeErr := erri.(error)
		requestDeps.RequestError.Errorf(403, "%s", scopeErr.Error())
		return nil
	})
	if err = requestDeps.Responser.Execute(c.view, rows); err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
}