package services

import "github.com/goatcms/goat-core/dependency"

// InitDep initialize a new mux router dependency
func InitDep(prov Provider) error {
	if err := prov.AddService(ProviderID, func(dependency.Provider) (dependency.Instance, error) {
		return Provider(prov), nil
	}); err != nil {
		return err
	}
	return nil
}
