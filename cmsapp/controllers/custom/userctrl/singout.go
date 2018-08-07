package userctrl

import (
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

// Signout is signout controller
type Signout struct {
	deps struct {
		Logger services.Logger `dependency:"LoggerService"`
	}
}

// NewSignout create new Signout controller
func NewSignout(dp dependency.Provider) (*Signout, error) {
	ctrl := &Signout{}
	if err := dp.InjectTo(&ctrl.deps); err != nil {
		return nil, err
	}
	return ctrl, nil
}

// Do is a standard endpoint for GET and POST request
func (c *Signout) Do(scope app.Scope) (err error) {
	var (
		deps struct {
			Responser      requestdep.Responser      `request:"ResponserService"`
			RequestAuth    requestdep.Auth           `request:"AuthService"`
			RequestSession requestdep.SessionManager `request:"SessionService"`
		}
	)
	if err = scope.InjectTo(&deps); err != nil {
		return err
	}
	if err = deps.RequestSession.DestroySession(); err != nil {
		return err
	}
	c.deps.Logger.TestLog("userctrl.Signout: Session destroyed")
	if err = scope.Trigger(app.CommitEvent, nil); err != nil {
		return err
	}
	deps.Responser.Redirect("/")
	return nil
}
