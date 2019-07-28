package userc

import (
	"fmt"
	"strings"

	"github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app"
)

// RunConnectAdd execute user:connected:add command
func RunConnectAdd(a app.App, ctxScope app.Scope) (err error) {
	var (
		deps struct {
			Input             app.Input             `dependency:"InputService"`
			Output            app.Output            `dependency:"OutputService"`
			Query             dao.UserSigninQuery   `dependency:"UserSigninQuery"`
			UserConnectInsert dao.UserConnectInsert `dependency:"UserConnectInsert"`
			ServiceName       string                `command:"?service"`
			RemoteIdentyfier  string                `command:"?remote"`
			LocalIdentyfier   string                `command:"?local"`
		}
		user *entities.User
	)
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	if err = ctxScope.InjectTo(&deps); err != nil {
		return err
	}
	if deps.ServiceName == "" {
		return fmt.Errorf("Service is required")
	}
	if deps.LocalIdentyfier == "" {
		return fmt.Errorf("User identyfier (email/username) is required")
	}
	if deps.RemoteIdentyfier == "" {
		return fmt.Errorf("User remote identyfier is required")
	}
	deps.ServiceName = strings.ToLower(deps.ServiceName)
	deps.RemoteIdentyfier = strings.ToLower(deps.RemoteIdentyfier)
	if user, err = deps.Query.Signin(ctxScope, &entities.UserFields{
		ID:       true,
		Username: true,
		Email:    true,
	}, &dao.UserSigninQueryParams{
		Username: deps.LocalIdentyfier,
		Email:    deps.LocalIdentyfier,
	}); err != nil {
		return err
	}
	if _, err = deps.UserConnectInsert.Insert(ctxScope, &entities.UserConnect{
		Service:  &deps.ServiceName,
		RemoteID: &deps.RemoteIdentyfier,
		UserID:   user.ID,
	}); err != nil {
		return err
	}
	if err = ctxScope.Trigger(app.CommitEvent, nil); err != nil {
		return err
	}
	deps.Output.Printf("User %v (%v) connected to %v... success\n", *user.Username, *user.Email, deps.ServiceName)
	return nil
}
