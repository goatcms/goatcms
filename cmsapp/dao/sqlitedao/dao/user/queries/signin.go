package queries

import (
	"database/sql"
	"fmt"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	dao "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/dao/user"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	entities "github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"strconv"
)

// UserSigninQuery is a query object for
type UserSigninQuery struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewUserSigninQuery(di dependency.Injector) (*UserSigninQuery, error) {
	instance := &UserSigninQuery{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func UserSigninQueryFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewUserSigninQuery(dp)
	if err != nil {
		return nil, err
	}
	return maindef.UserSigninQuery(instance), nil
}

func (query UserSigninQuery) Signin(scope app.Scope, fields []string, params *maindef.UserSigninQueryParams) (*entities.User, error) {
	var (
		err    error
		sqlq   string
		tx     *sql.Tx
		row    maindef.UserRow
		entity *entities.User
	)
	if tx, err = helpers.TX(scope, query.deps.DB); err != nil {
		return nil, err
	}
	sqlq = query.SQL(fields, params)
	row = dao.NewUserRow(tx.QueryRow(sqlq), fields)
	if entity, err = row.Get(); err != nil {
		return nil, fmt.Errorf("%v: %v", err, sqlq)
	}
	return entity, nil
}

func (query UserSigninQuery) SQL(fields []string, params *maindef.UserSigninQueryParams) string {
	sqlq := "SELECT "
	// selected fields
	i := 0
	for _, row := range fields {
		if i > 0 {
			sqlq += ", "
		}
		sqlq += row
		i++
	}
	// fields
	sqlq += " FROM User WHERE "
	if params.Username != "" {
		sqlq += "(Username=" + strconv.Quote(params.Username)
	}
	if params.Email != "" {
		sqlq += ") OR (Email=" + strconv.Quote(params.Email)
	}
	sqlq += ") LIMIT 1"
	return sqlq
}
