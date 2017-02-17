package articlectrl

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/goatcms/goat-core/app"
	"github.com/goatcms/goat-core/db"
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goat-core/goathtml"
	"github.com/goatcms/goatcms/cmsapp/models"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/gorilla/mux"
)

// ViewCtrl is a controler to show a single article
type ViewCtrl struct {
	deps struct {
		Template services.Template `dependency:"TemplateService"`
		FindByID db.FindByID       `dependency:"db.query.article.FindByID"`
	}
	view *template.Template
}

// NewViewCtrl create instance of a list articles controller
func NewViewCtrl(dp dependency.Provider) (*ViewCtrl, error) {
	var err error
	ctrl := &ViewCtrl{}
	if err = dp.InjectTo(&ctrl.deps); err != nil {
		return nil, err
	}
	ctrl.view, err = ctrl.deps.Template.View(goathtml.DefaultLayout, "articles/view", nil)
	return ctrl, nil
}

// Get is handler to serve template to view a article
func (c *ViewCtrl) Get(requestScope app.Scope) {
	var (
		tx          db.TX
		err         error
		requestDeps struct {
			RequestDB    services.RequestDB    `request:"RequestDBService"`
			RequestError services.RequestError `request:"RequestErrorService"`
			Request      *http.Request         `request:"Request"`
			Response     http.ResponseWriter   `request:"Response"`
		}
	)
	if err = requestScope.InjectTo(&requestDeps); err != nil {
		fmt.Println(err)
		return
	}
	vars := mux.Vars(requestDeps.Request)
	if vars["id"] == "" {
		requestDeps.RequestError.Errorf(312, "id is required")
		return
	}
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
	if tx, err = requestDeps.RequestDB.TX(); err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
	row, err := c.deps.FindByID(tx, id)
	if err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
	article := &models.Article{}
	row.StructScan(article)
	err = c.view.Execute(requestDeps.Response, article)
	if err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
}
