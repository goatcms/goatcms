package dao

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

// TranslationSearch is a Data Access Object for translation entity
type TranslationSearch struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewTranslationSearch(di dependency.Injector) (*TranslationSearch, error) {
	instance := &TranslationSearch{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func TranslationSearchFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewTranslationSearch(dp)
	if err != nil {
		return nil, err
	}
	return maindef.TranslationSearch(instance), nil
}

func (dao TranslationSearch) Search(scope app.Scope, fields []string, params *maindef.TranslationSearchParams) (maindef.TranslationRows, error) {
	var (
		err  error
		sqlq string
		tx   *sql.Tx
		rows *sql.Rows
	)
	if tx, err = helpers.TX(scope, dao.deps.DB); err != nil {
		return nil, err
	}
	sqlq = dao.SQL(fields, params)
	if rows, err = tx.Query(sqlq); err != nil {
		return nil, fmt.Errorf("%s: %s", err.Error(), sqlq)
	}
	return NewTranslationRows(rows), nil
}

func (query TranslationSearch) SQL(fields []string, params *maindef.TranslationSearchParams) string {
	var criterias []string = []string{}
	sqlq := "SELECT "
	// selected fields
	i := 0
	for _, row := range fields {
		if i > 0 {
			sqlq += ", "
		}
		sqlq += row
		i++
	}
	// fields
	sqlq += " FROM Translation"
	if params.Key != "" {
		criterias = append(criterias, "Key="+strconv.Quote(params.Key))
	}
	if params.Value != "" {
		criterias = append(criterias, "Value="+strconv.Quote(params.Value))
	}
	if len(criterias) > 0 {
		sqlq += " WHERE " + strings.Join(criterias, " AND ")
	}
	return sqlq
}
