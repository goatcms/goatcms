package articlemodel

import (
	dep "github.com/goatcms/goat-core/dependency"
	"github.com/s3c0nDD/goatcms/models"
	"github.com/s3c0nDD/goatcms/services"
)

// Factory is a database depondency builder
func Factory(dp dep.Provider) (dep.Instance, error) {
	dbIns, err := dp.Get(services.DBID)
	if err != nil {
		return nil, err
	}
	db := dbIns.(services.Database)

	return NewArticleDAO(db), nil
}

// InitDep inicjalize a new database dependency
func InitDep(prov dep.Provider) error {
	if err := prov.AddService(models.ArticleDAOID, Factory); err != nil {
		return err
	}
	return nil
}
