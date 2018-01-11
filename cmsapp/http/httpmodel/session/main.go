package session

import (
	"net/http"

	entities "github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/http/httphelpers"
)

// NewForm create new entity instance with data from HTTP request
func NewForm(scope app.Scope, fields []string) (entity *entities.Session, err error) {
	var deps struct {
		Req *http.Request `request:"Request"`
	}
	if scope.InjectTo(&deps); err != nil {
		return nil, err
	}
	entity = &entities.Session{}
	if err = DecodeFields("", entity, deps.Req, fields); err != nil {
		return nil, err
	}
	return entity, nil
}

// DecodeFields decode HTTP request and set result do entity
func DecodeFields(prefix string, e *entities.Session, req *http.Request, fields []string) (err error) {
	if err = httphelpers.ParseForm(req); err != nil {
		return err
	}
	for _, fieldName := range fields {
		switch fieldName {
		case "Secret":
			plainSecret, ok := req.Form[prefix+"Secret"]
			if !ok {
				continue
			}
			e.Secret = &plainSecret[0]
		}
	}
	return nil
}
