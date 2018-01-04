package register

import (
	"github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcms/cmsapp/forms/user/registerform"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/goatmail"
	"github.com/goatcms/goatcore/messages"
	"github.com/goatcms/goatcore/messages/msgcollection"
)

// UserRegister is global user register service
type UserRegister struct {
	deps struct {
		Crypt    services.Crypt  `dependency:"CryptService"`
		Mailer   services.Mailer `dependency:"MailerService"`
		Logger   services.Logger `dependency:"LoggerService"`
		Inserter dao.UserInsert  `dependency:"UserInsert"`
	}
}

// RegisterFactory create a user register instance
func RegisterFactory(dp dependency.Provider) (interface{}, error) {
	ins := &UserRegister{}
	if err := dp.InjectTo(&ins.deps); err != nil {
		return nil, err
	}
	return services.UserRegisterAction(ins), nil
}

func (service *UserRegister) Register(form *registerform.RegisterForm, scope app.Scope) (msgs messages.MessageMap, err error) {
	var (
		hash string
	)
	msgs = msgcollection.NewMessageMap()
	if err = form.Valid("", msgs); err != nil {
		return nil, err
	}
	if len(msgs.GetAll()) > 0 {
		service.deps.Logger.DevLog("UserRegister: valid errors for %v", msgs.GetAll())
		return msgs, nil
	}
	if hash, err = service.deps.Crypt.Hash(form.Password); err != nil {
		return nil, err
	}
	user := &entities.User{
		Email:    &form.Email,
		Login:    &form.Email,
		Password: &hash,
	}
	service.deps.Logger.DevLog("UserRegister: %v... insert start", form.Email)
	if _, err = service.deps.Inserter.Insert(scope, user); err != nil {
		return nil, err
	}
	if err = scope.Trigger(app.CommitEvent, nil); err != nil {
		return nil, err
	}
	service.deps.Logger.DevLog("UserRegister: new user (id:%v) %v... start send email", user.ID, form.Email)
	service.deps.Mailer.Send(form.Email, "register", form, []goatmail.Attachment{}, scope)
	return msgs, nil
}
