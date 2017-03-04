package userform

import (
	"github.com/goatcms/goatcms/cmsapp/models"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/messages"
	"github.com/goatcms/goatcore/varutil/validator"
)

const (
	passwordMinLength = 8
	EmailKey          = "Email"
	PasswordHashKey   = "PasswordHash"
)

// UserForm is structure with register form values
type UserForm models.User

func NewForm(dp dependency.Injector) (*UserForm, error) {
	form := &UserForm{}
	if err := dp.InjectTo(form); err != nil {
		return nil, err
	}
	return form, nil
}

func (f *UserForm) Valid(basekey string, mm messages.MessageMap) error {
	if err := validator.EmailValid(f.Email, basekey+EmailKey, mm); err != nil {
		return err
	}
	return nil
}
