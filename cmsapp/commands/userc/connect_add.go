package userc

import (
	"fmt"

	"github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app"
)

// RunConnectAdd execute user:connected:add command
func RunConnectAdd(a app.App) (err error) {
	var (
		deps struct {
			Query             dao.UserSigninQuery   `dependency:"UserSigninQuery"`
			UserConnectInsert dao.UserConnectInsert `dependency:"UserConnectInsert"`
			ServiceName       string                `argument:"service"`
			RemoteIdentyfier  string                `argument:"remote"`
			LocalIdentyfier   string                `argument:"local"`
		}
		user  *entities.User
		scope = a.AppScope()
	)
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	if deps.ServiceName == "" {
		return fmt.Errorf("Service is required")
	}
	if deps.LocalIdentyfier == "" {
		return fmt.Errorf("User identyfier (email/username) is required")
	}
	if user, err = deps.Query.Signin(scope, &entities.UserFields{
		ID:       true,
		Username: true,
		Email:    true,
	}, &dao.UserSigninQueryParams{
		Username: deps.LocalIdentyfier,
		Email:    deps.LocalIdentyfier,
	}); err != nil {
		return err
	}
	/*if err = deps.Action.SimpleReset(scope, user, deps.Password); err != nil {
		return err
	}*/
	if _, err = deps.UserConnectInsert.Insert(scope, &entities.UserConnect{
		Service:  &deps.ServiceName,
		RemoteID: &deps.RemoteIdentyfier,
		UserID:   user.ID,
	}); err != nil {
		return err
	}
	if err = scope.Trigger(app.CommitEvent, nil); err != nil {
		return err
	}
	fmt.Printf("User %v (%v) connected to %v... success\n", *user.Username, *user.Email, deps.ServiceName)
	return nil
}
