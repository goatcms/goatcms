package imagemodel

import (
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/models"
	"github.com/goatcms/goatcms/services"
)

// Factory is an image model dependency builder
func Factory(dp dependency.Provider) (dependency.Instance, error) {
	dbIns, err := dp.Get(services.DBID)
	if err != nil {
		return nil, err
	}
	db := dbIns.(services.Database)
	return NewImageDAO(db)
}

// InitDep initialize a new image model dependency
func InitDep(prov dependency.Provider) error {
	if err := prov.AddService(models.ImageDAOID, Factory); err != nil {
		return err
	}
	return nil
}
