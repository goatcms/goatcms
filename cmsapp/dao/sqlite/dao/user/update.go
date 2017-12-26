package userdao

import (
	"database/sql"
	"fmt"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlite/helpers"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/jmoiron/sqlx"
)

// UserUpdate is a Data Access Object for user entity
type UserUpdate struct {
	deps struct {
		DB *sqlx.DB `dependency:"sqlitedb"`
	}
}

func NewUserUpdate(di dependency.Injector) (*UserUpdate, error) {
	instance := &UserUpdate{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func UserUpdateFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewUserUpdate(dp)
	if err != nil {
		return nil, err
	}
	return maindef.Update(instance), nil
}

func (dao UserUpdate) Update(scope app.Scope, entity interface{}, fields []string) (err error) {
	var (
		res   sql.Result
		count int64
		sql   string
		tx    *sqlx.Tx
	)
	if tx, err = helpers.TX(scope, dao.deps.DB); err != nil {
		return err
	}
	if sql, err = dao.SQL(fields); err != nil {
		return err
	}
	if res, err = tx.NamedExec(sql, entity); err != nil {
		return fmt.Errorf("%s: %s", err.Error(), sql)
	}
	if count, err = res.RowsAffected(); err != nil {
		return fmt.Errorf("%s: %s", err.Error(), sql)
	}
	if count != 1 {
		return fmt.Errorf("Update modified more then one record (%v records modyfieds): %s", count, sql)
	}
	return nil
}

func (dao UserUpdate) SQL(fields []string) (string, error) {
	sql := "UPDATE User SET "
	for i, row := range fields {
		if i == 0 {
			sql += row + " = :" + row
		} else {
			sql += ", " + row + " = :" + row
		}
	}
	return sql, nil
}
