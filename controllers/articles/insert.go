package articles

import (
	"fmt"
	"log"
	"net/http"

	"github.com/goatcms/goat-core/http/post"
	"github.com/goatcms/goatcms/models"
	"github.com/goatcms/goatcms/services"
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
func (c *InsertArticleController) Get(scope services.RequestScope) {
	err := c.d.TMPL.ExecuteTemplate(scope.Response(), "articles/new", nil)
	if err != nil {
		log.Fatal("error rendering a template: ", err)
		http.Error(scope.Response(), err.Error(), http.StatusInternalServerError)
		return
	}
}

// Post is handler to save article from form obtained data
func (c *InsertArticleController) Post(scope services.RequestScope) {
	/*err := r.ParseForm()
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
	http.Redirect(w, r, "/", http.StatusSeeOther)*/

	article := &models.ArticleEntity{}
	decoder := post.NewDecoder(c.d.ArticleType)
	if err := decoder.Decode(article, scope.Request()); err != nil {
		fmt.Println(err)
		return
	}
	tx, err := scope.TX()
	if err != nil {
		log.Fatal("error rendering a template: ", err)
		http.Error(scope.Response(), err.Error(), http.StatusInternalServerError)
		return
	}
	c.d.ArticleDAO.Insert(tx, article)
	tx.Commit()
	http.Redirect(scope.Response(), scope.Request(), "/", http.StatusSeeOther)
}
