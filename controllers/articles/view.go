package articles

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/goatcms/goatcms/models"
	"github.com/goatcms/goatcms/services"
	"github.com/gorilla/mux"
)

// ViewArticleController is a controler to show a single article
type ViewArticleController struct {
	d *Dependency
}

// NewViewArticleController create instance of a list articles controller
func NewViewArticleController(d *Dependency) *ViewArticleController {
	return &ViewArticleController{d}
}

// Get is handler to serve template to view a article
func (c *ViewArticleController) Get(scope services.RequestScope) {
	vars := mux.Vars(scope.Request())
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Fatal("parse int fail: ", err)
		http.Error(scope.Response(), err.Error(), http.StatusInternalServerError)
		return
	}
	database, err := c.d.DP.Database()
	if err != nil {
		fmt.Println("error rendering a template: ", err)
		http.Error(scope.Response(), err.Error(), http.StatusInternalServerError)
		return
	}
	row := c.d.ArticleDAO.FindByID(database.Adapter(), id)
	if row.Err() != nil {
		log.Fatal("find article fail: ", row.Err())
		http.Error(scope.Response(), row.Err().Error(), http.StatusInternalServerError)
		return
	}
	article := &models.ArticleEntity{}
	row.StructScan(article)
	err = c.d.TMPL.ExecuteTemplate(scope.Response(), "articles/view", article)
	if err != nil {
		log.Fatal("error rendering a template: ", err)
		http.Error(scope.Response(), err.Error(), http.StatusInternalServerError)
		return
	}
}
