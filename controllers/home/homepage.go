package home

import (
	"log"
	"net/http"

	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/services"
)

// HomepageController is main page endpoint
type HomepageController struct {
	tmpl services.Template
}

// NewHomepageController create instance of a home controller
func NewHomepageController(dp dependency.Provider) (*HomepageController, error) {
	ctrl := &HomepageController{}
	// load template service from dependency provider
	tmplIns, err := dp.Get(services.TemplateID)
	if err != nil {
		return nil, err
	}
	ctrl.tmpl = tmplIns.(services.Template)
	return ctrl, nil
}

// Show is http get endpoint to serve homepage template
func (c *HomepageController) Show(w http.ResponseWriter, r *http.Request) {
	log.Println("responding to", r.Method, r.URL)
	err := c.tmpl.ExecuteTemplate(w, "homePage", nil)
	if err != nil {
		log.Fatal("error rendering a template: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
