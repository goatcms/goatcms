package userctrl

import (
	"fmt"
	"html/template"

	"github.com/goatcms/goatcms/cmsapp/forms"
	httpsignup "github.com/goatcms/goatcms/cmsapp/http/httpform/signup"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/goathtml"
	"github.com/goatcms/goatcore/messages"
	"github.com/goatcms/goatcore/messages/msgcollection"
)

// Signup is signup controller
type Signup struct {
	deps struct {
		Template services.Template     `dependency:"TemplateService"`
		Logger   services.Logger       `dependency:"LoggerService"`
		Action   services.SignupAction `dependency:"SignupAction"`
	}
	view *template.Template
}

// NewSignup create instance of a signup form controller
func NewSignup(dp dependency.Provider) (*Signup, error) {
	var err error
	ctrl := &Signup{}
	if err = dp.InjectTo(&ctrl.deps); err != nil {
		return nil, err
	}
	ctrl.view, err = ctrl.deps.Template.View(goathtml.DefaultLayout, "custom/users/signup", nil)
	if err != nil {
		return nil, err
	}
	return ctrl, nil
}

// Get is handler to serve template where one can add new article
func (c *Signup) Get(scope app.Scope) {
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
		"Form":  map[string]interface{}{},
	}); err != nil {
		c.deps.Logger.ErrorLog("%v", err)
		deps.RequestError.Error(312, err)
		return
	}
}

func (c *Signup) Post(scope app.Scope) {
	var (
		err  error
		msgs messages.MessageMap
		form *forms.Signup
		deps struct {
			RequestError requestdep.Error     `request:"ErrorService"`
			Responser    requestdep.Responser `request:"ResponserService"`
		}
	)
	if err = scope.InjectTo(&deps); err != nil {
		fmt.Println(err)
		return
	}
	if form, err = httpsignup.NewForm(scope, forms.SignupAllFields); err != nil {
		c.deps.Logger.ErrorLog("%v", err)
		deps.RequestError.Error(312, err)
		return
	}
	if msgs, err = c.deps.Action.Signup(form, scope); err != nil {
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
