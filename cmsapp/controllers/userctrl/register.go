package userctrl

import (
	"fmt"
	"html/template"

	"github.com/goatcms/goatcms/cmsapp/forms/user/registerform"
	"github.com/goatcms/goatcms/cmsapp/models"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/db"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/goathtml"
	"github.com/goatcms/goatcore/goatmail"
	"github.com/goatcms/goatcore/messages/msgcollection"
)

// Register is register controller
type Register struct {
	deps struct {
		Template services.Template   `dependency:"TemplateService"`
		Register models.UserRegister `dependency:"UserRegister"`
		Crypt    services.Crypt      `dependency:"CryptService"`
		Mailer   services.Mailer     `dependency:"MailerService"`
		Logger   services.Logger     `dependency:"LoggerService"`
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
func (c *Register) Get(requestScope app.Scope) {
	var requestDeps struct {
		RequestError requestdep.Error     `request:"ErrorService"`
		Responser    requestdep.Responser `request:"ResponserService"`
	}
	if err := requestScope.InjectTo(&requestDeps); err != nil {
		fmt.Println(err)
		return
	}
	if err := requestDeps.Responser.Execute(c.view, map[string]interface{}{
		"Valid": msgcollection.NewMessageMap(),
	}); err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
}

func (c *Register) Post(requestScope app.Scope) {
	var (
		tx          db.TX
		err         error
		requestDeps struct {
			RequestDB    requestdep.DB        `request:"DBService"`
			RequestError requestdep.Error     `request:"ErrorService"`
			Responser    requestdep.Responser `request:"ResponserService"`
		}
	)
	if err = requestScope.InjectTo(&requestDeps); err != nil {
		fmt.Println(err)
		return
	}
	form, err := registerform.NewForm(requestScope)
	if err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
	validResult := msgcollection.NewMessageMap()
	if err = form.Valid("", validResult); err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
	if len(validResult.GetAll()) == 0 {
		// success
		c.deps.Logger.DevLog("Register new user %v", form.User)
		tx, err = requestDeps.RequestDB.TX()
		if err != nil {
			requestDeps.RequestError.Error(312, err)
			return
		}
		user := models.User(form.User)
		_, err = c.deps.Register(tx, &user, form.Password)
		if err != nil {
			requestDeps.RequestError.Error(312, err)
			return
		}
		if err = tx.Commit(); err != nil {
			requestDeps.RequestError.Error(312, err)
			return
		}
		c.deps.Mailer.Send(form.User.Email, "register", form, []goatmail.Attachment{}, requestScope)
		requestDeps.Responser.Redirect("/")
	} else {
		// fail
		c.deps.Logger.DevLog("Register user valid error: %v %v %v", form, form.User, validResult.GetAll())
		if err := requestDeps.Responser.Execute(c.view, map[string]interface{}{
			"Valid": validResult,
			"Form":  form,
		}); err != nil {
			requestDeps.RequestError.Error(312, err)
			return
		}
	}
}
