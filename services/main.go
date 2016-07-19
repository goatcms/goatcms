package services

import (
	"database/sql"
	"html/template"
	"io"
	"net/http"

	"github.com/goatcms/goat-core/dependency"
)

const (
	// DBID is a name representing default database service
	DBID = "database"
	// MuxID is a name representing default mux service
	MuxID = "mux"
	// TemplateID is a name representing default template service
	TemplateID = "template"
	// CryptID is a name representing default crypt service
	CryptID = "crypt"
	// AuthID is a name representing default authentification service
	AuthID = "auth"
	// SessionManagerID is a name representing default session manager service
	SessionManagerID = "session"
)

// Database is global elementary database interface
type Database interface {
	Open() error
	Close() error
	CreateTables() error
	// Deprecated: It shouldn't be use
	Adapter() *sql.DB
}

// MuxHandler function for routing dispatcher
type MuxHandler func(http.ResponseWriter, *http.Request)

// Mux is global elementary routing interface
type Mux interface {
	Get(string, MuxHandler)
	Post(string, MuxHandler)
	Put(string, MuxHandler)
	Delete(string, MuxHandler)
	Start() error
}

// Template is global elementary routing interface
type Template interface {
	ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	Funcs(funcMap template.FuncMap) error
}

// Crypt is global elementary cryptographic interface
type Crypt interface {
	Hash(pass string) (string, error)
	Compare(hashedPass, pass string) (bool, error)
}

// Auth is global elementary authentification interface
type Auth interface {
	GetUserID(sessid string) (string, error)
	Auth(sessid string, userid string) error
	Clear(sessid string) error
}

// SessionManager is global elementary session interface
type SessionManager interface {
	Init(w http.ResponseWriter, r *http.Request) (string, error)
	Get(string, string) (string, error)
	Set(string, string, string) error
}

// Provider is service dependency provider extension
type Provider interface {
	dependency.Provider

	Database() (Database, error)
	Mux() (Mux, error)
	Template() (Template, error)
	Crypt() (Crypt, error)
	Auth() (Auth, error)
	SessionManager() (SessionManager, error)
}
