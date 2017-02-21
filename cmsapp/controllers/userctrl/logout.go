package userctrl

import (
	"fmt"
	"net/http"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcms/cmsapp/services"
)

// UserLogoutController is logout controller
type UserLogoutController struct {
	deps struct{}
}

// NewUserRegisterController create instance of a register form controller
func NewUserLogoutController(dp dependency.Provider) (*UserLogoutController, error) {
	ctrl := &UserLogoutController{}
	if err := dp.InjectTo(&ctrl.deps); err != nil {
		return nil, err
	}
	return ctrl, nil
}

func (c *UserLogoutController) All(requestScope app.Scope) {
	var requestDeps struct {
		Request      *http.Request         `request:"Request"`
		Response     http.ResponseWriter   `request:"Response"`
		RequestError services.RequestError `request:"RequestErrorService"`
		RequestAuth  services.RequestAuth  `request:"RequestAuthService"`
	}
	if err := requestScope.InjectTo(&requestDeps); err != nil {
		fmt.Println(err)
		return
	}
	if err := requestDeps.RequestAuth.Clear(); err != nil {
		requestDeps.RequestError.Error(312, err)
	}
	http.Redirect(requestDeps.Response, requestDeps.Request, "/", http.StatusSeeOther)
}
