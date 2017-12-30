package database

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// Factory create a new database instance
func Factory(dp dependency.Provider) (interface{}, error) {
	var (
		deps struct {
			DependencyScope app.Scope `dependency:"DependencyScope"`
			URL             string    `config:"?database.url"`
		}
		db  *sqlx.DB
		err error
	)
	if err = dp.InjectTo(&deps); err != nil {
		return nil, err
	}
	if db, err = sqlx.Open("sqlite3", deps.URL); err != nil {
		return nil, err
	}
	deps.DependencyScope.On(app.BeforeCloseEvent, func(interface{}) error {
		return db.Close()
	})
	return db, nil
}
