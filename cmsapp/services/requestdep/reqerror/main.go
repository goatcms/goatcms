package reqerror

import (
	"fmt"
	"html/template"
	"reflect"

	"github.com/goatcms/goatcms/cmsapp/cmserror"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/goathtml"
)

// RequestError is error manager for request
type RequestError struct {
	deps struct {
		RequestScope app.Scope            `request:"RequestScope"`
		Responser    requestdep.Responser `request:"ResponserService"`
		Logger       services.Logger      `dependency:"LoggerService"`
		Template     services.Template    `dependency:"TemplateService"`
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
	err := fmt.Errorf(msgKey, params...)
	re.Error(httpCode, err)
	return err
}

// Error print request errror
func (re *RequestError) Error(httpCode int, err error) {
	re.deps.Logger.ErrorLog("%v", err)
	re.deps.RequestScope.Trigger(app.ErrorEvent, err)
}

// DO process a error: log the error and send a error response to client
func (re *RequestError) DO(e error) {
	re.deps.Logger.ErrorLog("%v", e)
	re.deps.RequestScope.Trigger(app.ErrorEvent, e)
	switch v := e.(type) {
	case cmserror.JSONError:
		re.doJSON(v)
	case error:
		re.doTemplate(v)
	default:
		panic(fmt.Errorf("unknow %v error type: %v", reflect.TypeOf(e), e))
	}
}

func (re *RequestError) doTemplate(e error) {
	var (
		err  error
		view *template.Template
	)
	if view, err = re.deps.Template.View(goathtml.DefaultLayout, "custom/error/main", nil); err != nil {
		re.deps.Logger.ErrorLog("%v", err)
		panic(err)
	}
	if err = re.deps.Responser.Execute(view, map[string]interface{}{
		"Error": e.Error(),
	}); err != nil {
		re.deps.Logger.ErrorLog("%v", err)
		panic(err)
	}
}

func (re *RequestError) doJSON(e cmserror.JSONError) {
	var err error
	if err = re.deps.Responser.JSON(e.HTTPCode(), e.Error()); err != nil {
		re.deps.Logger.ErrorLog("%v", err)
		panic(err)
	}
}
