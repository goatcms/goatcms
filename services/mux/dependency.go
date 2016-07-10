package mux

import (
	dep "github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/services"
)

const (
	defaultDatabasePath = "sqlite3.db"
)

// Factory is a mux router dependency builder
func Factory(dp dep.Provider) (dep.Instance, error) {
	return NewMux()
}

// InitDep initialize a new mux router dependency
func InitDep(prov dep.Provider) error {
	if err := prov.AddService(services.MuxID, Factory); err != nil {
		return err
	}
	return nil
}
