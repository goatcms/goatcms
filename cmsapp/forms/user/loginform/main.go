package loginform

import (
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goat-core/messages"
	"github.com/goatcms/goat-core/messages/msgcollection"
	"github.com/goatcms/goat-core/varutil/validator"
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

func NewForm(dp dependency.Provider) (*LoginForm, error) {
	form := &LoginForm{}
	if err := dp.InjectTo(form); err != nil {
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
