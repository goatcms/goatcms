package dao

import "github.com/goatcms/goatcore/app"

// Database provide database public plain api
type Database interface {
	Exec(scope app.Scope, query string) error
	Commit(scope app.Scope) error
	Rollback(scope app.Scope) error
}

// CreateTable describe create table action as DAO separated service
type CreateTable interface {
	CreateTable(scope app.Scope) error
	SQL() string
}

// Delete describe delete action as DAO separated service
type Delete interface {
	Delete(scope app.Scope, id int64) error
	SQL(where string) string
}

// DropTable describe drop action as DAO separated service
type DropTable interface {
	DropTable(scope app.Scope) error
	SQL() string
}

// FindAll describe find all action as DAO separated service
type FindAll interface {
	Find(scope app.Scope, fields []string) (Rows, error)
	SQL(fields []string) (string, error)
}

// FindByID describe find by id action as DAO separated service
type FindByID interface {
	Find(scope app.Scope, fields []string, id int64) (row Row, err error)
	SQL(fields []string, id int64) (string, error)
}

// Insert describe insert action as DAO separated service
type Insert interface {
	Insert(scope app.Scope, entity interface{}, fields []string) (id int64, err error)
	SQL(fields []string) (string, error)
}

// Update describe update action as DAO separated service
type Update interface {
	Update(scope app.Scope, entity interface{}, fields []string) (err error)
	SQL(fields []string) (string, error)
}

// TX represent a database transaction accessor
/*type TX interface {
	Queryx(query string, args ...interface{}) (Rows, error)
	QueryRowx(query string, args ...interface{}) (Row, error)
	NamedExec(query string, arg interface{}) (sql.Result, error)
	MustExec(query string, args ...interface{}) sql.Result
	Commit() error
	Rollback() error
}*/

// Rows represent a query response
type Rows interface {
	Close() error
	Next() bool
	Columns() ([]string, error)
	StructScan(dest interface{}) error
}

// Row represent a single row query response
type Row interface {
	Scan(...interface{}) error
	StructScan(interface{}) error
	Columns() ([]string, error)
	Err() error
}
