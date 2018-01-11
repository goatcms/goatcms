package dao

import (
	"database/sql"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	entities "github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"strconv"
)

// SessionFindByID is a Data Access Object for session entity
type SessionFindByID struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewSessionFindByID(di dependency.Injector) (*SessionFindByID, error) {
	instance := &SessionFindByID{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func SessionFindByIDFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewSessionFindByID(dp)
	if err != nil {
		return nil, err
	}
	return maindef.SessionFindByID(instance), nil
}

func (dao SessionFindByID) Find(scope app.Scope, fields []string, id int64) (*entities.Session, error) {
	var (
		err   error
		query string
		tx    *sql.Tx
		row   *SessionRow
	)
	if tx, err = helpers.TX(scope, dao.deps.DB); err != nil {
		return nil, err
	}
	if query, err = dao.SQL(fields, id); err != nil {
		return nil, err
	}
	row = NewSessionRow(tx.QueryRow(query), fields)
	return row.Get()
}

func (dao SessionFindByID) SQL(fields []string, id int64) (string, error) {
	sql := "SELECT "
	i := 0
	for _, row := range fields {
		if i > 0 {
			sql += ", "
		}
		sql += row
		i++
	}
	return sql + " FROM Session WHERE id=" + strconv.FormatInt(id, 10) + " LIMIT 1", nil
}
