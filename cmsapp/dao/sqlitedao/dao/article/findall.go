package dao

import (
	"database/sql"
	"fmt"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

// ArticleFindAll is a Data Access Object for article entity
type ArticleFindAll struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewArticleFindAll(di dependency.Injector) (*ArticleFindAll, error) {
	instance := &ArticleFindAll{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func ArticleFindAllFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewArticleFindAll(dp)
	if err != nil {
		return nil, err
	}
	return maindef.ArticleFindAll(instance), nil
}

func (dao ArticleFindAll) Find(scope app.Scope, fields []string) (maindef.ArticleRows, error) {
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
	return NewArticleRows(rows), nil
}

func (dao ArticleFindAll) SQL(fields []string) (string, error) {
	sql := "SELECT "
	i := 0
	for _, row := range fields {
		if i > 0 {
			sql += ", "
		}
		sql += row
		i++
	}
	return sql + " FROM Article", nil
}
