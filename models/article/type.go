package articlemodel

import (
	"reflect"

	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/abstracttype"
	"github.com/goatcms/goat-core/types/simpletype"
	"github.com/goatcms/goat-core/types/validator"
	"github.com/goatcms/goatcms/models"
)

var (
	articleType types.CustomType
)

// ArticleType is type for articles
type ArticleType struct {
	abstracttype.MetaType
	abstracttype.FilespaceConverter
	validator.ObjectValidator
}

// NewArticleType create new instance of article type
func newArticleType() types.CustomType {
	/*var ptr *string
	types := map[string]types.CustomType{
		"title":   simpletype.NewTitleType(map[string]string{types.Required: "true"}),
		"content": simpletype.NewContentType(map[string]string{types.Required: "true"}),
	}
	return &ArticleType{
		ObjectCustomType: &abstracttype.ObjectCustomType{
			MetaType: abstracttype.MetaType{
				SQLTypeName:  "text",
				HTMLTypeName: "text",
				GoTypeRef:    reflect.TypeOf(ptr).Elem(),
				Attributes:   map[string]string{},
			},
		},
		ObjectValidator: validator.ObjectValidator{
			Types: types,
		},
	}*/
	var ptr *models.ArticleDAO
	types := map[string]types.CustomType{
		"title":   simpletype.NewTitleType(map[string]string{types.Required: "true"}),
		"content": simpletype.NewContentType(map[string]string{types.Required: "true"}),
	}
	return &abstracttype.ObjectCustomType{
		SingleCustomType: &ArticleType{
			MetaType: abstracttype.MetaType{
				SQLTypeName:  "varchar(500)",
				HTMLTypeName: "file",
				GoTypeRef:    reflect.TypeOf(ptr).Elem(),
			},
		},
		Validator: validator.ObjectValidator{
			Types: types,
		},
	}
}

// NewArticleType create new instance of article type
func NewArticleType() types.CustomType {
	if articleType == nil {
		articleType = newArticleType()
	}
	return articleType
}
