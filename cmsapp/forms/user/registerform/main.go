package registerform

import (
	"github.com/goatcms/goatcms/cmsapp/forms/user/userform"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/messages"
	"github.com/goatcms/goatcore/varutil/validator"
)

// RegisterForm is structure with register form values
type RegisterForm struct {
	User           userform.UserForm `form:"User."`
	Password       string            `form:"Password"`
	RepeatPassword string            `form:"RepeatPassword"`
}

func NewForm(dp dependency.Injector) (*RegisterForm, error) {
	var err error
	if err != nil {
		return nil, err
	}
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
	if err := validator.MinStringValid(f.Password, "Password", mm, 8); err != nil {
		return err
	}
	return nil
}
