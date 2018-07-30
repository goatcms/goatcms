package ruserctrl

import (
	"net/http"

	"github.com/goatcms/goatcms/cmsapp/cmserror"
	"github.com/goatcms/goatcms/cmsapp/forms"
	httpsignup "github.com/goatcms/goatcms/cmsapp/http/httpform/signup"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/messages"
)

// Signup is signup controller
type Signup struct {
	deps struct {
		Template services.Template     `dependency:"TemplateService"`
		Logger   services.Logger       `dependency:"LoggerService"`
		Action   services.SignupAction `dependency:"SignupAction"`
	}
}

// NewSignup create instance of a signup form controller
func NewSignup(dp dependency.Provider) (*Signup, error) {
	var err error
	ctrl := &Signup{}
	if err = dp.InjectTo(&ctrl.deps); err != nil {
		return nil, err
	}
	return ctrl, nil
}

// DO is signup endpoint for POST and GET queries
func (c *Signup) DO(scope app.Scope) (err error) {
	var (
		msgs messages.MessageMap
		form *forms.Signup
		deps struct {
			Logger       services.Logger      `dependency:"LoggerService"`
			RequestError requestdep.Error     `request:"ErrorService"`
			Responser    requestdep.Responser `request:"ResponserService"`
		}
	)
	if err = scope.InjectTo(&deps); err != nil {
		return cmserror.NewJSONError(err, http.StatusBadRequest, "{\"status\":\"StatusInternalServerError\"}")
	}
	if form, err = httpsignup.NewForm(scope, forms.SignupAllFields); err != nil {
		return cmserror.NewJSONError(err, http.StatusBadRequest, "{\"status\":\"StatusInternalServerError\"}")
	}
	if msgs, err = c.deps.Action.Signup(form, scope); err != nil {
		return cmserror.NewJSONError(err, http.StatusBadRequest, "{\"status\":\"StatusBadRequest\"}")
	}
	if len(msgs.GetAll()) != 0 {
		deps.Logger.ErrorLog("%v", err)
		deps.Responser.JSON(http.StatusBadRequest, msgs.ToJSON())
		return
	}
	deps.Responser.JSON(http.StatusCreated, "{\"status\":\"success\"}")
	return nil
}
