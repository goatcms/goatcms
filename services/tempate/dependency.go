package template

import (
	dep "github.com/goatcms/goat-core/dependency"
	"github.com/s3c0nDD/goatcms/services"
)

// Factory is a database depondency builder
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
