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

// UserLoginQuery is a query object for
type UserLoginQuery struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewUserLoginQuery(di dependency.Injector) (*UserLoginQuery, error) {
	instance := &UserLoginQuery{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func UserLoginQueryFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewUserLoginQuery(dp)
	if err != nil {
		return nil, err
	}
	return maindef.UserLoginQuery(instance), nil
}

func (query UserLoginQuery) Login(scope app.Scope, fields []string, params *maindef.UserLoginQueryParams) (*entities.User, error) {
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

func (query UserLoginQuery) SQL(fields []string, params *maindef.UserLoginQueryParams) string {
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
	if params.Login != "" || params.Password != "" {
		sqlq += "(Login=" + strconv.Quote(params.Login) + " AND Password=" + strconv.Quote(params.Password)
	}
	if params.Email != "" || params.Password != "" {
		sqlq += ") OR (Email=" + strconv.Quote(params.Email) + " AND Password=" + strconv.Quote(params.Password)
	}
	sqlq += ") LIMIT 1"
	return sqlq
}
