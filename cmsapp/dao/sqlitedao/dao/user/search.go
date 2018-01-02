package dao

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

// UserSearch is a Data Access Object for user entity
type UserSearch struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewUserSearch(di dependency.Injector) (*UserSearch, error) {
	instance := &UserSearch{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func UserSearchFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewUserSearch(dp)
	if err != nil {
		return nil, err
	}
	return maindef.UserSearch(instance), nil
}

func (dao UserSearch) Search(scope app.Scope, fields []string, params *maindef.UserSearchParams) (maindef.UserRows, error) {
	var (
		err  error
		sqlq string
		tx   *sql.Tx
		rows *sql.Rows
	)
	if tx, err = helpers.TX(scope, dao.deps.DB); err != nil {
		return nil, err
	}
	sqlq = dao.SQL(fields, params)
	if rows, err = tx.Query(sqlq); err != nil {
		return nil, fmt.Errorf("%s: %s", err.Error(), sqlq)
	}
	return NewUserRows(rows), nil
}

func (query UserSearch) SQL(fields []string, params *maindef.UserSearchParams) string {
	var criterias []string = []string{}
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
	sqlq += " FROM User"
	if params.Password != "" {
		criterias = append(criterias, "Password="+strconv.Quote(params.Password))
	}
	if params.Firstname != "" {
		criterias = append(criterias, "Firstname="+strconv.Quote(params.Firstname))
	}
	if params.Email != "" {
		criterias = append(criterias, "Email="+strconv.Quote(params.Email))
	}
	if params.Login != "" {
		criterias = append(criterias, "Login="+strconv.Quote(params.Login))
	}
	if len(criterias) > 0 {
		sqlq += " WHERE " + strings.Join(criterias, " AND ")
	}
	return sqlq
}
