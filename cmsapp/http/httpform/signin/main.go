package signin

import (
	"net/http"

	forms "github.com/goatcms/goatcms/cmsapp/forms"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/http/httphelpers"
)

// NewForm create new form instance with data from HTTP request
func NewForm(scope app.Scope, fields []string) (entity *forms.Signin, err error) {
	var deps struct {
		Req *http.Request `request:"Request"`
	}
	if scope.InjectTo(&deps); err != nil {
		return nil, err
	}
	entity = &forms.Signin{}
	if err = DecodeFields("", entity, deps.Req, fields); err != nil {
		return nil, err
	}
	return entity, nil
}

// DecodeFields decode HTTP request and set result do form
func DecodeFields(prefix string, form *forms.Signin, req *http.Request, fields []string) (err error) {
	if err = httphelpers.ParseForm(req); err != nil {
		return err
	}
	for _, fieldName := range fields {
		switch fieldName {
		case "Username":
			plainUsername, ok := req.Form[prefix+"Username"]
			if !ok {
				continue
			}
			form.Username = &plainUsername[0]
		case "Password":
			plainPassword, ok := req.Form[prefix+"Password"]
			if !ok {
				continue
			}
			form.Password = &plainPassword[0]
		}
	}
	return nil
}
