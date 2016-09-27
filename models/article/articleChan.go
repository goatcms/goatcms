package articlemodel

import (
	"github.com/goatcms/goat-core/db/entityChan"
	"github.com/goatcms/goat-core/scope"
	"github.com/goatcms/goatcms/models"
	"github.com/jmoiron/sqlx"
)

func newArticle() interface{} {
	return &models.ArticleEntity{}
}

// NewArticleChan create new article converter chan
func NewArticleChan(scope scope.Scope, rows *sqlx.Rows) *entityChan.ChanCorverter {
	return entityChan.NewChanCorverter(scope, rows, newArticle)
}
