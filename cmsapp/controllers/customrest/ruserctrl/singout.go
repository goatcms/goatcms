package ruserctrl

import (
	"net/http"

	"github.com/goatcms/goatcms/cmsapp/services"
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

func (c *Signout) DO(requestScope app.Scope) {
	var deps struct {
		Logger         services.Logger           `dependency:"LoggerService"`
		Responser      requestdep.Responser      `request:"ResponserService"`
		RequestError   requestdep.Error          `request:"ErrorService"`
		RequestAuth    requestdep.Auth           `request:"AuthService"`
		RequestSession requestdep.SessionManager `request:"SessionService"`
	}
	if err := requestScope.InjectTo(&deps); err != nil {
		deps.Logger.ErrorLog("%v", err)
		deps.Responser.JSON(http.StatusBadRequest, "{\"status\":\"StatusInternalServerError\"}")
		return
	}
	if err := deps.RequestSession.DestroySession(); err != nil {
		deps.Logger.ErrorLog("%v", err)
		deps.Responser.JSON(http.StatusBadRequest, "{\"status\":\"StatusInternalServerError\"}")
		return
	}
	deps.Responser.JSON(http.StatusCreated, "{\"status\":\"success\"}")
}
