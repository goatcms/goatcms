package ruserctrl

import (
	"net/http"

	"github.com/goatcms/goatcms/cmsapp/cmserror"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

// Signout is signout controller
type Signout struct {
	deps struct{}
}

// NewSignout create new signout controller instance
func NewSignout(dp dependency.Provider) (*Signout, error) {
	ctrl := &Signout{}
	if err := dp.InjectTo(&ctrl.deps); err != nil {
		return nil, err
	}
	return ctrl, nil
}

// DO is signout endpoint for POST and GET queries
func (c *Signout) DO(requestScope app.Scope) (err error) {
	var deps struct {
		Logger         services.Logger           `dependency:"LoggerService"`
		Responser      requestdep.Responser      `request:"ResponserService"`
		RequestError   requestdep.Error          `request:"ErrorService"`
		RequestAuth    requestdep.Auth           `request:"AuthService"`
		RequestSession requestdep.SessionManager `request:"SessionService"`
	}
	if err = requestScope.InjectTo(&deps); err != nil {
		return cmserror.NewJSONError(err, http.StatusBadRequest, "{\"status\":\"StatusInternalServerError\"}")
	}
	deps.Logger.TestLog("ruserctrl.Signout: Session destroyed")
	if err = deps.RequestSession.DestroySession(); err != nil {
		return cmserror.NewJSONError(err, http.StatusBadRequest, "{\"status\":\"StatusInternalServerError\"}")
	}
	deps.Responser.JSON(http.StatusCreated, "{\"status\":\"success\"}")
	return nil
}
