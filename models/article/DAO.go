package articlemodel

import (
	"github.com/goatcms/goat-core/db/orm"
	"github.com/goatcms/goatcms/models"
)

// ArticleDAO is describing entity of article
type ArticleDAO struct {
	*orm.BaseDAO
}

// NewArticleDAO create new article DAO
func NewArticleDAO(table *ArticleTable) models.ArticleDAO {
	return &ArticleDAO{
		BaseDAO: &orm.BaseDAO{
			Table: table.BaseTable,
		},
	}
}
