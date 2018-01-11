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

// SessionUpdate is a Data Access Object for session entity
type SessionUpdate struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewSessionUpdate(di dependency.Injector) (*SessionUpdate, error) {
	instance := &SessionUpdate{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func SessionUpdateFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewSessionUpdate(dp)
	if err != nil {
		return nil, err
	}
	return maindef.SessionUpdate(instance), nil
}

func (dao SessionUpdate) Update(scope app.Scope, entity *entities.Session, fields []string) (err error) {
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

func (dao SessionUpdate) SQL(fields []string, entity *entities.Session) (string, error) {
	sql := "UPDATE Session SET "
	for i, row := range fields {
		if i == 0 {
			sql += row + "="
		} else {
			sql += ", " + row + "="
		}
		switch row {
		case "Secret":
			sql += helpers.Quote(entity.Secret)
		case "UserID":
			sql += helpers.FormatInt(entity.UserID, 10)
		}
	}
	return sql, nil
}
