package images

import (
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/services"
)

// Init initialize the images controller package
func Init(dp dependency.Provider) error {
	muxIns, err := dp.Get(services.MuxID)
	if err != nil {
		return err
	}
	mux := muxIns.(services.Mux)

	ctrl, err := NewImageController(dp)
	if err != nil {
		return err
	}

	mux.Get("/image/add/{id:[0-9]+}", ctrl.TemplateAddImage)
	mux.Post("/image/add/{id:[0-9]+}", ctrl.TrySaveImage)

	return nil
}
