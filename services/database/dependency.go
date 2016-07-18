package database

import (
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/services"
)

const (
	defaultDatabasePath = "sqlite3.db"
)

// Factory is a database dependency builder
func Factory(dp dependency.Provider) (dependency.Instance, error) {
	return NewDatabase(defaultDatabasePath)
}

// InitDep initialize a new database dependency
func InitDep(prov dependency.Provider) error {
	if err := prov.AddService(services.DBID, Factory); err != nil {
		return err
	}
	return nil
}
