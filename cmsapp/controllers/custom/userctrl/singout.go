package userctrl

import (
	"net/http"

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

func NewSignout(dp dependency.Provider) (*Signout, error) {
	ctrl := &Signout{}
	if err := dp.InjectTo(&ctrl.deps); err != nil {
		return nil, err
	}
	return ctrl, nil
}

func (c *Signout) Do(scope app.Scope) {
	var (
		err  error
		deps struct {
			Responser      requestdep.Responser      `request:"ResponserService"`
			RequestError   requestdep.Error          `request:"ErrorService"`
			RequestAuth    requestdep.Auth           `request:"AuthService"`
			RequestSession requestdep.SessionManager `request:"SessionService"`
		}
	)
	if err = scope.InjectTo(&deps); err != nil {
		c.deps.Logger.ErrorLog("%v", err)
		deps.RequestError.Error(http.StatusInternalServerError, err)
		return
	}
	if err = deps.RequestSession.DestroySession(); err != nil {
		c.deps.Logger.ErrorLog("%v", err)
		deps.RequestError.Error(http.StatusInternalServerError, err)
		return
	}
	if err = scope.Trigger(app.CommitEvent, nil); err != nil {
		c.deps.Logger.ErrorLog("%v", err)
		deps.RequestError.Error(http.StatusInternalServerError, err)
		return
	}
	deps.Responser.Redirect("/")
}
