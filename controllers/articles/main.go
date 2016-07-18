package articles

import (
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/services"
)

// Init initialize the article controllers package
func Init(dp dependency.Provider) error {
	muxIns, err := dp.Get(services.MuxID)
	if err != nil {
		return err
	}
	mux := muxIns.(services.Mux)

	ctrl, err := NewArticleController(dp)
	if err != nil {
		return err
	}

	mux.Get("/article/add", ctrl.AddArticle)
	mux.Post("/article/add", ctrl.SaveArticle)
	mux.Get("/article", ctrl.ListArticle)
	mux.Get("/article/{id:[0-9]+}", ctrl.ViewArticle)

	return nil
}
