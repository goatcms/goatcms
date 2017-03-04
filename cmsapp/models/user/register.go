package user

import (
	"github.com/goatcms/goatcms/cmsapp/models"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/db"
	"github.com/goatcms/goatcore/dependency"
)

// UserRegister is register user query
type UserRegister struct {
	deps struct {
		Insert db.Insert      `dependency:"UserInsert"`
		Crypt  services.Crypt `dependency:"CryptService"`
	}
}

// RegisterFactory create new register query object
func RegisterFactory(dp dependency.Provider) (interface{}, error) {
	register := &UserRegister{}
	if err := dp.InjectTo(&register.deps); err != nil {
		return nil, err
	}
	return models.UserRegister(register.Register), nil
}

// Register create new user record
func (q UserRegister) Register(tx db.TX, user *models.User, password string) (int64, error) {
	var (
		hashPassword string
		err          error
	)
	if hashPassword, err = q.deps.Crypt.Hash(password); err != nil {
		return 0, err
	}
	user.PasswordHash = hashPassword
	return q.deps.Insert(tx, user)
}
