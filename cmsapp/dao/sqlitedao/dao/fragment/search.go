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

// FragmentSearch is a Data Access Object for fragment entity
type FragmentSearch struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewFragmentSearch(di dependency.Injector) (*FragmentSearch, error) {
	instance := &FragmentSearch{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func FragmentSearchFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewFragmentSearch(dp)
	if err != nil {
		return nil, err
	}
	return maindef.FragmentSearch(instance), nil
}

func (dao FragmentSearch) Search(scope app.Scope, fields []string, params *maindef.FragmentSearchParams) (maindef.FragmentRows, error) {
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
	return NewFragmentRows(rows), nil
}

func (query FragmentSearch) SQL(fields []string, params *maindef.FragmentSearchParams) string {
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
	sqlq += " FROM Fragment"
	if params.Lang != "" {
		criterias = append(criterias, "Lang="+strconv.Quote(params.Lang))
	}
	if params.Name != "" {
		criterias = append(criterias, "Name="+strconv.Quote(params.Name))
	}
	if params.Content != "" {
		criterias = append(criterias, "Content="+strconv.Quote(params.Content))
	}
	if len(criterias) > 0 {
		sqlq += " WHERE " + strings.Join(criterias, " AND ")
	}
	return sqlq
}
