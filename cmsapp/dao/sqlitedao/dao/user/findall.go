package dao

import (
	"fmt"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/jmoiron/sqlx"
)

// UserFindAll is a Data Access Object for user entity
type UserFindAll struct {
	deps struct {
		DB *sqlx.DB `dependency:"db0.engine"`
	}
}

func NewUserFindAll(di dependency.Injector) (*UserFindAll, error) {
	instance := &UserFindAll{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func UserFindAllFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewUserFindAll(dp)
	if err != nil {
		return nil, err
	}
	return maindef.FindAll(instance), nil
}

func (dao UserFindAll) Find(scope app.Scope, fields []string) (rows maindef.Rows, err error) {
	var (
		sql string
		tx  *sqlx.Tx
	)
	if tx, err = helpers.TX(scope, dao.deps.DB); err != nil {
		return nil, err
	}
	if sql, err = dao.SQL(fields); err != nil {
		return nil, err
	}
	if rows, err = tx.Queryx(sql); err != nil {
		return nil, fmt.Errorf("%s: %s", err.Error(), sql)
	}
	return rows.(maindef.Rows), nil
}

func (dao UserFindAll) SQL(fields []string) (string, error) {
	sql := "SELECT "
	i := 0
	for _, row := range fields {
		if i > 0 {
			sql += ", "
		}
		sql += row
		i++
	}
	return sql + " FROM User", nil
}
