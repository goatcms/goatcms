package userc

import (
	"fmt"

	"github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
)

// RunUpdatePassword execute user:password:update command
func RunUpdatePassword(a app.App, ctxScope app.Scope) (err error) {
	var (
		deps struct {
			Router     services.Router              `dependency:"RouterService"`
			Action     services.ResetPasswordAction `dependency:"ResetPasswordAction"`
			Query      dao.UserSigninQuery          `dependency:"UserSigninQuery"`
			Identyfier string                       `argument:"by"`
			Password   string                       `argument:"password"`
		}
		user  *entities.User
		scope = a.AppScope()
	)
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	if deps.Password == "" {
		return fmt.Errorf("Password is required")
	}
	if deps.Identyfier == "" {
		return fmt.Errorf("User identyfier (email/username) is required")
	}
	if user, err = deps.Query.Signin(scope, &entities.UserFields{
		ID: true,
	}, &dao.UserSigninQueryParams{
		Username: deps.Identyfier,
		Email:    deps.Identyfier,
	}); err != nil {
		return err
	}
	if err = deps.Action.SimpleReset(scope, user, deps.Password); err != nil {
		return err
	}
	if err = scope.Trigger(app.CommitEvent, nil); err != nil {
		return err
	}
	fmt.Printf("Password updated... success\n")
	return nil
}
