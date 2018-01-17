package services

import (
	"html/template"
	"net/http"

	"github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcms/cmsapp/forms"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/db"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/goatmail"
	"github.com/goatcms/goatcore/messages"
)

const (
	// DatabaseService is a key to access database storage service
	DatabaseService = "DatabaseService"
	// RouterKey is a key to access router service
	RouterService = "RouterService"
	// SessionStorageService is a key to access session storage service
	SessionStorageService = "SessionStorageService"
	// SessionManagerService is a key to access session manager service
	SessionManagerService = "SessionManagerService"
	// TemplateProviderService is a key to access template manager service
	TemplateService = "TemplateService"
	// CryptProviderService is a key to access crypting/encrypting manager service
	CryptService = "CryptService"
	// MailerService provide mail system
	MailerService = "MailerService"
	// LoggerService provide logger
	LoggerService = "LoggerService"
	// LoggerService provide logger
	TranslateService = "TranslateService"
	// SignupActionService provide user register service
	SignupActionService = "SignupAction"
)

// MuxHandler is function for standard mux input
type MuxHandler func(http.ResponseWriter, *http.Request)

// ScopeHandler is a router service handler
type ScopeHandler func(app.Scope)

// Database is global elementary database interface
type Database interface {
	Open() error
	Close() error
	TX() (db.TX, error)
	FlushTX() (db.TX, error)
}

type Router interface {
	OnGet(path string, handler ScopeHandler)
	OnPost(path string, handler ScopeHandler)
	OnPut(path string, handler ScopeHandler)
	OnDelete(path string, handler ScopeHandler)
	On(methods []string, path string, handler ScopeHandler)
	Host() string
	Start() error
	AddFactory(name string, factory dependency.Factory) error
}

/*type SessionStorage interface {
	Get(id string) (app.DataScope, error)
	Create() (string, app.DataScope, error)
	SessionLifetime() (int64, error)
}*/

type SessionManager interface {
	Get(scope app.Scope, secret string) (session *entities.Session, err error)
	Create(scope app.Scope, user *entities.User) (session *entities.Session, err error)
	Delete(scope app.Scope, secret string) (err error)
}

type Template interface {
	AddFunc(name string, f interface{}) error
	View(layoutName, viewName string, eventScope app.EventScope) (*template.Template, error)
}

type Crypt interface {
	Hash(pass string) (string, error)
	Compare(hashedPass, pass string) (bool, error)
}

type Mailer interface {
	Send(to, name string, data interface{}, attachments []goatmail.Attachment, scope app.Scope)
}

type Logger interface {
	DevLog(format string, data ...interface{})
	TestLog(format string, data ...interface{})
	ProdLog(format string, data ...interface{})
	ErrorLog(format string, data ...interface{})
}

type Translate interface {
	Translate(key string, values ...interface{}) (string, error)
	TranslateFor(key, prefix string, values ...interface{}) (string, error)
	Langs() []string
	Default() string
}

type SignupAction interface {
	Signup(form *forms.Signup, scope app.Scope) (msgs messages.MessageMap, err error)
}
