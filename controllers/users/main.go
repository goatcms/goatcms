package users

import "github.com/goatcms/goatcms/services"

// Init initialize the user controller package
func Init(dp services.Provider) error {
	mux, err := dp.Mux()
	if err != nil {
		return err
	}

	ctrl, err := NewUserController(dp)
	if err != nil {
		return err
	}

	mux.Get("/register", ctrl.TemplateSignUp)
	mux.Post("/register", ctrl.TryToSignUp)
	mux.Get("/login", ctrl.TemplateLogin)
	mux.Post("/login", ctrl.TryToLogin)
	mux.Get("/logout", ctrl.TryToLogout)

	return nil
}
