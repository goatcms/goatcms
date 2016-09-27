package articles

import (
	"strconv"

	"github.com/goatcms/goatcms/models"
	"github.com/goatcms/goatcms/services"
	"github.com/gorilla/mux"
)

// ViewCtrl is a controler to show a single article
type ViewCtrl struct {
	d *Dependency
}

// NewViewCtrl create instance of a list articles controller
func NewViewCtrl(d *Dependency) *ViewCtrl {
	return &ViewCtrl{d}
}

// Get is handler to serve template to view a article
func (c *ViewCtrl) Get(scope services.RequestScope) {
	vars := mux.Vars(scope.Request())
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		scope.Error(err)
		return
	}
	row := c.d.ArticleDAO.FindByID(c.d.Database.Adapter(), id)
	if row.Err() != nil {
		scope.Error(err)
		return
	}
	article := &models.ArticleEntity{}
	row.StructScan(article)
	err = c.d.Template.ExecuteTemplate(scope.Response(), "articles/view", article)
	if err != nil {
		scope.Error(err)
		return
	}
}
