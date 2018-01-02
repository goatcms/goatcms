package dao

import (
	"database/sql"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"strconv"
)

// TranslationFindByID is a Data Access Object for translation entity
type TranslationFindByID struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewTranslationFindByID(di dependency.Injector) (*TranslationFindByID, error) {
	instance := &TranslationFindByID{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func TranslationFindByIDFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewTranslationFindByID(dp)
	if err != nil {
		return nil, err
	}
	return maindef.TranslationFindByID(instance), nil
}

func (dao TranslationFindByID) Find(scope app.Scope, fields []string, id int64) (maindef.TranslationRow, error) {
	var (
		err   error
		query string
		tx    *sql.Tx
		row   *sql.Row
	)
	if tx, err = helpers.TX(scope, dao.deps.DB); err != nil {
		return nil, err
	}
	if query, err = dao.SQL(fields, id); err != nil {
		return nil, err
	}
	row = tx.QueryRow(query)
	return NewTranslationRow(row, fields), nil
}

func (dao TranslationFindByID) SQL(fields []string, id int64) (string, error) {
	sql := "SELECT "
	i := 0
	for _, row := range fields {
		if i > 0 {
			sql += ", "
		}
		sql += row
		i++
	}
	return sql + " FROM Translation WHERE id=" + strconv.FormatInt(id, 10) + " LIMIT 1", nil
}
