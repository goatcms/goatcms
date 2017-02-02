package userform

import (
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goat-core/messages"
	"github.com/goatcms/goat-core/varutil/validator"
	"github.com/goatcms/goatcms/cmsapp/models"
)

const (
	passwordMinLength = 8
	EmailKey          = "email"
	PasswordHashKey   = "passwordHash"
)

// UserForm is structure with register form values
type UserForm models.User

func NewForm(dp dependency.Provider) (*UserForm, error) {
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
	if err := validator.MinStringValid(f.PasswordHash, basekey+PasswordHashKey, mm, 1); err != nil {
		return err
	}
	return nil
}
