package articlectrl

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/goatcms/goat-core/app"
	"github.com/goatcms/goat-core/db"
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goat-core/goathtml"
	"github.com/goatcms/goat-core/messages/msgcollection"
	"github.com/goatcms/goatcms/cmsapp/forms/article/articleform"
	"github.com/goatcms/goatcms/cmsapp/services"
)

// InsertCtrl is a controler to create new article
type InsertCtrl struct {
	deps struct {
		Template    services.Template `dependency:"TemplateService"`
		InsertQuery db.Insert         `dependency:"db.query.article.Insert"`
	}
	view *template.Template
}

// NewInsertCtrl create instance of a insert articles controller
func NewInsertCtrl(dp dependency.Provider) (*InsertCtrl, error) {
	var err error
	ctrl := &InsertCtrl{}
	if err = dp.InjectTo(&ctrl.deps); err != nil {
		return nil, err
	}
	ctrl.view, err = ctrl.deps.Template.View(goathtml.DefaultLayout, "articles/insert", nil)
	if err != nil {
		return nil, err
	}
	return ctrl, nil
}

// Get is handler to serve template where one can add new article
func (c *InsertCtrl) Get(requestScope app.Scope) {
	var requestDeps struct {
		RequestError services.RequestError `request:"RequestErrorService"`
		Response     http.ResponseWriter   `request:"Response"`
	}
	if err := requestScope.InjectTo(&requestDeps); err != nil {
		fmt.Println(err)
		return
	}
	if err := c.view.Execute(requestDeps.Response, map[string]interface{}{
		"Valid": msgcollection.NewMessageMap(),
	}); err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
}

// Post is handler to save article from form obtained data
func (c *InsertCtrl) Post(requestScope app.Scope) {
	var (
		tx          db.TX
		requestDeps struct {
			RequestDB    services.RequestDB    `request:"RequestDBService"`
			RequestError services.RequestError `request:"RequestErrorService"`
			Request      *http.Request         `request:"Request"`
			Response     http.ResponseWriter   `request:"Response"`
		}
	)
	if err := requestScope.InjectTo(&requestDeps); err != nil {
		fmt.Println(err)
		return
	}
	form, err := articleform.NewForm(requestScope)
	if err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
	if err = requestScope.InjectTo(form); err != nil {
		fmt.Println(err)
		return
	}
	validResult := msgcollection.NewMessageMap()
	if err = form.Valid("", validResult); err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
	if len(validResult.GetAll()) != 0 {
		if err = c.view.Execute(requestDeps.Response, map[string]interface{}{
			"Valid": validResult,
		}); err != nil {
			requestDeps.RequestError.Error(312, err)
			return
		}
		return
	}
	if tx, err = requestDeps.RequestDB.TX(); err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
	if _, err = c.deps.InsertQuery(tx, form); err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
	if err = tx.Commit(); err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
	http.Redirect(requestDeps.Response, requestDeps.Request, ListURL, http.StatusSeeOther)
}
