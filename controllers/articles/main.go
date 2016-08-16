package articles

import (
	"github.com/goatcms/goatcms/models"
	"github.com/goatcms/goatcms/services"
)

// Dependency is default set of dependency
type Dependency struct {
	TMPL       services.Template
	ArticleDAO models.ArticleDAO
}

// NewDependency is default set of dependency
func NewDependency(dp services.Provider) (*Dependency, error) {
	var (
		d   = &Dependency{}
		err error
	)
	d.TMPL, err = dp.Template()
	if err != nil {
		return nil, err
	}
	d.ArticleDAO, err = dp.ArticleDAO()
	if err != nil {
		return nil, err
	}
	return d, nil
}

// Init initialize the article controllers package
func Init(dp services.Provider) error {
	d, err := NewDependency(dp)
	if err != nil {
		return err
	}
	insertCtrl := NewInsertArticleController(d)
	listCtrl := NewListArticleController(d)
	viewCtrl := NewViewArticleController(d)

	mux, err := dp.Mux()
	if err != nil {
		return err
	}

	mux.Get("/article/add", insertCtrl.Get)
	mux.Post("/article/add", insertCtrl.Post)
	mux.Get("/article", listCtrl.Get)
	mux.Get("/article/{id:[0-9]+}", viewCtrl.Get)

	return nil
}
