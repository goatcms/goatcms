package homectrl

import (
	"html/template"

	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/goathtml"
)

// Home is home page controller
type Home struct {
	deps struct {
		Template services.Template `dependency:"TemplateService"`
	}
	view *template.Template
}

// NewHome create new Home controller instance
func NewHome(dp dependency.Provider) (*Home, error) {
	var err error
	ctrl := &Home{}
	if err = dp.InjectTo(&ctrl.deps); err != nil {
		return nil, err
	}
	ctrl.view, err = ctrl.deps.Template.View(goathtml.DefaultLayout, "custom/home/main", nil)
	if err != nil {
		return nil, err
	}
	return ctrl, nil
}

// Get render Home controller
func (c *Home) Get(requestScope app.Scope) (err error) {
	var deps struct {
		RequestError requestdep.Error     `request:"ErrorService"`
		Responser    requestdep.Responser `request:"ResponserService"`
		ACL          requestdep.ACL       `request:"ACLService"`
	}
	if err = requestScope.InjectTo(&deps); err != nil {
		return err
	}
	if err = deps.Responser.Execute(c.view, map[string]interface{}{
		"Editable": deps.ACL.HasAnyRole([]string{"admin"}),
	}); err != nil {
		return err
	}
	return err
}
