package dao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app"
)

type UserSigninQueryParams struct {
	Email    string
	Username string
}

type UserSigninQuery interface {
	Signin(scope app.Scope, fields []string, params *UserSigninQueryParams) (*entities.User, error)
}
