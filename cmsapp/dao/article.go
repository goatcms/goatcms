package dao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app"
)

// ArticleRows is the result of a query. Its cursor starts before the first row of the result set. Use Next to advance through the rows
type ArticleRows interface {
	Rows
	InjectTo(*entities.Article) error
	Get() (*entities.Article, error)
}

// ArticleRow is the result of calling QueryRow to select a single row.
type ArticleRow interface {
	Row
	InjectTo(*entities.Article) error
	Get() (*entities.Article, error)
}

// ArticleFindAll is the DAO find all provider interface
type ArticleFindAll interface {
	Find(scope app.Scope, fields []string) (ArticleRows, error)
	SQL(fields []string) (string, error)
}

// ArticleFindByID is the DAO find by id provider interface
type ArticleFindByID interface {
	Find(scope app.Scope, fields []string, id int64) (row ArticleRow, err error)
	SQL(fields []string, id int64) (string, error)
}

// ArticleInsert is the DAO insert provider interface
type ArticleInsert interface {
	Insert(scope app.Scope, entity *entities.Article) (id int64, err error)
	SQL(entity *entities.Article) (string, error)
}

// ArticleUpdate is the DAO update provider interface
type ArticleUpdate interface {
	Update(scope app.Scope, entity *entities.Article, fields []string) (err error)
	SQL(fields []string, entity *entities.Article) (string, error)
}

// ArticleSearchParams is the search criteria container
type ArticleSearchParams struct {
	Title   string
	Content string
}

// ArticleSearch is the DAO search provider interface
type ArticleSearch interface {
	Search(scope app.Scope, fields []string, params *ArticleSearchParams) (ArticleRows, error)
	SQL(fields []string, params *ArticleSearchParams) string
}
