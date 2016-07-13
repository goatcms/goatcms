package auth

import (
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/services"
)

// Factory is a authentification dependency builder
func Factory(dp dependency.Provider) (dependency.Instance, error) {
	return NewAuth(dp)
}

// InitDep initialize a new authentification dependency
func InitDep(prov dependency.Provider) error {
	if err := prov.AddService(services.AuthID, Factory); err != nil {
		return err
	}
	return nil
}
