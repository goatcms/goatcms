package services

import (
	"html/template"
	"net/http"

	"github.com/goatcms/goat-core/app"
	"github.com/goatcms/goat-core/db"
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/cmsapp/models"
)

const (
	CreateRequestScope = 1000

	// RequestTagName is a tag name used for a request injection
	RequestTagName = "request"
	// RequestTagName is a tag name used for a request injection
	FormTagName = "form"

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
	// RequestAuthService provide authentication for request
	RequestAuthService = "RequestAuthService"
	// RequestErrorService provide error system
	RequestErrorService = "RequestErrorService"
	// RequestErrorService provide error system
	RequestDBService = "RequestDBService"

	// SessionCookieID is default name of session cookie
	SessionCookieID = "session"
	// SessionCookieLength is default length of session id (storaged by cookie)
	SessionIDLength = 128
	// SessionCookieLifetime is a lifetime of cookie
	SessionLifetime = 365 * 24
	// SessionExpire is key to read expire time from session
	SessionExpireKey = "session.expire"

	// DefaultTemplatePath is a default path for temapates
	DefaultTemplatePath = "./cmsapp/templates"

	// DefaultDatabaseEngine is default engine for database
	DefaultDatabaseEngine = "sqlite3"
	// DefaultDatabaseUrl is default url/path for database
	DefaultDatabaseUrl = "./database/sqlite3_database.db"
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

type SessionStorage interface {
	Get(id string) (app.DataScope, error)
	Create() (string, app.DataScope, error)
	SessionLifetime() (int64, error)
}

type SessionManager interface {
	Init() error
	Scope() (app.DataScope, error)
	app.DataScope
}

type Template interface {
	AddFunc(name string, f interface{}) error
	View(layoutName, viewName string, eventScope app.EventScope) (*template.Template, error)
}

type Crypt interface {
	Hash(pass string) (string, error)
	Compare(hashedPass, pass string) (bool, error)
}

type RequestAuth interface {
	UserID() (string, error)
	Login(name, password string) (*models.User, error)
	Clear() error
}

type RequestError interface {
	Errorf(httpCode int, msgKey string, params ...interface{}) error
	Error(httpCode int, err error)
}

type RequestDB interface {
	TX() (db.TX, error)
}
