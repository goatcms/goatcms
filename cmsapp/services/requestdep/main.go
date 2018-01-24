package requestdep

import (
	"html/template"

	"github.com/goatcms/goatcms/cmsapp/entities"
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

type ACL interface {
	HasAnyRole(roles []string) bool
}

type SessionManager interface {
	LoadSession() (err error)
	Get() (session *entities.Session, err error)
	CreateSession(user *entities.User) (session *entities.Session, err error)
	DestroySession() (err error)
}

type Auth interface {
	Signin(name, password string) (*entities.Session, error)
	Signout() error
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
	JSON(code int, json string) (err error)
	Redirect(url string)
	IsSended() bool
}
