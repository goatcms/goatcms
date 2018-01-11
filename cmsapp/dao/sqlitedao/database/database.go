package database

import (
	"database/sql"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

type Database struct {
	*sql.DB
}

func (db *Database) Exec(scope app.Scope, query string) (err error) {
	var tx *sql.Tx
	if tx, err = helpers.TX(scope, db.DB); err != nil {
		return err
	}
	_, err = tx.Exec(query)
	return err
}

func (db *Database) Query(scope app.Scope, query string) (rows maindef.Rows, err error) {
	var (
		tx      *sql.Tx
		sqlRows *sql.Rows
	)
	if tx, err = helpers.TX(scope, db.DB); err != nil {
		return nil, err
	}
	if sqlRows, err = tx.Query(query); err != nil {
		return nil, err
	}
	return helpers.NewRows(sqlRows), nil
}

func (db *Database) Commit(scope app.Scope) (err error) {
	_, err = helpers.Commit(scope)
	return err
}

func (db *Database) Rollback(scope app.Scope) (err error) {
	_, err = helpers.Rollback(scope)
	return err
}

// Factory create a new database instance
func Factory(dp dependency.Provider) (interface{}, error) {
	var (
		deps struct {
			DB *sql.DB `dependency:"db0.engine"`
		}
		err error
	)
	if err = dp.InjectTo(&deps); err != nil {
		return nil, err
	}
	return maindef.Database(&Database{
		DB: deps.DB,
	}), nil
}
