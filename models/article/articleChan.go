package articlemodel

import (
	"github.com/goatcms/goat-core/db"
	"github.com/goatcms/goat-core/db/entityChan"
	"github.com/goatcms/goat-core/scope"
	"github.com/goatcms/goatcms/models"
)

func newArticle() interface{} {
	return &models.ArticleEntity{}
}

// NewArticleChan create new article converter chan
func NewArticleChan(scope scope.Scope, rows db.Rows) entityChan.EntityChan {
	converter := entityChan.NewChanCorverter(scope, rows, newArticle)
	go converter.Go()
	return converter.Chan
}
