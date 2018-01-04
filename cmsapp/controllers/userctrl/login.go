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

// Login is register controller
type Login struct {
	deps struct {
		Template services.Template `dependency:"TemplateService"`
		Logger   services.Logger   `dependency:"LoggerService"`
	}
	view *template.Template
}

// NewLogin create instance of a register form controller
func NewLogin(dp dependency.Provider) (*Login, error) {
	var err error
	ctrl := &Login{}
	if err = dp.InjectTo(&ctrl.deps); err != nil {
		return nil, err
	}
	ctrl.view, err = ctrl.deps.Template.View(goathtml.DefaultLayout, "users/login", nil)
	if err != nil {
		return nil, err
	}
	return ctrl, nil
}

func (c *Login) Get(requestScope app.Scope) {
	var requestDeps struct {
		RequestError requestdep.Error     `request:"ErrorService"`
		Responser    requestdep.Responser `request:"ResponserService"`
	}
	if err := requestScope.InjectTo(&requestDeps); err != nil {
		fmt.Println(err)
		return
	}
	if err := requestDeps.Responser.Execute(c.view, map[string]interface{}{
		"Error": false,
	}); err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
}

func (c *Login) Post(requestScope app.Scope) {
	var requestDeps struct {
		Request      *http.Request        `request:"Request"`
		Responser    requestdep.Responser `request:"ResponserService"`
		RequestAuth  requestdep.Auth      `request:"AuthService"`
		RequestError requestdep.Error     `request:"ErrorService"`
		Username     string               `form:"Username"`
		Password     string               `form:"Password"`
	}
	if err := requestScope.InjectTo(&requestDeps); err != nil {
		fmt.Println(err)
		return
	}
	_, err := requestDeps.RequestAuth.Login(requestDeps.Username, requestDeps.Password)
	c.deps.Logger.DevLog("Login.Post controller error: %v", err)
	if err == nil {
		requestDeps.Responser.Redirect("/")
		return
	}
	if err := requestDeps.Responser.Execute(c.view, map[string]interface{}{
		"Error": true,
	}); err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
}
