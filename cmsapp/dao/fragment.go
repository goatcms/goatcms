package dao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app"
)

// FragmentRows is the result of a query. Its cursor starts before the first row of the result set. Use Next to advance through the rows
type FragmentRows interface {
	Rows
	InjectTo(*entities.Fragment) error
	Get() (*entities.Fragment, error)
}

// FragmentRow is the result of calling QueryRow to select a single row.
type FragmentRow interface {
	Row
	InjectTo(*entities.Fragment) error
	Get() (*entities.Fragment, error)
}

// FragmentFindAll is the DAO find all provider interface
type FragmentFindAll interface {
	Find(scope app.Scope, fields []string) (FragmentRows, error)
	SQL(fields []string) (string, error)
}

// FragmentFindByID is the DAO find by id provider interface
type FragmentFindByID interface {
	Find(scope app.Scope, fields []string, id int64) (user *entities.Fragment, err error)
	SQL(fields []string, id int64) (string, error)
}

// FragmentInsert is the DAO insert provider interface
type FragmentInsert interface {
	Insert(scope app.Scope, entity *entities.Fragment) (id int64, err error)
	SQL(entity *entities.Fragment) (string, error)
}

// FragmentUpdate is the DAO update provider interface
type FragmentUpdate interface {
	Update(scope app.Scope, entity *entities.Fragment, fields []string) (err error)
	SQL(fields []string, entity *entities.Fragment) (string, error)
}

// FragmentSearchParams is the search criteria container
type FragmentSearchParams struct {
	Lang    string
	Name    string
	Content string
}

// FragmentSearch is the DAO search provider interface
type FragmentSearch interface {
	Search(scope app.Scope, fields []string, params *FragmentSearchParams) (FragmentRows, error)
	SQL(fields []string, params *FragmentSearchParams) string
}
