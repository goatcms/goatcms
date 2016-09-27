package articles

import (
	"github.com/goatcms/goat-core/db"
	"github.com/goatcms/goatcms/models/article"
	"github.com/goatcms/goatcms/services"
)

// ListCtrl is a controler to show a list of article
type ListCtrl struct {
	d *Dependency
}

// NewListCtrl create instance of a list articles controller
func NewListCtrl(d *Dependency) *ListCtrl {
	return &ListCtrl{d}
}

// Get is handler to serve template where one can add new article
func (c *ListCtrl) Get(scope services.RequestScope) {
	var (
		rows db.Rows
		err  error
	)
	if rows, err = c.d.ArticleDAO.FindAll(c.d.Database.Adapter()); err != nil {
		scope.Error(err)
		return
	}
	articleChan := articlemodel.NewArticleChan(scope, rows)
	if err = c.d.Template.ExecuteTemplate(scope.Response(), "articles/list", articleChan); err != nil {
		scope.Error(err)
		return
	}
}
