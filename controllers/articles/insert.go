package articles

import (
	"net/http"

	"github.com/goatcms/goat-core/db"
	"github.com/goatcms/goat-core/types"
	"github.com/goatcms/goatcms/models"
	"github.com/goatcms/goatcms/services"
)

// InsertCtrl is a controler to create new article
type InsertCtrl struct {
	d *Dependency
}

// NewInsertCtrl create instance of a insert articles controller
func NewInsertCtrl(d *Dependency) *InsertCtrl {
	return &InsertCtrl{d}
}

// Get is handler to serve template where one can add new article
func (c *InsertCtrl) Get(scope services.RequestScope) {
	if err := c.d.Template.ExecuteTemplate(scope.Response(), "articles/new", nil); err != nil {
		scope.Error(err)
		return
	}
}

// Post is handler to save article from form obtained data
func (c *InsertCtrl) Post(scope services.RequestScope) {
	var (
		tx    db.TX
		err   error
		mesgs types.MessageMap
	)
	article := &models.ArticleEntity{}
	if err := c.d.ArticleDecoder.Decode(article, scope.Request()); err != nil {
		scope.Error(err)
		return
	}
	mesgs, err = c.d.ArticleType.Valid(article)
	if err != nil {
		scope.Error(err)
		return
	}
	if len(mesgs.GetAll()) != 0 {
		if err := c.d.Template.ExecuteTemplate(scope.Response(), "articles/new", mesgs); err != nil {
			scope.Error(err)
			return
		}
		return
	}
	if tx, err = scope.TX(); err != nil {
		scope.Error(err)
		return
	}
	c.d.ArticleDAO.Insert(tx, article)
	scope.Commit()
	http.Redirect(scope.Response(), scope.Request(), ListURL, http.StatusSeeOther)
}
