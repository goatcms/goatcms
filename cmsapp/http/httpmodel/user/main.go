package user

import (
	"net/http"

	entities "github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/http/httphelpers"
)

// NewForm create new entity instance with data from HTTP request
func NewForm(scope app.Scope, fields []string) (entity *entities.User, err error) {
	var deps struct {
		Req *http.Request `request:"Request"`
	}
	if scope.InjectTo(&deps); err != nil {
		return nil, err
	}
	entity = &entities.User{}
	if err = DecodeFields("", entity, deps.Req, fields); err != nil {
		return nil, err
	}
	return entity, nil
}

// DecodeFields decode HTTP request and set result do entity
func DecodeFields(prefix string, e *entities.User, req *http.Request, fields []string) (err error) {
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
			e.Username = &plainUsername[0]
		case "Email":
			plainEmail, ok := req.Form[prefix+"Email"]
			if !ok {
				continue
			}
			e.Email = &plainEmail[0]
		case "Lastname":
			plainLastname, ok := req.Form[prefix+"Lastname"]
			if !ok {
				continue
			}
			e.Lastname = &plainLastname[0]
		case "Password":
			plainPassword, ok := req.Form[prefix+"Password"]
			if !ok {
				continue
			}
			e.Password = &plainPassword[0]
		case "Roles":
			plainRoles, ok := req.Form[prefix+"Roles"]
			if !ok {
				continue
			}
			e.Roles = &plainRoles[0]
		case "Firstname":
			plainFirstname, ok := req.Form[prefix+"Firstname"]
			if !ok {
				continue
			}
			e.Firstname = &plainFirstname[0]
		}
	}
	return nil
}
