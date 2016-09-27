package articlemodel

import "github.com/goatcms/goat-core/db/orm"

// ArticleTable describe table structure and simple queries
type ArticleTable struct {
	*orm.BaseTable
}

// NewArticleTable create new article table
func NewArticleTable() *ArticleTable {
	return &ArticleTable{
		BaseTable: orm.NewBaseTable("articles", NewArticleType().GetSubTypes()),
	}
}
