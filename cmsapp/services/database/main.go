package database

import (
	"fmt"

	"github.com/goatcms/goat-core/app"
	"github.com/goatcms/goat-core/db"
	"github.com/goatcms/goat-core/db/adapter"
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // import sqlite3 adapter as default db driver
)

// Database is global database connection provider
type Database struct {
	deps struct {
		DependencyScope app.Scope `dependency:"DependencyScope"`
		Engine          string    `config:"?database.engine"`
		Url             string    `config:"?database.url"`
	}
	instance *sqlx.DB
	flushTX  db.TX
}

// DatabaseFactory create a database instance
func DatabaseFactory(dp dependency.Provider) (interface{}, error) {
	db := &Database{
		instance: nil,
	}
	if err := dp.InjectTo(&db.deps); err != nil {
		return nil, err
	}
	if err := db.Open(); err != nil {
		return nil, err
	}
	db.deps.DependencyScope.On(app.BeforeCloseEvent, func(interface{}) error {
		return db.Close()
	})
	return services.Database(db), nil
}

// Open initialize a database connection
func (db *Database) Open() error {
	var err error
	db.instance, err = sqlx.Open(db.engine(), db.url())
	if err != nil {
		return err
	}
	if db.instance == nil {
		return fmt.Errorf("Can not open database conection (pointer is nil)")
	}
	db.flushTX = adapter.NewTXFromDB(db.instance)
	return nil
}

// Close close a connection
func (db *Database) Close() error {
	return db.instance.Close()
}

// TX return new transaction
func (db *Database) TX() (db.TX, error) {
	xtx, err := db.instance.Beginx()
	if err != nil {
		return nil, err
	}
	return adapter.NewTX(xtx), nil
}

// TX return new flush transaction (all queries are run directly from database)
func (db *Database) FlushTX() (db.TX, error) {
	return db.flushTX, nil
}

// Adapter return current golang database interface instance
func (db *Database) Adapter() *sqlx.DB {
	return db.instance
}

// engine return database engine name
func (db *Database) engine() string {
	if db.deps.Engine != "" {
		return db.deps.Engine
	}
	return services.DefaultDatabaseEngine
}

// engine return database engine name
func (db *Database) url() string {
	if db.deps.Url != "" {
		return db.deps.Url
	}
	return services.DefaultDatabaseUrl
}
