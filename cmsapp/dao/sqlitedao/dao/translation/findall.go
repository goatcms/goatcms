package dao

import (
	"database/sql"
	"fmt"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

// TranslationFindAll is a Data Access Object for translation entity
type TranslationFindAll struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewTranslationFindAll(di dependency.Injector) (*TranslationFindAll, error) {
	instance := &TranslationFindAll{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func TranslationFindAllFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewTranslationFindAll(dp)
	if err != nil {
		return nil, err
	}
	return maindef.TranslationFindAll(instance), nil
}

func (dao TranslationFindAll) Find(scope app.Scope, fields []string) (maindef.TranslationRows, error) {
	var (
		err   error
		query string
		tx    *sql.Tx
		rows  *sql.Rows
	)
	if tx, err = helpers.TX(scope, dao.deps.DB); err != nil {
		return nil, err
	}
	if query, err = dao.SQL(fields); err != nil {
		return nil, err
	}
	if rows, err = tx.Query(query); err != nil {
		return nil, fmt.Errorf("%s: %s", err.Error(), query)
	}
	return NewTranslationRows(rows), nil
}

func (dao TranslationFindAll) SQL(fields []string) (string, error) {
	sql := "SELECT "
	i := 0
	for _, row := range fields {
		if i > 0 {
			sql += ", "
		}
		sql += row
		i++
	}
	return sql + " FROM Translation", nil
}
