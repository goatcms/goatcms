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

// SessionInsert is a Data Access Object for session entity
type SessionInsert struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewSessionInsert(di dependency.Injector) (*SessionInsert, error) {
	instance := &SessionInsert{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func SessionInsertFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewSessionInsert(dp)
	if err != nil {
		return nil, err
	}
	return maindef.SessionInsert(instance), nil
}

func (dao SessionInsert) Insert(scope app.Scope, entity *entities.Session) (id int64, err error) {
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

func (dao SessionInsert) SQL(entity *entities.Session) (string, error) {
	sql := "INSERT INTO Session ("
	if entity.ID != nil {
		sql += "ID, "
	}
	sql += "Secret, UserID) VALUES ("
	if entity.ID != nil {
		sql += strconv.FormatInt(*entity.ID, 10) + ", "
	}
	sql += "" + helpers.Quote(entity.Secret) + ", " + helpers.FormatInt(entity.UserID, 10) + ")"
	return sql, nil
}
