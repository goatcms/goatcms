package articles

import (
	"log"
	"net/http"
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
func (c *ListArticleController) Get(w http.ResponseWriter, r *http.Request) {
	rows, err := c.d.ArticleDAO.FindAll()
	if err != nil {
		log.Fatal("error findAll articles: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	articles, err := c.d.ArticleDAO.ToEntities(rows)
	if err != nil {
		log.Fatal("error scan articles rows: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = c.d.TMPL.ExecuteTemplate(w, "articles/list", articles)
	if err != nil {
		log.Fatal("error rendering a template: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
