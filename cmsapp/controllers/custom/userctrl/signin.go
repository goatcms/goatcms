package userctrl

import (
	"html/template"
	"net/http"

	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/goathtml"
)

// Signin is register controller
type Signin struct {
	deps struct {
		Template services.Template `dependency:"TemplateService"`
		Logger   services.Logger   `dependency:"LoggerService"`
	}
	view *template.Template
}

// NewSignin create instance of a register form controller
func NewSignin(dp dependency.Provider) (*Signin, error) {
	var err error
	ctrl := &Signin{}
	if err = dp.InjectTo(&ctrl.deps); err != nil {
		return nil, err
	}
	ctrl.view, err = ctrl.deps.Template.View(goathtml.DefaultLayout, "custom/users/signin", nil)
	if err != nil {
		return nil, err
	}
	return ctrl, nil
}

// Get is signin endpoint for GET method
func (c *Signin) Get(scope app.Scope) (err error) {
	var (
		deps struct {
			Responser requestdep.Responser `request:"ResponserService"`
		}
	)
	if err = scope.InjectTo(&deps); err != nil {
		return err
	}
	return deps.Responser.Execute(c.view, map[string]interface{}{
		"Error": false,
	})
}

// Post is signin endpoint for POST method
func (c *Signin) Post(scope app.Scope) (err error) {
	var (
		deps struct {
			Request     *http.Request        `request:"Request"`
			Responser   requestdep.Responser `request:"ResponserService"`
			RequestAuth requestdep.Auth      `request:"AuthService"`
			Username    string               `form:"Username"`
			Password    string               `form:"Password"`
		}
	)
	if err = scope.InjectTo(&deps); err != nil {
		return err
	}
	if _, err = deps.RequestAuth.Signin(deps.Username, deps.Password); err != nil {
		c.deps.Logger.ErrorLog("%v", err)
		return deps.Responser.Execute(c.view, map[string]interface{}{
			"Error": true,
		})
	}
	if err = scope.Trigger(app.CommitEvent, nil); err != nil {
		return err
	}
	deps.Responser.Redirect("/")
	return nil
}
