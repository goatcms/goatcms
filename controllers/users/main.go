package users

import (
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/services"
)

// Init initialize the user controller package
func Init(dp dependency.Provider) error {
	muxIns, err := dp.Get(services.MuxID)
	if err != nil {
		return err
	}
	mux := muxIns.(services.Mux)

	ctrl, err := NewUserController(dp)
	if err != nil {
		return err
	}

	mux.Get("/register", ctrl.TemplateSignUp)
	mux.Post("/register", ctrl.TryToSignUp)
	mux.Get("/login", ctrl.TemplateLogin)
	mux.Post("/login", ctrl.TryToLogin)
	mux.Post("/logout", ctrl.TryToLogout)

	return nil
}
