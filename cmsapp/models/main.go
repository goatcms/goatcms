package models

import "github.com/goatcms/goatcore/db"

const (
	UserLoginService    = "UserLogin"
	UserRegisterService = "UserRegister"
)

type UserLogin func(tx db.TX, name, password string) (*User, error)
type UserRegister func(tx db.TX, user *User, password string) (int64, error)
