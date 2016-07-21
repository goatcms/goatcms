package home

import "github.com/goatcms/goatcms/services"

// Init initialize the homepage controller package
func Init(dp services.Provider) error {
	mux, err := dp.Mux()
	if err != nil {
		return err
	}

	homepage, err := NewHomepageController(dp)
	if err != nil {
		return err
	}

	mux.Get("/", homepage.Show)

	return nil
}
