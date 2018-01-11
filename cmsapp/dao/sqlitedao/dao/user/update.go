package dao

import (
	"database/sql"
	"fmt"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	entities "github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

// UserUpdate is a Data Access Object for user entity
type UserUpdate struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
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
	return maindef.UserUpdate(instance), nil
}

func (dao UserUpdate) Update(scope app.Scope, entity *entities.User, fields []string) (err error) {
	var (
		res   sql.Result
		count int64
		query string
		tx    *sql.Tx
	)
	if tx, err = helpers.TX(scope, dao.deps.DB); err != nil {
		return err
	}
	if query, err = dao.SQL(fields, entity); err != nil {
		return fmt.Errorf("%s: %s", err.Error(), query)
	}
	if res, err = tx.Exec(query); err != nil {
	}
	if count, err = res.RowsAffected(); err != nil {
		return fmt.Errorf("%s: %s", err.Error(), query)
	}
	if count != 1 {
		return fmt.Errorf("Update modified more then one record (%v records modyfieds): %s", count, query)
	}
	return nil
}

func (dao UserUpdate) SQL(fields []string, entity *entities.User) (string, error) {
	sql := "UPDATE User SET "
	for i, row := range fields {
		if i == 0 {
			sql += row + "="
		} else {
			sql += ", " + row + "="
		}
		switch row {
		case "Firstname":
			sql += helpers.Quote(entity.Firstname)
		case "Lastname":
			sql += helpers.Quote(entity.Lastname)
		case "Email":
			sql += helpers.Quote(entity.Email)
		case "Password":
			sql += helpers.Quote(entity.Password)
		case "Roles":
			sql += helpers.Quote(entity.Roles)
		case "Username":
			sql += helpers.Quote(entity.Username)
		}
	}
	return sql, nil
}
