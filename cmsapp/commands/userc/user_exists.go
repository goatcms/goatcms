package userc

import (
	"github.com/goatcms/goatcms/cmsapp/dao"
	entities "github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app"
)

// RunUserExists exit with zero code if user exist, otherwise exit with non zero code
func RunUserExists(a app.App) (err error) {
	var (
		deps struct {
			Input  app.Input  `dependency:"InputService"`
			Output app.Output `dependency:"OutputService"`

			Search     dao.UserSigninQuery `dependency:"UserSigninQuery"`
			Identyfier string              `argument:"by"`
		}
		user  *entities.User
		scope = a.AppScope()
	)
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	if user, err = deps.Search.Signin(scope, &entities.UserFields{
		ID: true,
	}, &dao.UserSigninQueryParams{
		Username: deps.Identyfier,
		Email:    deps.Identyfier,
	}); err != nil {
		return err
	}
	deps.Output.Printf("User exists: %v\n", *user.ID)
	return nil
}
