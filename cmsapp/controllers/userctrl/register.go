package userctrl

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/goathtml"
	"github.com/goatcms/goatcore/messages/msgcollection"
	"github.com/goatcms/goatcms/cmsapp/forms/user/registerform"
	"github.com/goatcms/goatcms/cmsapp/models"
	"github.com/goatcms/goatcms/cmsapp/services"
)

// UserRegisterController is register controller
type UserRegisterController struct {
	deps struct {
		Template      services.Template        `dependency:"TemplateService"`
		RegisterQuery models.UserRegisterQuery `dependency:"db.query.user.RegisterQuery"`
		Crypt         services.Crypt           `dependency:"CryptService"`
	}
	view *template.Template
}

// NewUserRegisterController create instance of a register form controller
func NewUserRegisterController(dp dependency.Provider) (*UserRegisterController, error) {
	var err error
	ctrl := &UserRegisterController{}
	if err = dp.InjectTo(&ctrl.deps); err != nil {
		return nil, err
	}
	ctrl.view, err = ctrl.deps.Template.View(goathtml.DefaultLayout, "users/register", nil)
	if err != nil {
		return nil, err
	}
	return ctrl, nil
}

func (c *UserRegisterController) Get(requestScope app.Scope) {
	var requestDeps struct {
		RequestError services.RequestError `request:"RequestErrorService"`
		Response     http.ResponseWriter   `request:"Response"`
	}
	if err := requestScope.InjectTo(&requestDeps); err != nil {
		fmt.Println(err)
		return
	}
	if err := c.view.Execute(requestDeps.Response, nil); err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
}

func (c *UserRegisterController) Post(requestScope app.Scope) {
	var requestDeps struct {
		Request      *http.Request         `request:"Request"`
		Response     http.ResponseWriter   `request:"Response"`
		DB           services.RequestDB    `request:"RequestDBService"`
		RequestError services.RequestError `request:"RequestErrorService"`
	}
	if err := requestScope.InjectTo(&requestDeps); err != nil {
		fmt.Println(err)
		return
	}
	form, err := registerform.NewForm(requestScope)
	if err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
	validResult := msgcollection.NewMessageMap()
	if err := form.Valid("", validResult); err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
	if len(validResult.GetAll()) == 0 {
		// success
		tx, err := requestDeps.DB.TX()
		if err != nil {
			requestDeps.RequestError.Error(312, err)
			return
		}
		user := models.User(form.User)
		_, err = c.deps.RegisterQuery(tx, &user)
		if err != nil {
			requestDeps.RequestError.Error(312, err)
			return
		}
		http.Redirect(requestDeps.Response, requestDeps.Request, "/", http.StatusSeeOther)
	} else {
		if err := c.view.Execute(requestDeps.Response, map[string]interface{}{
			"formMessages": validResult.GetAll(),
		}); err != nil {
			requestDeps.RequestError.Error(312, err)
			return
		}
	}
}
