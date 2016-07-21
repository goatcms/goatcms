package articles

import "github.com/goatcms/goatcms/services"

// Init initialize the article controllers package
func Init(dp services.Provider) error {
	mux, err := dp.Mux()
	if err != nil {
		return err
	}

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
