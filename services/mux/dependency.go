package mux

import (
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/services"
)

// InitDep initialize a new mux router dependency
func InitDep(prov services.Provider) error {
	if err := prov.AddService(services.MuxID, func(dependency.Provider) (dependency.Instance, error) {
		return NewMux(prov)
	}); err != nil {
		return err
	}
	return nil
}
