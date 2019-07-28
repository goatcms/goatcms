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
			Input      app.Input                    `dependency:"InputService"`
			Output     app.Output                   `dependency:"OutputService"`
			Action     services.ResetPasswordAction `dependency:"ResetPasswordAction"`
			Query      dao.UserSigninQuery          `dependency:"UserSigninQuery"`
			Identyfier string                       `command:"?by"`
			Password   string                       `command:"?password"`
		}
		user *entities.User
	)
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	if err = ctxScope.InjectTo(&deps); err != nil {
		return err
	}
	if deps.Password == "" {
		return fmt.Errorf("Password is required")
	}
	if deps.Identyfier == "" {
		return fmt.Errorf("User identyfier (email/username) is required")
	}
	if user, err = deps.Query.Signin(ctxScope, &entities.UserFields{
		ID: true,
	}, &dao.UserSigninQueryParams{
		Username: deps.Identyfier,
		Email:    deps.Identyfier,
	}); err != nil {
		return err
	}
	if err = deps.Action.SimpleReset(ctxScope, user, deps.Password); err != nil {
		return err
	}
	if err = ctxScope.Trigger(app.CommitEvent, nil); err != nil {
		return err
	}
	deps.Output.Printf("Password updated... success\n")
	return nil
}
