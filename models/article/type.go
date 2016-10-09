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
	articleTypeInstance types.CustomType
)

// newArticleType create new instance of article type
func newArticleType() types.CustomType {
	var ptr *models.ArticleDAO
	goTypeRef := reflect.TypeOf(ptr).Elem()
	types := map[string]types.CustomType{
		"Title":   simpletype.NewTitleType(map[string]string{types.Required: "true"}),
		"Content": simpletype.NewContentType(map[string]string{types.Required: "true"}),
	}
	return &abstracttype.ObjectCustomType{
		MetaType: &abstracttype.MetaType{
			SQLTypeName:  "varchar(500)",
			HTMLTypeName: "file",
			GoTypeRef:    goTypeRef,
		},
		TypeConverter: abstracttype.NewObjectConverterFromType(goTypeRef),
		TypeValidator: validator.NewObjectValidator(types),
		Types:         types,
	}
}

// NewArticleType create new instance of article type
func NewArticleType() types.CustomType {
	if articleTypeInstance == nil {
		articleTypeInstance = newArticleType()
	}
	return articleTypeInstance
}
