package dao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app"
)

// TranslationRows is the result of a query. Its cursor starts before the first row of the result set. Use Next to advance through the rows
type TranslationRows interface {
	Rows
	InjectTo(*entities.Translation) error
	Get() (*entities.Translation, error)
}

// TranslationRow is the result of calling QueryRow to select a single row.
type TranslationRow interface {
	Row
	InjectTo(*entities.Translation) error
	Get() (*entities.Translation, error)
}

// TranslationFindAll is the DAO find all provider interface
type TranslationFindAll interface {
	Find(scope app.Scope, fields []string) (TranslationRows, error)
	SQL(fields []string) (string, error)
}

// TranslationFindByID is the DAO find by id provider interface
type TranslationFindByID interface {
	Find(scope app.Scope, fields []string, id int64) (row TranslationRow, err error)
	SQL(fields []string, id int64) (string, error)
}

// TranslationInsert is the DAO insert provider interface
type TranslationInsert interface {
	Insert(scope app.Scope, entity *entities.Translation) (id int64, err error)
	SQL(entity *entities.Translation) (string, error)
}

// TranslationUpdate is the DAO update provider interface
type TranslationUpdate interface {
	Update(scope app.Scope, entity *entities.Translation, fields []string) (err error)
	SQL(fields []string, entity *entities.Translation) (string, error)
}

// TranslationSearchParams is the search criteria container
type TranslationSearchParams struct {
	Key   string
	Value string
}

// TranslationSearch is the DAO search provider interface
type TranslationSearch interface {
	Search(scope app.Scope, fields []string, params *TranslationSearchParams) (TranslationRows, error)
	SQL(fields []string, params *TranslationSearchParams) string
}
