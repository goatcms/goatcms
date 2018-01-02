package dao

import (
	"database/sql"
	"fmt"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	entities "github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"strconv"
)

// ArticleUpdate is a Data Access Object for article entity
type ArticleUpdate struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewArticleUpdate(di dependency.Injector) (*ArticleUpdate, error) {
	instance := &ArticleUpdate{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func ArticleUpdateFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewArticleUpdate(dp)
	if err != nil {
		return nil, err
	}
	return maindef.ArticleUpdate(instance), nil
}

func (dao ArticleUpdate) Update(scope app.Scope, entity *entities.Article, fields []string) (err error) {
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

func (dao ArticleUpdate) SQL(fields []string, entity *entities.Article) (string, error) {
	sql := "UPDATE Article SET "
	for i, row := range fields {
		if i == 0 {
			sql += row + "="
		} else {
			sql += ", " + row + "="
		}
		switch row {
		case "Content":
			sql += strconv.Quote(*entity.Content)
		case "Title":
			sql += strconv.Quote(*entity.Title)
		}
	}
	return sql, nil
}
