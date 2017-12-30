package queries

import (
	"fmt"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/jmoiron/sqlx"
	"strconv"
)

// UserLoginQuery is a query object for
type UserLoginQuery struct {
	deps struct {
		DB *sqlx.DB `dependency:"sqlitedb0"`
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

func (dao UserLoginQuery) Login(scope app.Scope, fields []string, params *maindef.UserLoginQueryParams) (row maindef.Row, err error) {
	var (
		sql string
		tx  *sqlx.Tx
	)
	if tx, err = helpers.TX(scope, dao.deps.DB); err != nil {
		return nil, err
	}
	sql = dao.SQL(fields, params)
	row = tx.QueryRowx(sql)
	if row.Err() != nil {
		return nil, fmt.Errorf("%v: %v", row.Err(), sql)
	}
	return row, nil
}

func (dao UserLoginQuery) SQL(fields []string, params *maindef.UserLoginQueryParams) string {
	sql := "SELECT "
	// selected fields
	i := 0
	for _, row := range fields {
		if i > 0 {
			sql += ", "
		}
		sql += row
		i++
	}
	// fields
	sql += " FROM User WHERE "
	if params.Email != "" || params.Password != "" {
		sql += "(Email=" + strconv.Quote(params.Email) + " AND Password=" + strconv.Quote(params.Password)
	}
	if params.Login != "" || params.Password != "" {
		sql += ") OR (Login=" + strconv.Quote(params.Login) + " AND Password=" + strconv.Quote(params.Password)
	}
	sql += ") LIMIT 1"
	return sql
}
