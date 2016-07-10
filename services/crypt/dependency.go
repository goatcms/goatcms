package crypt

import (
	dep "github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/services"
)

//Factory is a database dependency builder
func Factory(dp dep.Provider) (dep.Instance, error) {
	return NewCrypt(dp)
}

//InitDep inicjalize a new database dependency
func InitDep(prov dep.Provider) error {
	if err := prov.AddService(services.CryptID, Factory); err != nil {
		return err
	}
	return nil
}
