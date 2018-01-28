package userctrl

import (
	"fmt"
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

func (c *Signin) Get(scope app.Scope) {
	var deps struct {
		RequestError requestdep.Error     `request:"ErrorService"`
		Responser    requestdep.Responser `request:"ResponserService"`
	}
	if err := scope.InjectTo(&deps); err != nil {
		fmt.Println(err)
		return
	}
	if err := deps.Responser.Execute(c.view, map[string]interface{}{
		"Error": false,
	}); err != nil {
		deps.RequestError.Error(312, err)
		return
	}
}

func (c *Signin) Post(scope app.Scope) {
	var (
		err  error
		deps struct {
			Request      *http.Request        `request:"Request"`
			Responser    requestdep.Responser `request:"ResponserService"`
			RequestAuth  requestdep.Auth      `request:"AuthService"`
			RequestError requestdep.Error     `request:"ErrorService"`
			Username     string               `form:"Username"`
			Password     string               `form:"Password"`
		}
	)
	if err = scope.InjectTo(&deps); err != nil {
		c.deps.Logger.ErrorLog("%v", err)
		deps.RequestError.Error(http.StatusInternalServerError, err)
		return
	}
	if _, err = deps.RequestAuth.Signin(deps.Username, deps.Password); err != nil {
		// there can be incorrect password error - only log it
		c.deps.Logger.ErrorLog("%v", err)
		// show login panel again
		if err := deps.Responser.Execute(c.view, map[string]interface{}{
			"Error": true,
		}); err != nil {
			c.deps.Logger.ErrorLog("%v", err)
			deps.RequestError.Error(http.StatusInternalServerError, err)
			return
		}
	}
	if err = scope.Trigger(app.CommitEvent, nil); err != nil {
		c.deps.Logger.ErrorLog("%v", err)
		deps.RequestError.Error(http.StatusInternalServerError, err)
		return
	}
	if err == nil {
		deps.Responser.Redirect("/")
		return
	}
}
