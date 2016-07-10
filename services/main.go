package services

import (
	"database/sql"
	"io"
	"net/http"
)

const (
	// DBID is a name representing default database service
	DBID = "database"
	// MuxID is a name representing default mux service
	MuxID = "mux"
	// TemplateID is a name represents default template service
	TemplateID = "template"
	// CryptID is a name represents default crypt service
	CryptID = "crypt"
)

// Database is global elementary database interface
type Database interface {
	Open() error
	Close() error
	CreateTables() error
	// Deprecated: It shouldn't be use	// what shouldn't?
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
}

// Crypt is global elementary cryptographic interface
type Crypt interface {
	Hash(pass string) (string, error)
}
