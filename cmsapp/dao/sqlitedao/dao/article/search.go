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

// ArticleSearch is a Data Access Object for article entity
type ArticleSearch struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewArticleSearch(di dependency.Injector) (*ArticleSearch, error) {
	instance := &ArticleSearch{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func ArticleSearchFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewArticleSearch(dp)
	if err != nil {
		return nil, err
	}
	return maindef.ArticleSearch(instance), nil
}

func (dao ArticleSearch) Search(scope app.Scope, fields []string, params *maindef.ArticleSearchParams) (maindef.ArticleRows, error) {
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
	return NewArticleRows(rows), nil
}

func (query ArticleSearch) SQL(fields []string, params *maindef.ArticleSearchParams) string {
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
	sqlq += " FROM Article"
	if params.Content != "" {
		criterias = append(criterias, "Content="+strconv.Quote(params.Content))
	}
	if params.Title != "" {
		criterias = append(criterias, "Title="+strconv.Quote(params.Title))
	}
	if len(criterias) > 0 {
		sqlq += " WHERE " + strings.Join(criterias, " AND ")
	}
	return sqlq
}
