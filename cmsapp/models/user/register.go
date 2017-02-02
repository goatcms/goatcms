package user

import (
	"github.com/goatcms/goat-core/db"
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/cmsapp/models"
)

// UserRegisterQuery is register user query
type UserRegisterQuery struct {
	deps struct {
		InsertQuery db.Insert `dependency:"db.query.user.Insert"`
	}
}

// RegisterQueryFactory create new register query object
func RegisterQueryFactory(dp dependency.Provider) (interface{}, error) {
	query := &UserRegisterQuery{}
	if err := dp.InjectTo(query); err != nil {
		return nil, err
	}
	return models.UserRegisterQuery(query.Register), nil
}

// Insert create new record
func (q UserRegisterQuery) Register(tx db.TX, user *models.User) (int64, error) {
	return q.deps.InsertQuery(tx, user)
}
