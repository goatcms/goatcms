package userctrl

import (
	"fmt"
	"net/http"

	"github.com/goatcms/goat-core/app"
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/cmsapp/services"
)

// UserLoginController is register controller
type UserLoginController struct {
	deps struct {
		Template services.Template `dependency:"TemplateService"`
	}
}

// NewUserLoginController create instance of a register form controller
func NewUserLoginController(dp dependency.Provider) (*UserLoginController, error) {
	ctrl := &UserLoginController{}
	if err := dp.InjectTo(&ctrl.deps); err != nil {
		return nil, err
	}
	return ctrl, nil
}

func (c *UserLoginController) Get(requestScope app.Scope) {
	var requestDeps struct {
		RequestError services.RequestError `request:"RequestErrorService"`
		Response     http.ResponseWriter   `request:"Response"`
	}
	if err := requestScope.InjectTo(&requestDeps); err != nil {
		fmt.Println(err)
		return
	}
	if err := c.deps.Template.ExecuteTemplate(requestDeps.Response, "users/login", nil); err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
}

func (c *UserLoginController) Post(requestScope app.Scope) {
	var requestDeps struct {
		Request      *http.Request         `request:"Request"`
		Response     http.ResponseWriter   `request:"Response"`
		Template     services.Template     `dependency:"TemplateService"`
		RequestAuth  services.RequestAuth  `request:"RequestAuthService"`
		RequestError services.RequestError `request:"RequestErrorService"`
		Username     string                `form:"Username"`
		Password     string                `form:"Password"`
	}
	if err := requestScope.InjectTo(&requestDeps); err != nil {
		fmt.Println(err)
		return
	}
	_, err := requestDeps.RequestAuth.Login(requestDeps.Username, requestDeps.Password)
	if err == nil {
		http.Redirect(requestDeps.Response, requestDeps.Request, "/", http.StatusSeeOther)
		return
	}
	if err := requestDeps.Template.ExecuteTemplate(requestDeps.Response, "users/login", map[string]interface{}{
		"Errors": []string{"Username or password incorrect"},
	}); err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
}
