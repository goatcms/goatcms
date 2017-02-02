package models

import "github.com/goatcms/goat-core/db"

const (
	UserLogin    = "UserLoginQuery"
	UserRegister = "UserRegisterQuery"
)

type UserLoginQuery func(tx db.TX, name, password string) (db.Row, error)
type UserRegisterQuery func(tx db.TX, user *User) (int64, error)
