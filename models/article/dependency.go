package articlemodel

import (
	dep "github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/models"
	"github.com/goatcms/goatcms/services"
)

// Factory is a database dependency builder
func Factory(dp dep.Provider) (dep.Instance, error) {
	dbIns, err := dp.Get(services.DBID)
	if err != nil {
		return nil, err
	}
	db := dbIns.(services.Database)

	return NewArticleDAO(db), nil
}

// InitDep initialize a new article model dependency
func InitDep(prov dep.Provider) error {
	if err := prov.AddService(models.ArticleDAOID, Factory); err != nil {
		return err
	}
	return nil
}
