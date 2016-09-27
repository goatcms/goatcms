package usermodel

import (
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/models"
	"github.com/goatcms/goatcms/services"
)

// Factory is a user model dependency builder
func Factory(dp dependency.Provider) (dependency.Instance, error) {
	dbIns, err := dp.Get(services.DBID)
	if err != nil {
		return nil, err
	}
	db := dbIns.(services.Database)
	return NewUserDAO(db)
}

// InitDep initialize a new user model dependency
func InitDep(prov dependency.Provider) error {
	if err := prov.AddService(models.UserDAOID, Factory); err != nil {
		return err
	}
	return nil
}
