package dao

import (
	"database/sql"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"strconv"
)

// ArticleFindByID is a Data Access Object for article entity
type ArticleFindByID struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewArticleFindByID(di dependency.Injector) (*ArticleFindByID, error) {
	instance := &ArticleFindByID{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func ArticleFindByIDFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewArticleFindByID(dp)
	if err != nil {
		return nil, err
	}
	return maindef.ArticleFindByID(instance), nil
}

func (dao ArticleFindByID) Find(scope app.Scope, fields []string, id int64) (maindef.ArticleRow, error) {
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
	return NewArticleRow(row, fields), nil
}

func (dao ArticleFindByID) SQL(fields []string, id int64) (string, error) {
	sql := "SELECT "
	i := 0
	for _, row := range fields {
		if i > 0 {
			sql += ", "
		}
		sql += row
		i++
	}
	return sql + " FROM Article WHERE id=" + strconv.FormatInt(id, 10) + " LIMIT 1", nil
}
