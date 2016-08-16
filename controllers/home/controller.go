package home

import (
	"log"
	"net/http"

	"github.com/goatcms/goatcms/services"
)

// HomepageController is main page endpoint
type HomepageController struct {
	tmpl services.Template
}

// NewHomepageController create instance of a home controller
func NewHomepageController(dp services.Provider) (*HomepageController, error) {
	var err error
	ctrl := &HomepageController{}
	ctrl.tmpl, err = dp.Template()
	if err != nil {
		return nil, err
	}
	return ctrl, nil
}

// Get is http get endpoint to serve home page
func (c *HomepageController) Get(w http.ResponseWriter, r *http.Request) {
	err := c.tmpl.ExecuteTemplate(w, "home/index", nil)
	if err != nil {
		log.Fatal("error rendering a template: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
