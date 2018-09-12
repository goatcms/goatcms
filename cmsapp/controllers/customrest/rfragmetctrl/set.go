package ruserctrl

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/goatcms/goatcms/cmsapp/cmserror"
	"github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcms/cmsapp/forms"
	"github.com/goatcms/goatcms/cmsapp/http/httpform/signin"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/goathtml"
)

// SetCtrl is set value controller
type SetCtrl struct {
	deps struct {
		Template services.Template `dependency:"TemplateService"`
		Logger   services.Logger   `dependency:"LoggerService"`
	}
}

// NewSignin create instance of a register form controller
func NewSignin(dp dependency.Provider) (*SetCtrl, error) {
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

// DO is signin endpoint for POST and GET queries
func (c *Signin) DO(scope app.Scope) (err error) {
	var (
		deps struct {
			Logger       services.Logger      `dependency:"LoggerService"`
			Request      *http.Request        `request:"Request"`
			Responser    requestdep.Responser `request:"ResponserService"`
			RequestAuth  requestdep.Auth      `request:"AuthService"`
			RequestError requestdep.Error     `request:"ErrorService"`
		}
		form      *forms.Signin
		session   *entities.Session
		rolesJSON string
	)
	if err = scope.InjectTo(&deps); err != nil {
		return cmserror.NewJSONError(err, http.StatusBadRequest, "{\"status\":\"StatusInternalServerError\"}")
	}
	if form, err = signin.NewForm(scope, forms.SigninAllFields); err != nil {
		return cmserror.NewJSONError(err, http.StatusBadRequest, "{\"status\":\"StatusInternalServerError\"}")
	}
	c.deps.Logger.DevLog("restctrl.Signin: Created user session %v", session)
	if session, err = deps.RequestAuth.Signin(*form.Username, *form.Password); err != nil {
		return cmserror.NewJSONError(err, http.StatusBadRequest, "{\"status\":\"StatusBadRequest\"}")
	}
	if err = scope.Trigger(app.CommitEvent, nil); err != nil {
		return cmserror.NewJSONError(err, http.StatusBadRequest, "{\"status\":\"StatusInternalServerError\"}")
	}
	if session.User != nil && session.User.Roles != nil {
		arr := strings.Split(*session.User.Roles, " ")
		for i, v := range arr {
			arr[i] = strconv.Quote(v)
		}
		rolesJSON = "[" + strings.Join(arr, ",") + "]"
	} else {
		rolesJSON = "[]"
	}
	deps.Responser.JSON(http.StatusOK, "{\"status\":\"success\", \"secret\":"+strconv.Quote(*session.Secret)+", \"roles\":"+rolesJSON+"}")
	return nil
}
