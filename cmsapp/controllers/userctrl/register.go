package userctrl

import (
	"fmt"
	"html/template"

	"github.com/goatcms/goatcms/cmsapp/forms/user/registerform"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/goathtml"
	"github.com/goatcms/goatcore/messages"
	"github.com/goatcms/goatcore/messages/msgcollection"
)

// Register is register controller
type Register struct {
	deps struct {
		Template services.Template           `dependency:"TemplateService"`
		Logger   services.Logger             `dependency:"LoggerService"`
		Action   services.UserRegisterAction `dependency:"UserRegisterAction"`
	}
	view *template.Template
}

// NewRegister create instance of a register form controller
func NewRegister(dp dependency.Provider) (*Register, error) {
	var err error
	ctrl := &Register{}
	if err = dp.InjectTo(&ctrl.deps); err != nil {
		return nil, err
	}
	ctrl.view, err = ctrl.deps.Template.View(goathtml.DefaultLayout, "users/register", nil)
	if err != nil {
		return nil, err
	}
	return ctrl, nil
}

// Get is handler to serve template where one can add new article
func (c *Register) Get(scope app.Scope) {
	var deps struct {
		RequestError requestdep.Error     `request:"ErrorService"`
		Responser    requestdep.Responser `request:"ResponserService"`
	}
	if err := scope.InjectTo(&deps); err != nil {
		c.deps.Logger.ErrorLog("%v", err)
		return
	}
	if err := deps.Responser.Execute(c.view, map[string]interface{}{
		"Valid": msgcollection.NewMessageMap(),
	}); err != nil {
		c.deps.Logger.ErrorLog("%v", err)
		deps.RequestError.Error(312, err)
		return
	}
}

func (c *Register) Post(scope app.Scope) {
	var (
		err  error
		msgs messages.MessageMap
		form *registerform.RegisterForm
		deps struct {
			RequestError requestdep.Error     `request:"ErrorService"`
			Responser    requestdep.Responser `request:"ResponserService"`
		}
	)
	if err = scope.InjectTo(&deps); err != nil {
		fmt.Println(err)
		return
	}
	if form, err = registerform.NewForm(scope); err != nil {
		c.deps.Logger.ErrorLog("%v", err)
		deps.RequestError.Error(312, err)
		return
	}
	if msgs, err = c.deps.Action.Register(form, scope); err != nil {
		c.deps.Logger.ErrorLog("%v", err)
		deps.RequestError.Error(312, err)
		return
	}
	if len(msgs.GetAll()) == 0 {
		deps.Responser.Redirect("/")
		return
	}
	if err := deps.Responser.Execute(c.view, map[string]interface{}{
		"Valid": msgs,
		"Form":  form,
	}); err != nil {
		c.deps.Logger.ErrorLog("%v", err)
		deps.RequestError.Error(312, err)
		return
	}
}
