package dao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app"
)

type UserLoginQueryParams struct {
	Email    string
	Password string
	Login    string
}

type UserLoginQuery interface {
	Login(scope app.Scope, fields []string, params *UserLoginQueryParams) (*entities.User, error)
}
