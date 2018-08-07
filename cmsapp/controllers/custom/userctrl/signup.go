package userctrl

import (
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

// Get is signup endpoint for GET method
func (c *Signup) Get(scope app.Scope) (err error) {
	var deps struct {
		RequestError requestdep.Error     `request:"ErrorService"`
		Responser    requestdep.Responser `request:"ResponserService"`
	}
	if err = scope.InjectTo(&deps); err != nil {
		return err
	}
	return deps.Responser.Execute(c.view, map[string]interface{}{
		"Valid": msgcollection.NewMessageMap(),
		"Form":  map[string]interface{}{},
	})
}

// Post is signup endpoint for POST method
func (c *Signup) Post(scope app.Scope) (err error) {
	var (
		msgs messages.MessageMap
		form *forms.Signup
		deps struct {
			Responser requestdep.Responser `request:"ResponserService"`
		}
	)
	if err = scope.InjectTo(&deps); err != nil {
		return err
	}
	if form, err = httpsignup.NewForm(scope, forms.SignupAllFields); err != nil {
		return err
	}
	c.deps.Logger.DevLog("userctrl.Signup: Process signup form %v", form)
	if msgs, err = c.deps.Action.Signup(form, scope); err != nil {
		return err
	}
	if len(msgs.GetAll()) == 0 {
		c.deps.Logger.TestLog("userctrl.Signup: Process signup form correct")
		deps.Responser.Redirect("/")
		return nil
	}
	c.deps.Logger.TestLog("userctrl.Signup: Process signup form incorrect %v", msgs.GetAll())
	return deps.Responser.Execute(c.view, map[string]interface{}{
		"Valid": msgs,
		"Form":  form,
	})
}
