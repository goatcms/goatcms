package userctrl

import (
	"fmt"

	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

// Signout is signout controller
type Signout struct {
	deps struct{}
}

func NewSignout(dp dependency.Provider) (*Signout, error) {
	ctrl := &Signout{}
	if err := dp.InjectTo(&ctrl.deps); err != nil {
		return nil, err
	}
	return ctrl, nil
}

func (c *Signout) Do(requestScope app.Scope) {
	var requestDeps struct {
		Responser      requestdep.Responser      `request:"ResponserService"`
		RequestError   requestdep.Error          `request:"ErrorService"`
		RequestAuth    requestdep.Auth           `request:"AuthService"`
		RequestSession requestdep.SessionManager `request:"SessionService"`
	}
	if err := requestScope.InjectTo(&requestDeps); err != nil {
		fmt.Println(err)
		return
	}
	if err := requestDeps.RequestSession.DestroySession(); err != nil {
		requestDeps.RequestError.Error(312, err)
	}
	requestDeps.Responser.Redirect("/")
}
