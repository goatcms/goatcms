package dao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app"
)

// UserRows is the result of a query. Its cursor starts before the first row of the result set. Use Next to advance through the rows
type UserRows interface {
	Rows
	InjectTo(*entities.User) error
	Get() (*entities.User, error)
}

// UserRow is the result of calling QueryRow to select a single row.
type UserRow interface {
	Row
	InjectTo(*entities.User) error
	Get() (*entities.User, error)
}

// UserFindAll is the DAO find all provider interface
type UserFindAll interface {
	Find(scope app.Scope, fields []string) (UserRows, error)
	SQL(fields []string) (string, error)
}

// UserFindByID is the DAO find by id provider interface
type UserFindByID interface {
	Find(scope app.Scope, fields []string, id int64) (row UserRow, err error)
	SQL(fields []string, id int64) (string, error)
}

// UserInsert is the DAO insert provider interface
type UserInsert interface {
	Insert(scope app.Scope, entity *entities.User) (id int64, err error)
	SQL(entity *entities.User) (string, error)
}

// UserUpdate is the DAO update provider interface
type UserUpdate interface {
	Update(scope app.Scope, entity *entities.User, fields []string) (err error)
	SQL(fields []string, entity *entities.User) (string, error)
}

// UserSearchParams is the search criteria container
type UserSearchParams struct {
	Firstname string
	Login     string
	Email     string
	Password  string
}

// UserSearch is the DAO search provider interface
type UserSearch interface {
	Search(scope app.Scope, fields []string, params *UserSearchParams) (UserRows, error)
	SQL(fields []string, params *UserSearchParams) string
}
