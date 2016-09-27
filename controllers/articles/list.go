package articles

import (
	"fmt"
	"net/http"

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
		fmt.Println("error rendering a template: ", err)
		http.Error(scope.Response(), err.Error(), http.StatusInternalServerError)
		return
	}
	rows, err := c.d.ArticleDAO.FindAll(database.Adapter())
	if err != nil {
		fmt.Println("error findAll articles: ", err)
		http.Error(scope.Response(), err.Error(), http.StatusInternalServerError)
		return
	}
	articleChan := articlemodel.NewArticleChan(scope, rows)
	go articleChan.Go()
	//articles, err := c.d.ArticleDAO.ToEntities(rows)
	if err != nil {
		fmt.Println("error scan articles rows: ", err)
		http.Error(scope.Response(), err.Error(), http.StatusInternalServerError)
		return
	}
	err = c.d.TMPL.ExecuteTemplate(scope.Response(), "articles/list", articleChan.Chan)
	if err != nil {
		fmt.Println("error rendering a template: ", err)
		http.Error(scope.Response(), err.Error(), http.StatusInternalServerError)
		return
	}
}
