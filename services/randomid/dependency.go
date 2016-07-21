package randomid

import (
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/services"
)

// Factory is a random id dependency builder
func Factory(dp dependency.Provider) (dependency.Instance, error) {
	return NewRandomID(dp)
}

// InitDep initialize a new random id dependency
func InitDep(prov dependency.Provider) error {
	if err := prov.AddService(services.RandomidID, Factory); err != nil {
		return err
	}
	return nil
}
