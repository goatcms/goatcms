package useremodel

import (
	dep "github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/models"
	"github.com/goatcms/goatcms/services"
)

// Factory is a database depondency builder
func Factory(dp dep.Provider) (dep.Instance, error) {
	dbIns, err := dp.Get(services.DBID)
	if err != nil {
		return nil, err
	}
	db := dbIns.(services.Database)
	return NewUserDAO(db)
}

// InitDep inicjalize a new database dependency
func InitDep(prov dep.Provider) error {
	if err := prov.AddService(models.UserDAOID, Factory); err != nil {
		return err
	}
	return nil
}
