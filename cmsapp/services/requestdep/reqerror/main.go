package reqerror

import (
	"fmt"

	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

// RequestError is error manager for request
type RequestError struct {
	deps struct {
		RequestScope app.Scope `request:"RequestScope"`
	}
}

// ErrorFactory create an RequestError service instance
func ErrorFactory(dp dependency.Provider) (interface{}, error) {
	instance := &RequestError{}
	if err := dp.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return requestdep.Error(instance), nil
}

// Errorf print request errror
func (re *RequestError) Errorf(httpCode int, msgKey string, params ...interface{}) error {
	err := fmt.Errorf(msgKey, params)
	re.Error(httpCode, err)
	return err
}

// Error print request errror
func (re *RequestError) Error(httpCode int, err error) {
	fmt.Printf("\n error: %v\n", err)
	re.deps.RequestScope.Trigger(app.ErrorEvent, err)
}
