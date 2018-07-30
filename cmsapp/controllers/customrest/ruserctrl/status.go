package ruserctrl

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/goatcms/goatcms/cmsapp/cmserror"
	"github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

// Status is register controller
type Status struct {
	deps struct {
		Logger services.Logger `dependency:"LoggerService"`
	}
}

// NewStatus create instance of a status controller
func NewStatus(dp dependency.Provider) (*Status, error) {
	var err error
	ctrl := &Status{}
	if err = dp.InjectTo(&ctrl.deps); err != nil {
		return nil, err
	}
	return ctrl, nil
}

// DO is status endpoint for POST and GET queries
func (c *Status) DO(scope app.Scope) (err error) {
	var (
		deps struct {
			Logger         services.Logger           `dependency:"LoggerService"`
			Request        *http.Request             `request:"Request"`
			Responser      requestdep.Responser      `request:"ResponserService"`
			SessionManager requestdep.SessionManager `request:"SessionService"`
			RequestError   requestdep.Error          `request:"ErrorService"`
		}
		session   *entities.Session
		rolesJSON string
	)
	if err = scope.InjectTo(&deps); err != nil {
		return cmserror.NewJSONError(err, http.StatusBadRequest, "{\"status\":\"StatusInternalServerError\"}")
	}
	if session, err = deps.SessionManager.Get(); err != nil {
		deps.Logger.ErrorLog("%v", err)
		deps.Responser.JSON(http.StatusOK, "{\"engine\":\"goatapp\",\"status\":\"unauthorized\", \"roles\":[]}")
		return nil
	}
	if session.User != nil && session.User.Roles != nil {
		arr := strings.Split(*session.User.Roles, " ")
		for i, v := range arr {
			arr[i] = strconv.Quote(v)
		}
		rolesJSON = "[" + strings.Join(arr, ",") + "]"
	} else {
		rolesJSON = "[]"
	}
	deps.Responser.JSON(http.StatusOK, "{\"engine\":\"goatapp\",\"status\":\"loggedin\", \"roles\":"+rolesJSON+"}")
	return nil
}
