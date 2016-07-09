package articles

import (
	"log"
	"net/http"
	"strconv"

	"github.com/goatcms/goat-core/dependency"
	"github.com/gorilla/mux"
	"github.com/goatcms/goatcms/models"
	"github.com/goatcms/goatcms/models/article"
	"github.com/goatcms/goatcms/services"
)

//ArticleController is main page endpoint
type ArticleController struct {
	tmpl       services.Template
	articleDAO models.ArticleDAO
}

//NewArticleController create instance of a articles controller
func NewArticleController(dp dependency.Provider) (*ArticleController, error) {
	ctrl := &ArticleController{}
	// load template service from dependency provider
	tmplIns, err := dp.Get(services.TemplateID)
	if err != nil {
		return nil, err
	}
	ctrl.tmpl = tmplIns.(services.Template)
	// load template service from dependency provider
	daoIns, err := dp.Get(models.ArticleDAOID)
	if err != nil {
		return nil, err
	}
	ctrl.articleDAO = daoIns.(models.ArticleDAO)
	return ctrl, nil
}

// AddArticle is handler which serves template when one can add new article
func (c *ArticleController) AddArticle(w http.ResponseWriter, r *http.Request) {
	log.Println(" AddArticle rendering ")
	err := c.tmpl.ExecuteTemplate(w, "addArticlePage", nil)
	if err != nil {
		log.Fatal("error rendering a template: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// SaveArticle is handler which saves article from form obtained data
func (c *ArticleController) SaveArticle(w http.ResponseWriter, r *http.Request) {
	// TODO: http://www.gorillatoolkit.org/pkg/schema
	// like: err := decoder.Decode(person, r.PostForm)
	// By Sebastian
	err := r.ParseForm()
	if err != nil {
		log.Fatal("error parsing a form: ", err)
	}
	// obtain data from form...
	title := r.PostFormValue("title")
	content := r.PostFormValue("content")
	article := articlemodel.ArticleDTO{Title: title, Content: content}
	// ...and save to database

	var queryArticles []models.ArticleDTO //make slice, append article, send to DB
	queryArticles = append(queryArticles, models.ArticleDTO(&article))
	c.articleDAO.PersistAll(queryArticles)
	//redirect
	// log.Println("DB: add article [ " + title + " ] = " + content)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// ListArticle - handler serving template with list of articles
func (c *ArticleController) ListArticle(w http.ResponseWriter, r *http.Request) {
	articles := c.articleDAO.GetAll()

	err := c.tmpl.ExecuteTemplate(w, "articleListPage", articles)
	if err != nil {
		log.Fatal("error rendering a template: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ViewArticle - handler serving template with list of articles
func (c *ArticleController) ViewArticle(w http.ResponseWriter, r *http.Request) {
	log.Println("ArticleViewHandler: responding to GET", r.URL)
	vars := mux.Vars(r)

	articleID, _ := strconv.Atoi(vars["id"])
	article := c.articleDAO.GetOne(articleID)

	if article == nil { // if fe. user gives id of non existent article
		http.Error(w, http.StatusText(404), 403)
		return
	}

	err := c.tmpl.ExecuteTemplate(w, "articleViewPage", article)
	if err != nil {
		log.Fatal("error rendering a template: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
