package articles

import (
	"log"
	"net/http"
	"strconv"

	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/models"
	"github.com/goatcms/goatcms/models/article"
	"github.com/goatcms/goatcms/services"
	"github.com/gorilla/mux"
)

// ArticleController is article controller endpoint
type ArticleController struct {
	tmpl       services.Template
	articleDAO models.ArticleDAO
}

// NewArticleController create instance of a articles controller
func NewArticleController(dp dependency.Provider) (*ArticleController, error) {
	ctrl := &ArticleController{}
	// load template service from dependency provider
	tmplIns, err := dp.Get(services.TemplateID)
	if err != nil {
		return nil, err
	}
	ctrl.tmpl = tmplIns.(services.Template)
	// load articleDAO service from dependency provider
	daoIns, err := dp.Get(models.ArticleDAOID)
	if err != nil {
		return nil, err
	}
	ctrl.articleDAO = daoIns.(models.ArticleDAO)
	return ctrl, nil
}

// AddArticle is handler to serve template where one can add new article
func (c *ArticleController) AddArticle(w http.ResponseWriter, r *http.Request) {
	log.Println("responding to", r.Method, r.URL)
	err := c.tmpl.ExecuteTemplate(w, "articles/new", nil)
	if err != nil {
		log.Fatal("error rendering a template: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// SaveArticle is handler to save article from form obtained data
func (c *ArticleController) SaveArticle(w http.ResponseWriter, r *http.Request) {
	log.Println("responding to", r.Method, r.URL)
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

	var articlesToAdd []models.ArticleDTO
	articlesToAdd = append(articlesToAdd, models.ArticleDTO(&article))
	c.articleDAO.PersistAll(articlesToAdd)
	// redirect
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// ListArticle is handler serving template with list of articles
func (c *ArticleController) ListArticle(w http.ResponseWriter, r *http.Request) {
	articles := c.articleDAO.FindAll()

	err := c.tmpl.ExecuteTemplate(w, "articles/list", articles)
	if err != nil {
		log.Fatal("error rendering a template: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ViewArticle is handler serving template with list of articles
func (c *ArticleController) ViewArticle(w http.ResponseWriter, r *http.Request) {
	log.Println("responding to", r.Method, r.URL)
	vars := mux.Vars(r)

	articleID, _ := strconv.Atoi(vars["id"])
	article := c.articleDAO.FindByID(articleID)

	if article == nil { // if fe. user gives id of non existent article
		http.Error(w, http.StatusText(404), 403)
		// TODO maybe handle above some better way?
		return
	}

	err := c.tmpl.ExecuteTemplate(w, "articles/view", article)
	if err != nil {
		log.Fatal("error rendering a template: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
