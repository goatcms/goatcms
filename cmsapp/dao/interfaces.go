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

// Rows represent a query response
type Rows interface {
	Close() error
	Next() bool
	Columns() ([]string, error)
}

// Row represent a single row query response
type Row interface {
	Columns() ([]string, error)
}
