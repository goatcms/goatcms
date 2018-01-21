package userc

import (
	"fmt"
	"strings"

	"github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
)

// RunUpdateRoles execute user:roles:update command
func RunUpdateRoles(a app.App) (err error) {
	var (
		deps struct {
			Router     services.Router     `dependency:"RouterService"`
			Search     dao.UserSigninQuery `dependency:"UserSigninQuery"`
			Updater    dao.UserUpdate      `dependency:"UserUpdate"`
			Identyfier string              `argument:"by"`
			Roles      string              `argument:"roles"`
		}
		user  *entities.User
		scope = a.AppScope()
	)
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	if user, err = deps.Search.Signin(scope, []string{"ID"}, &dao.UserSigninQueryParams{
		Username: deps.Identyfier,
		Email:    deps.Identyfier,
	}); err != nil {
		return err
	}
	deps.Roles = strings.Replace(deps.Roles, "&", " ", -1)
	user.Roles = &deps.Roles
	if err = deps.Updater.Update(scope, user, []string{"Roles"}); err != nil {
		return err
	}
	if err = scope.Trigger(app.CommitEvent, nil); err != nil {
		return err
	}
	fmt.Printf("Role updated... success\n")
	return nil
}
