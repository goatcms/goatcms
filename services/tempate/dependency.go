package template

import (
	dep "github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/services"
)

// Factory is a database dependency builder
func Factory(dp dep.Provider) (dep.Instance, error) {
	return NewTemplate()
}

// InitDep inicjalize a new database dependency
func InitDep(prov dep.Provider) error {
	if err := prov.AddService(services.TemplateID, Factory); err != nil {
		return err
	}
	return nil
}
