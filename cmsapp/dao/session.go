package dao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app"
)

// SessionRows is the result of a query. Its cursor starts before the first row of the result set. Use Next to advance through the rows
type SessionRows interface {
	Rows
	InjectTo(*entities.Session) error
	Get() (*entities.Session, error)
}

// SessionRow is the result of calling QueryRow to select a single row.
type SessionRow interface {
	Row
	InjectTo(*entities.Session) error
	Get() (*entities.Session, error)
}

// SessionFindAll is the DAO find all provider interface
type SessionFindAll interface {
	Find(scope app.Scope, fields []string) (SessionRows, error)
	SQL(fields []string) (string, error)
}

// SessionFindByID is the DAO find by id provider interface
type SessionFindByID interface {
	Find(scope app.Scope, fields []string, id int64) (user *entities.Session, err error)
	SQL(fields []string, id int64) (string, error)
}

// SessionInsert is the DAO insert provider interface
type SessionInsert interface {
	Insert(scope app.Scope, entity *entities.Session) (id int64, err error)
	SQL(entity *entities.Session) (string, error)
}

// SessionUpdate is the DAO update provider interface
type SessionUpdate interface {
	Update(scope app.Scope, entity *entities.Session, fields []string) (err error)
	SQL(fields []string, entity *entities.Session) (string, error)
}

// SessionSearchParams is the search criteria container
type SessionSearchParams struct {
	Secret string
}

// SessionSearch is the DAO search provider interface
type SessionSearch interface {
	Search(scope app.Scope, fields []string, params *SessionSearchParams) (SessionRows, error)
	SQL(fields []string, params *SessionSearchParams) string
}
