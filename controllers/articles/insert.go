package articles

import (
	"log"
	"net/http"

	"github.com/goatcms/goatcms/models"
	"github.com/gorilla/schema"
)

// InsertArticleController is a controler to create new article
type InsertArticleController struct {
	d *Dependency
}

// NewInsertArticleController create instance of a insert articles controller
func NewInsertArticleController(d *Dependency) *InsertArticleController {
	return &InsertArticleController{d}
}

// Get is handler to serve template where one can add new article
func (c *InsertArticleController) Get(w http.ResponseWriter, r *http.Request) {
	err := c.d.TMPL.ExecuteTemplate(w, "articles/new", nil)
	if err != nil {
		log.Fatal("error rendering a template: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Post is handler to save article from form obtained data
func (c *InsertArticleController) Post(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal("error parsing a form: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	article := &models.ArticleEntity{}
	decoder := schema.NewDecoder()
	err = decoder.Decode(article, r.PostForm)
	if err != nil {
		log.Fatal("error decode a form: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	c.d.ArticleDAO.Insert(article)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
