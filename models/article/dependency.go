package articlemodel

import (
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goat-core/http/post"
	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goatcms/models"
)

// DAOFactory is article dao builder
func DAOFactory(dp dependency.Provider) (dependency.Instance, error) {
	table := NewArticleTable()
	return NewArticleDAO(table), nil
}

// TypeFactory is type builder
func TypeFactory(dp dependency.Provider) (dependency.Instance, error) {
	return NewArticleType(), nil
}

// DecoderFactory is decoder builder
func DecoderFactory(dp dependency.Provider) (dependency.Instance, error) {
	typeIns, err := dp.Get(models.ArticleTypeID)
	if err != nil {
		return nil, err
	}
	t := typeIns.(types.CustomType)
	return post.NewDecoder(t), nil
}

// InitDep initialize a new article model dependency
func InitDep(prov dependency.Provider) error {
	if err := prov.AddService(models.ArticleDAOID, DAOFactory); err != nil {
		return err
	}
	if err := prov.AddService(models.ArticleTypeID, TypeFactory); err != nil {
		return err
	}
	if err := prov.AddService(models.ArticleDecoderID, DecoderFactory); err != nil {
		return err
	}
	return nil
}
