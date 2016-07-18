package crypt

import (
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/services"
)

// Factory is a crypt dependency builder
func Factory(dp dependency.Provider) (dependency.Instance, error) {
	return NewCrypt(dp)
}

// InitDep initialize a new crypt dependency
func InitDep(prov dependency.Provider) error {
	if err := prov.AddService(services.CryptID, Factory); err != nil {
		return err
	}
	return nil
}
