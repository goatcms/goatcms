package userc

import (
	"fmt"
	"strings"

	"github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app"
)

// RunUpdateRoles execute user:roles:update command
func RunUpdateRoles(a app.App, ctxScope app.Scope) (err error) {
	var (
		deps struct {
			Input      app.Input           `dependency:"InputService"`
			Output     app.Output          `dependency:"OutputService"`
			Search     dao.UserSigninQuery `dependency:"UserSigninQuery"`
			Updater    dao.UserUpdate      `dependency:"UserUpdate"`
			Identyfier string              `command:"?by"`
			Roles      string              `command:"?roles"`
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
	if deps.Roles == "" {
		return fmt.Errorf("User roles are required")
	}
	if user, err = deps.Search.Signin(ctxScope, &entities.UserFields{
		ID: true,
	}, &dao.UserSigninQueryParams{
		Username: deps.Identyfier,
		Email:    deps.Identyfier,
	}); err != nil {
		return err
	}
	deps.Roles = strings.Replace(deps.Roles, "&", " ", -1)
	user.Roles = &deps.Roles
	if err = deps.Updater.Update(ctxScope, user, &entities.UserFields{
		Roles: true,
	}); err != nil {
		return err
	}
	if err = ctxScope.Trigger(app.CommitEvent, nil); err != nil {
		return err
	}
	deps.Output.Printf("Role updated... success\n")
	return nil
}
