package dao

import (
	"database/sql"
	"fmt"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

// SessionFindAll is a Data Access Object for session entity
type SessionFindAll struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewSessionFindAll(di dependency.Injector) (*SessionFindAll, error) {
	instance := &SessionFindAll{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func SessionFindAllFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewSessionFindAll(dp)
	if err != nil {
		return nil, err
	}
	return maindef.SessionFindAll(instance), nil
}

func (dao SessionFindAll) Find(scope app.Scope, fields []string) (maindef.SessionRows, error) {
	var (
		err   error
		query string
		tx    *sql.Tx
		rows  *sql.Rows
	)
	if tx, err = helpers.TX(scope, dao.deps.DB); err != nil {
		return nil, err
	}
	if query, err = dao.SQL(fields); err != nil {
		return nil, err
	}
	if rows, err = tx.Query(query); err != nil {
		return nil, fmt.Errorf("%s: %s", err.Error(), query)
	}
	return NewSessionRows(rows), nil
}

func (dao SessionFindAll) SQL(fields []string) (string, error) {
	sql := "SELECT "
	i := 0
	for _, row := range fields {
		if i > 0 {
			sql += ", "
		}
		sql += row
		i++
	}
	return sql + " FROM Session", nil
}
