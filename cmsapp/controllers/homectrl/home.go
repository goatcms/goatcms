package homectrl

import (
	"fmt"
	"html/template"

	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/goathtml"
)

type Home struct {
	deps struct {
		Template services.Template `dependency:"TemplateService"`
	}
	view *template.Template
}

func NewHome(dp dependency.Provider) (*Home, error) {
	var err error
	ctrl := &Home{}
	if err = dp.InjectTo(&ctrl.deps); err != nil {
		return nil, err
	}
	ctrl.view, err = ctrl.deps.Template.View(goathtml.DefaultLayout, "home/index", nil)
	if err != nil {
		return nil, err
	}
	return ctrl, nil
}

func (c *Home) Get(requestScope app.Scope) {
	var requestDeps struct {
		RequestError requestdep.Error     `request:"ErrorService"`
		Responser    requestdep.Responser `request:"ResponserService"`
	}
	if err := requestScope.InjectTo(&requestDeps); err != nil {
		fmt.Println(err)
		return
	}
	if err := requestDeps.Responser.Execute(c.view, nil); err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
}
