package articlectrl

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/db"
	"github.com/goatcms/goatcore/db/entitychan"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/goathtml"
	"github.com/goatcms/goatcms/cmsapp/models"
	"github.com/goatcms/goatcms/cmsapp/services"
)

// ListCtrl is a controler to show a list of article
type ListCtrl struct {
	deps struct {
		Template services.Template `dependency:"TemplateService"`
		FindAll  db.FindAll        `dependency:"db.query.article.FindAll"`
	}
	view *template.Template
}

// NewListCtrl create instance of a list articles controller
func NewListCtrl(dp dependency.Provider) (*ListCtrl, error) {
	var err error
	ctrl := &ListCtrl{}
	if err = dp.InjectTo(&ctrl.deps); err != nil {
		return nil, err
	}
	ctrl.view, err = ctrl.deps.Template.View(goathtml.DefaultLayout, "articles/list", nil)
	if err != nil {
		return nil, err
	}
	return ctrl, nil
}

// Get is handler to serve template where one can add new article
func (c *ListCtrl) Get(requestScope app.Scope) {
	var (
		tx          db.TX
		rows        db.Rows
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
	if tx, err = requestDeps.RequestDB.TX(); err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
	if rows, err = c.deps.FindAll(tx); err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
	articleChan := entitychan.NewChanCorverter(requestScope, rows, models.ArticleFactory)
	articleChan.Go()
	if err = c.view.Execute(requestDeps.Response, articleChan.Chan); err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
}
