package requestdep

import (
	"html/template"

	"github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/db"
)

const (
	// AuthService provide authentication for request
	AuthService = "AuthService"
	// ErrorService provide error system
	ErrorService = "ErrorService"
	// ErrorService provide error system
	DBService = "DBService"
	// SessionService provide sessions accessor
	SessionService = "SessionService"
	// ResponseService provide http response system
	ResponserService = "ResponserService"
	// TranslateService provide trasnlate system (for current user language)
	TranslateService = "TranslateService"
)

type Session interface {
	Init() error
	Scope() (app.DataScope, error)
	app.DataScope
}

type Auth interface {
	UserID() (string, error)
	Login(name, password string) (*entities.User, error)
	Clear() error
}

type Error interface {
	Errorf(httpCode int, msgKey string, params ...interface{}) error
	Error(httpCode int, err error)
}

type DB interface {
	TX() (db.TX, error)
}

type Translate interface {
	Translate(key string, values ...interface{}) (string, error)
	Lang() string
}

type Responser interface {
	Execute(template *template.Template, data interface{}) error
	Redirect(url string)
	IsSended() bool
}
