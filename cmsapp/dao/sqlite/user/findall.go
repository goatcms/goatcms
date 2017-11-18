package userdao

import (
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	sqlitebase "github.com/goatcms/goatcms/cmsapp/dao/sqlite"
	"github.com/goatcms/goatcore/app"
)

// UserFindAll is a Data Access Object for user entity
type UserFindAll struct {
	deps struct {
		DB *sql.DB `dependency:"sqlitedb"`
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
		tx  *sql.Tx
	)
	if tx, err = sqlitebase.TX(scope, dao.deps.DB); err != nil {
		return err
	}
	if sql, err = dao.SQL(fields); err != nil {
		return nil, err
	}
	if rows, err = tx.Queryx(sql); err != nil {
		return nil, fmt.Errorf("%s: %s", err.Error(), q.query)
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
