package home

import (
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/services"
)

// Init initialize the homepage controller package
func Init(dp dependency.Provider) error {
	muxIns, err := dp.Get(services.MuxID)
	if err != nil {
		return err
	}
	mux := muxIns.(services.Mux)

	homepage, err := NewHomepageController(dp)
	if err != nil {
		return err
	}

	mux.Get("/", homepage.Show)

	return nil
}
