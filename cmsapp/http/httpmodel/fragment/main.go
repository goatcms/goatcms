package fragment

import (
	"net/http"

	entities "github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/http/httphelpers"
)

// NewForm create new entity instance with data from HTTP request
func NewForm(scope app.Scope, fields []string) (entity *entities.Fragment, err error) {
	var deps struct {
		Req *http.Request `request:"Request"`
	}
	if scope.InjectTo(&deps); err != nil {
		return nil, err
	}
	entity = &entities.Fragment{}
	if err = DecodeFields("", entity, deps.Req, fields); err != nil {
		return nil, err
	}
	return entity, nil
}

// DecodeFields decode HTTP request and set result do entity
func DecodeFields(prefix string, e *entities.Fragment, req *http.Request, fields []string) (err error) {
	if err = httphelpers.ParseForm(req); err != nil {
		return err
	}
	for _, fieldName := range fields {
		switch fieldName {
		case "Name":
			plainName, ok := req.Form[prefix+"Name"]
			if !ok {
				continue
			}
			e.Name = &plainName[0]
		case "Lang":
			plainLang, ok := req.Form[prefix+"Lang"]
			if !ok {
				continue
			}
			e.Lang = &plainLang[0]
		case "Content":
			plainContent, ok := req.Form[prefix+"Content"]
			if !ok {
				continue
			}
			e.Content = &plainContent[0]
		}
	}
	return nil
}
