package signup

import (
	"net/http"

	forms "github.com/goatcms/goatcms/cmsapp/forms"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/http/httphelpers"
)

// NewForm create new form instance with data from HTTP request
func NewForm(scope app.Scope, fields []string) (entity *forms.Signup, err error) {
	var deps struct {
		Req *http.Request `request:"Request"`
	}
	if scope.InjectTo(&deps); err != nil {
		return nil, err
	}
	entity = &forms.Signup{}
	if err = DecodeFields("", entity, deps.Req, fields); err != nil {
		return nil, err
	}
	return entity, nil
}

// DecodeFields decode HTTP request and set result do form
func DecodeFields(prefix string, form *forms.Signup, req *http.Request, fields []string) (err error) {
	if err = httphelpers.ParseForm(req); err != nil {
		return err
	}
	for _, fieldName := range fields {
		switch fieldName {
		case "Password":
			plainPasswordFirst, ok := req.Form[prefix+"Password.First"]
			if !ok {
				continue
			}
			plainPasswordSecond, ok := req.Form[prefix+"Password.Second"]
			if !ok {
				continue
			}
			form.Password = &forms.RepeatPassword{
				First:  plainPasswordFirst[0],
				Second: plainPasswordSecond[0],
			}
		case "Firstname":
			plainFirstname, ok := req.Form[prefix+"Firstname"]
			if !ok {
				continue
			}
			form.Firstname = &plainFirstname[0]
		case "Username":
			plainUsername, ok := req.Form[prefix+"Username"]
			if !ok {
				continue
			}
			form.Username = &plainUsername[0]
		case "Lastname":
			plainLastname, ok := req.Form[prefix+"Lastname"]
			if !ok {
				continue
			}
			form.Lastname = &plainLastname[0]
		case "Email":
			plainEmail, ok := req.Form[prefix+"Email"]
			if !ok {
				continue
			}
			form.Email = &plainEmail[0]
		}
	}
	return nil
}
