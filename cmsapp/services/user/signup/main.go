package signup

import (
	"github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcms/cmsapp/forms"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/goatmail"
	"github.com/goatcms/goatcore/messages"
)

var defaultUserRoles = ""

// SignupAction is global user register service
type SignupAction struct {
	deps struct {
		Crypt    services.Crypt  `dependency:"CryptService"`
		Mailer   services.Mailer `dependency:"MailerService"`
		Logger   services.Logger `dependency:"LoggerService"`
		Inserter dao.UserInsert  `dependency:"UserInsert"`
	}
}

// SignupActionFactory create a user register instance
func SignupActionFactory(dp dependency.Provider) (interface{}, error) {
	ins := &SignupAction{}
	if err := dp.InjectTo(&ins.deps); err != nil {
		return nil, err
	}
	return services.SignupAction(ins), nil
}

func (service *SignupAction) Signup(form *forms.Signup, scope app.Scope) (msgs messages.MessageMap, err error) {
	var (
		hash string
	)
	if msgs, err = forms.ValidSignup(form); err != nil {
		return nil, err
	}
	if len(msgs.GetAll()) > 0 {
		service.deps.Logger.DevLog("UserRegister: valid errors for %v", msgs.GetAll())
		return msgs, nil
	}
	if hash, err = service.deps.Crypt.Hash(form.Password.First); err != nil {
		return nil, err
	}
	user := &entities.User{
		Firstname: form.Firstname,
		Lastname:  form.Lastname,
		Email:     form.Email,
		Username:  form.Username,
		Roles:     &defaultUserRoles,
		Password:  &hash,
	}
	service.deps.Logger.DevLog("UserRegister: %v... insert start", form.Email)
	if _, err = service.deps.Inserter.Insert(scope, user); err != nil {
		return nil, err
	}
	if err = scope.Trigger(app.CommitEvent, nil); err != nil {
		return nil, err
	}
	service.deps.Logger.DevLog("UserRegister: new user (id:%v) %v... start send email", user.ID, form.Email)
	service.deps.Mailer.Send(*form.Email, "register", form, []goatmail.Attachment{}, scope)
	return msgs, nil
}
