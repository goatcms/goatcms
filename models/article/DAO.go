package articlemodel

import (
	"github.com/goatcms/goat-core/db/orm"
	"github.com/goatcms/goatcms/models"
	"github.com/jmoiron/sqlx"
)

// ArticleDAO is describing entity of article
type ArticleDAO struct {
	*orm.BaseDAO
}

// NewArticleDAO create new article DAO
func NewArticleDAO(db *sqlx.DB, table *ArticleTable) models.ArticleDAO {
	return &ArticleDAO{
		BaseDAO: orm.NewBaseDAO(table.BaseTable, db),
	}
}

// ToEntities read a query to entity list
func (dao *ArticleDAO) ToEntities(rows *sqlx.Rows) ([]*models.ArticleEntity, error) {
	var entities = []*models.ArticleEntity{}
	for rows.Next() {
		entity := &models.ArticleEntity{}
		if err := rows.StructScan(entity); err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}
	return entities, nil
}
