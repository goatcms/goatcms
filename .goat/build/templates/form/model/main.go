package loginform

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/messages"
	"github.com/goatcms/goatcore/messages/msgcollection"
	"github.com/goatcms/goatcore/varutil/validator"
)

const (
	passwordMinLength = 8
	UsernameKey       = "Username"
	PasswordKey       = "Password"
)

// LoginForm is structure with register form values
type LoginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func NewForm(scope app.Scope) (*LoginForm, error) {
	form := &LoginForm{}
	if err := scope.InjectTo(form); err != nil {
		return nil, err
	}
	return form, nil
}

func (f *LoginForm) Valid() (messages.MessageMap, error) {
	mm := msgcollection.NewMessageMap()
	if err := validator.EmailValid(f.Username, UsernameKey, mm); err != nil {
		return nil, err
	}
	if err := validator.MinStringValid(f.Password, PasswordKey, mm, 1); err != nil {
		return nil, err
	}
	if err := validator.MinStringValid(f.Username, UsernameKey, mm, 1); err != nil {
		return nil, err
	}
	return mm, nil
}
