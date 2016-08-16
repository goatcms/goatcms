package files

import (
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/services"
)

// Factory is a files dependency builder
func Factory(dp services.Provider) (dependency.Instance, error) {
	return NewFiles(dp)
}

// InitDep initialize a new files dependency
func InitDep(prov dependency.Provider) error {
	if err := prov.AddService(services.FilesID, Factory); err != nil {
		return err
	}
	return nil
}
