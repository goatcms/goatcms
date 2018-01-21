package router

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gorilla/mux"
)

// Parameters storage and provide url parameters
type Parameters struct {
	tagname string
	data    map[string]string
}

// NewParametersFromRequest create new Parameters instance from request URL data
func NewParametersFromRequest(req *http.Request) *Parameters {
	return NewParameters(mux.Vars(req))
}

// NewParameters create new Parameters instance
func NewParameters(data map[string]string) *Parameters {
	return &Parameters{
		tagname: "parameter",
		data:    data,
	}
}

// InjectTo inject URL parameters to object
func (params Parameters) InjectTo(obj interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	for i := 0; i < structValue.NumField(); i++ {
		var isRequired = true
		valueField := structValue.Field(i)
		structField := structValue.Type().Field(i)
		key := structField.Tag.Get(params.tagname)
		if key == "" {
			continue
		}
		if strings.HasPrefix(key, "?") {
			isRequired = false
			key = key[1:]
		}
		if !valueField.IsValid() {
			return fmt.Errorf("MapInjector.InjectTo: %s is not valid", structField.Name)
		}
		if !valueField.CanSet() {
			return fmt.Errorf("MapInjector.InjectTo: Cannot set %s field value", structField.Name)
		}
		newValue, ok := params.data[key]
		if !ok {
			if !isRequired {
				continue
			}
			return fmt.Errorf("parameter for %s is unknown", key)
		}
		refValue := reflect.ValueOf(newValue)
		valueField.Set(refValue)
	}
	return nil
}
