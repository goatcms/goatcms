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

// SessionSearch is a Data Access Object for session entity
type SessionSearch struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewSessionSearch(di dependency.Injector) (*SessionSearch, error) {
	instance := &SessionSearch{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func SessionSearchFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewSessionSearch(dp)
	if err != nil {
		return nil, err
	}
	return maindef.SessionSearch(instance), nil
}

func (dao SessionSearch) Search(scope app.Scope, fields []string, params *maindef.SessionSearchParams) (maindef.SessionRows, error) {
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
	return NewSessionRows(rows), nil
}

func (query SessionSearch) SQL(fields []string, params *maindef.SessionSearchParams) string {
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
	sqlq += " FROM Session"
	if params.Secret != "" {
		criterias = append(criterias, "Secret="+strconv.Quote(params.Secret))
	}
	if len(criterias) > 0 {
		sqlq += " WHERE " + strings.Join(criterias, " AND ")
	}
	return sqlq
}
