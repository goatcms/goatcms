package articlemodel

import (
	"reflect"

	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/abstracttype"
	"github.com/goatcms/goat-core/types/simpletype"
	"github.com/goatcms/goat-core/types/validator"
)

var (
	articleType types.CustomType
)

// ArticleType is type for articles
type ArticleType struct {
	abstracttype.MetaType
	abstracttype.ObjectConverter
	validator.EmptyValidator
}

// NewArticleType create new instance of article type
func newArticleType() types.CustomType {
	var ptr *string
	return &abstracttype.ObjectCustomType{
		SingleCustomType: &ArticleType{
			MetaType: abstracttype.MetaType{
				SQLTypeName:  "text",
				HTMLTypeName: "text",
				GoTypeRef:    reflect.TypeOf(ptr).Elem(),
				Attributes:   map[string]string{},
			},
		},
		Types: map[string]types.CustomType{
			"title":   simpletype.NewTitleType(map[string]string{types.Required: "true"}),
			"content": simpletype.NewContentType(map[string]string{types.Required: "true"}),
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
