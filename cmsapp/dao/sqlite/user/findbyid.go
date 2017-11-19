package userdao

import (
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	sqlitebase "github.com/goatcms/goatcms/cmsapp/dao/sqlite"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/jmoiron/sqlx"
	"strconv"
)

// UserFindByID is a Data Access Object for user entity
type UserFindByID struct {
	deps struct {
		DB *sqlx.DB `dependency:"sqlitedb"`
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
	return maindef.FindByID(instance), nil
}

func (dao UserFindByID) Find(scope app.Scope, fields []string, id int64) (row maindef.Row, err error) {
	var (
		sql string
		tx  *sqlx.Tx
	)
	if tx, err = sqlitebase.TX(scope, dao.deps.DB); err != nil {
		return nil, err
	}
	if sql, err = dao.SQL(fields, id); err != nil {
		return nil, err
	}
	row = tx.QueryRowx(sql)
	return row, nil
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
