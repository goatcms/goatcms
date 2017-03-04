package userctrl

import (
	"fmt"

	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

// Logout is logout controller
type Logout struct {
	deps struct{}
}

// NewUserRegisterController create instance of a register form controller
func NewLogout(dp dependency.Provider) (*Logout, error) {
	ctrl := &Logout{}
	if err := dp.InjectTo(&ctrl.deps); err != nil {
		return nil, err
	}
	return ctrl, nil
}

func (c *Logout) Do(requestScope app.Scope) {
	var requestDeps struct {
		Responser    requestdep.Responser `request:"ResponserService"`
		RequestError requestdep.Error     `request:"ErrorService"`
		RequestAuth  requestdep.Auth      `request:"AuthService"`
	}
	if err := requestScope.InjectTo(&requestDeps); err != nil {
		fmt.Println(err)
		return
	}
	if err := requestDeps.RequestAuth.Clear(); err != nil {
		requestDeps.RequestError.Error(312, err)
	}
	requestDeps.Responser.Redirect("/")
}
