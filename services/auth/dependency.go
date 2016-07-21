package auth

import (
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/services"
)

// Factory is a authentification dependency builder
func Factory(dp services.Provider) (dependency.Instance, error) {
	return NewAuth(dp)
}

// InitDep initialize a new authentification dependency
func InitDep(prov services.Provider) error {
	err := prov.AddService(services.AuthID, func(dependency.Provider) (dependency.Instance, error) {
		return NewAuth(prov)
	})
	if err != nil {
		return err
	}
	return nil
}
