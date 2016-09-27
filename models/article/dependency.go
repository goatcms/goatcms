package articlemodel

import (
	dep "github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/models"
)

// Factory is a database dependency builder
func Factory(dp dep.Provider) (dep.Instance, error) {
	table := NewArticleTable()
	return NewArticleDAO(table), nil
}

// InitDep initialize a new article model dependency
func InitDep(prov dep.Provider) error {
	if err := prov.AddService(models.ArticleDAOID, Factory); err != nil {
		return err
	}
	return nil
}
