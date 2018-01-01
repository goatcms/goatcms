package database

import (
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	*sqlx.DB
}

func (db *Database) Exec(scope app.Scope, sql string) (err error) {
	var tx *sqlx.Tx
	if tx, err = helpers.TX(scope, db.DB); err != nil {
		return err
	}
	_, err = tx.Exec(sql)
	return err
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
			DB *sqlx.DB `dependency:"db0.engine"`
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
