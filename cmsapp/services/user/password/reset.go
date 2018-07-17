package password

import (
	"fmt"

	"github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

var defaultUserRoles = ""

// ResetPasswordAction is global user reset password service
type ResetPasswordAction struct {
	deps struct {
		Crypt   services.Crypt  `dependency:"CryptService"`
		Logger  services.Logger `dependency:"LoggerService"`
		Updater dao.UserUpdate  `dependency:"UserUpdate"`
	}
}

// ResetPasswordActionFactory create a user reset password instance
func ResetPasswordActionFactory(dp dependency.Provider) (interface{}, error) {
	ins := &ResetPasswordAction{}
	if err := dp.InjectTo(&ins.deps); err != nil {
		return nil, err
	}
	return services.ResetPasswordAction(ins), nil
}

// SimpleReset set new password for user
func (service *ResetPasswordAction) SimpleReset(scope app.Scope, user *entities.User, password string) (err error) {
	var (
		hash string
	)
	if user == nil {
		return fmt.Errorf("[Code Error] SimpleReset: user is required")
	}
	if hash, err = service.deps.Crypt.Hash(password); err != nil {
		return err
	}
	user.Password = &hash
	if err = service.deps.Updater.Update(scope, user, &entities.UserFields{
		Password: true,
	}); err != nil {
		return err
	}
	if err = scope.Trigger(app.CommitEvent, nil); err != nil {
		return err
	}
	service.deps.Logger.DevLog("SimpleReset: user password reseted... success")
	return nil
}
