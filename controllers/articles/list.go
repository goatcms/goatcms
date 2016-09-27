package articles

import (
	"github.com/goatcms/goatcms/models/article"
	"github.com/goatcms/goatcms/services"
)

// ListArticleController is a controler to show a list of article
type ListArticleController struct {
	d *Dependency
}

// NewListArticleController create instance of a list articles controller
func NewListArticleController(d *Dependency) *ListArticleController {
	return &ListArticleController{d}
}

// Get is handler to serve template where one can add new article
func (c *ListArticleController) Get(scope services.RequestScope) {
	database, err := c.d.DP.Database()
	if err != nil {
		scope.Error(err)
		return
	}
	rows, err := c.d.ArticleDAO.FindAll(database.Adapter())
	if err != nil {
		scope.Error(err)
		return
	}
	articleChan := articlemodel.NewArticleChan(scope, rows)
	go articleChan.Go()
	//articles, err := c.d.ArticleDAO.ToEntities(rows)
	if err != nil {
		scope.Error(err)
		return
	}
	err = c.d.TMPL.ExecuteTemplate(scope.Response(), "articles/list", articleChan.Chan)
	if err != nil {
		scope.Error(err)
		return
	}
}
