package userc

import (
	"fmt"

	"github.com/goatcms/goatcms/cmsapp/dao"
	entities "github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app"
)

// RunUserExists exit with zero code if user exist, otherwise exit with non zero code
func RunUserExists(a app.App, ctxScope app.Scope) (err error) {
	var (
		deps struct {
			Input      app.Input           `dependency:"InputService"`
			Output     app.Output          `dependency:"OutputService"`
			Search     dao.UserSigninQuery `dependency:"UserSigninQuery"`
			Identyfier string              `command:"by"`
		}
		user *entities.User
	)
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	if err = ctxScope.InjectTo(&deps); err != nil {
		return err
	}
	if deps.Identyfier == "" {
		return fmt.Errorf("User identyfier (email/username) is required")
	}
	if user, err = deps.Search.Signin(ctxScope, &entities.UserFields{
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
