package articlectrl

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/goatcms/goatcms/cmsapp/forms/article/articleform"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/db"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/goathtml"
	"github.com/goatcms/goatcore/messages/msgcollection"
)

// InsertCtrl is a controler to create new article
type InsertCtrl struct {
	deps struct {
		Template services.Template `dependency:"TemplateService"`
		Insert   db.Insert         `dependency:"ArticleInsert"`
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
		RequestError requestdep.Error     `request:"ErrorService"`
		Responser    requestdep.Responser `request:"ResponserService"`
	}
	if err := requestScope.InjectTo(&requestDeps); err != nil {
		fmt.Println(err)
		return
	}
	if err := requestDeps.Responser.Execute(c.view, map[string]interface{}{
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
			RequestDB    requestdep.DB        `request:"DBService"`
			RequestError requestdep.Error     `request:"ErrorService"`
			Responser    requestdep.Responser `request:"ResponserService"`
			Request      *http.Request        `request:"Request"`
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
		requestDeps.RequestError.Error(312, err)
		return
	}
	validResult := msgcollection.NewMessageMap()
	if err = form.Valid("", validResult); err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
	if len(validResult.GetAll()) != 0 {
		// valid error
		if err = requestDeps.Responser.Execute(c.view, map[string]interface{}{
			"Valid": validResult,
		}); err != nil {
			requestDeps.RequestError.Error(312, err)
			return
		}
		return
	}
	// valid success
	if tx, err = requestDeps.RequestDB.TX(); err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
	if _, err = c.deps.Insert(tx, form); err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
	if err = tx.Commit(); err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
	requestDeps.Responser.Redirect(ListURL)
}
