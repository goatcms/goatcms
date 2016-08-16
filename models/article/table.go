package articlemodel

import (
	"github.com/goatcms/goat-core/db/orm"
	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goat-core/types/simpletype"
)

// ArticleTable describe table structure and simple queries
type ArticleTable struct {
	*orm.BaseTable
}

// NewArticleTable create new article table
func NewArticleTable() *ArticleTable {
	return &ArticleTable{
		BaseTable: orm.NewBaseTable("articles", map[string]types.CustomType{
			"title":   simpletype.NewTitleType([]string{types.NotNull}),
			"content": simpletype.NewContentType([]string{types.NotNull}),
			"image":   simpletype.NewImageType([]string{}),
		}),
	}
}
