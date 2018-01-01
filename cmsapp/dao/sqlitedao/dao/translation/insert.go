package dao

import (
	"database/sql"
	"fmt"

	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/varutil"
	"github.com/jmoiron/sqlx"
)

// TranslationInsert is a Data Access Object for translation entity
type TranslationInsert struct {
	deps struct {
		DB *sqlx.DB `dependency:"db0.engine"`
	}
}

func NewTranslationInsert(di dependency.Injector) (*TranslationInsert, error) {
	instance := &TranslationInsert{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func TranslationInsertFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewTranslationInsert(dp)
	if err != nil {
		return nil, err
	}
	return maindef.Insert(instance), nil
}

func (dao TranslationInsert) Insert(scope app.Scope, entity interface{}, fields []string) (id int64, err error) {
	var (
		sqlq   string
		tx     *sqlx.Tx
		result sql.Result
	)
	if tx, err = helpers.TX(scope, dao.deps.DB); err != nil {
		return -1, err
	}
	if sqlq, err = dao.SQL(fields); err != nil {
		return -1, err
	}
	if result, err = tx.NamedExec(sqlq, entity); err != nil {
		return -1, fmt.Errorf("%s: %s", err.Error(), sqlq)
	}
	if id, err = result.LastInsertId(); err != nil {
		return -1, fmt.Errorf("%s: %s", err.Error(), sqlq)
	}
	if err = varutil.SetField(entity, "ID", id); err != nil {
		return -1, fmt.Errorf("%s: %s", err.Error(), sqlq)
	}
	return id, nil
}

func (dao TranslationInsert) SQL(fields []string) (string, error) {
	sql := "INSERT INTO Translation ("
	sqlValues := "VALUES ("
	for i, row := range fields {
		if i == 0 {
			sql += "" + row
			sqlValues += ":" + row
		} else {
			sql += ", " + row
			sqlValues += ", :" + row
		}
	}
	return sql + ") " + sqlValues + ")", nil
}
