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

// UserFindByID is a Data Access Object for user entity
type UserFindByID struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewUserFindByID(di dependency.Injector) (*UserFindByID, error) {
	instance := &UserFindByID{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func UserFindByIDFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewUserFindByID(dp)
	if err != nil {
		return nil, err
	}
	return maindef.UserFindByID(instance), nil
}

func (dao UserFindByID) Find(scope app.Scope, fields []string, id int64) (*entities.User, error) {
	var (
		err   error
		query string
		tx    *sql.Tx
		row   *UserRow
	)
	if tx, err = helpers.TX(scope, dao.deps.DB); err != nil {
		return nil, err
	}
	if query, err = dao.SQL(fields, id); err != nil {
		return nil, err
	}
	row = NewUserRow(tx.QueryRow(query), fields)
	return row.Get()
}

func (dao UserFindByID) SQL(fields []string, id int64) (string, error) {
	sql := "SELECT "
	i := 0
	for _, row := range fields {
		if i > 0 {
			sql += ", "
		}
		sql += row
		i++
	}
	return sql + " FROM User WHERE id=" + strconv.FormatInt(id, 10) + " LIMIT 1", nil
}
