package dao

import (
	"database/sql"
	"fmt"
	"strconv"

	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	entities "github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

// UserInsert is a Data Access Object for user entity
type UserInsert struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewUserInsert(di dependency.Injector) (*UserInsert, error) {
	instance := &UserInsert{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func UserInsertFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewUserInsert(dp)
	if err != nil {
		return nil, err
	}
	return maindef.UserInsert(instance), nil
}

func (dao UserInsert) Insert(scope app.Scope, entity *entities.User) (id int64, err error) {
	var (
		sqlq   string
		tx     *sql.Tx
		result sql.Result
	)
	if tx, err = helpers.TX(scope, dao.deps.DB); err != nil {
		return -1, err
	}
	if sqlq, err = dao.SQL(entity); err != nil {
		return -1, err
	}
	if result, err = tx.Exec(sqlq); err != nil {
		return -1, fmt.Errorf("%s: %s", err.Error(), sqlq)
	}
	if id, err = result.LastInsertId(); err != nil {
		return -1, fmt.Errorf("%s: %s", err.Error(), sqlq)
	}
	entity.ID = &id
	return id, nil
}

func (dao UserInsert) SQL(entity *entities.User) (string, error) {
	sql := "INSERT INTO User ("
	if entity.ID != nil {
		sql += "ID, "
	}
	sql += "Firstname, Lastname, Email, Password, Roles, Username, ID, ID, ID, ID, ID, ID) VALUES ("
	if entity.ID != nil {
		sql += strconv.FormatInt(*entity.ID, 10) + ", "
	}
	sql += "" + helpers.Quote(entity.Firstname) + ", " + helpers.Quote(entity.Lastname) + ", " + helpers.Quote(entity.Email) + ", " + helpers.Quote(entity.Password) + ", " + helpers.Quote(entity.Roles) + ", " + helpers.Quote(entity.Username) + ", " + helpers.FormatInt(entity.ID, 10) + ", " + helpers.FormatInt(entity.ID, 10) + ", " + helpers.FormatInt(entity.ID, 10) + ", " + helpers.FormatInt(entity.ID, 10) + ", " + helpers.FormatInt(entity.ID, 10) + ", " + helpers.FormatInt(entity.ID, 10) + ")"
	return sql, nil
}
