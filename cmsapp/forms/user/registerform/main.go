package registerform

import (
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goat-core/messages"
	"github.com/goatcms/goat-core/varutil/validator"
	"github.com/goatcms/goatcms/cmsapp/forms/user/userform"
)

const (
	passwordMinLength = 8
	IDKey             = "id"
	PasswordKey       = "password"
)

// RegisterForm is structure with register form values
type RegisterForm struct {
	User           userform.UserForm `form:"User."`
	Password       string            `form:"Password"`
	RepeatPassword string            `form:"RepeatPassword"`
}

func NewForm(dp dependency.Injector) (*RegisterForm, error) {
	var err error
	form := &RegisterForm{}
	if err = dp.InjectTo(form); err != nil {
		return nil, err
	}
	return form, nil
}

func (f *RegisterForm) Valid(basekey string, mm messages.MessageMap) error {
	if err := f.User.Valid("User.", mm); err != nil {
		return err
	}
	if err := validator.MinStringValid(f.Password, PasswordKey, mm, passwordMinLength); err != nil {
		return err
	}
	return nil
}
