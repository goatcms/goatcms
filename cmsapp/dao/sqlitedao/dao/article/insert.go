package dao

import (
	"database/sql"
	"fmt"
	"strconv"

	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	entities "github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/varutil"
)

// ArticleInsert is a Data Access Object for article entity
type ArticleInsert struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewArticleInsert(di dependency.Injector) (*ArticleInsert, error) {
	instance := &ArticleInsert{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func ArticleInsertFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewArticleInsert(dp)
	if err != nil {
		return nil, err
	}
	return maindef.ArticleInsert(instance), nil
}

func (dao ArticleInsert) Insert(scope app.Scope, entity *entities.Article) (id int64, err error) {
	var (
		sqlq   string
		tx     *sql.Tx
		result sql.Result
	)
	if tx, err = helpers.TX(scope, dao.deps.DB); err != nil {
		return -1, err
	}
	if sqlq, err = dao.SQL(entity); err != nil {
		return -1, err
	}
	if result, err = tx.Exec(sqlq); err != nil {
		return -1, fmt.Errorf("%s: %s", err.Error(), sqlq)
	}
	if id, err = result.LastInsertId(); err != nil {
		return -1, fmt.Errorf("%s: %s", err.Error(), sqlq)
	}
	if err = varutil.SetField(entity, "ID", &id); err != nil {
		return -1, fmt.Errorf("%s: %s", err.Error(), sqlq)
	}
	return id, nil
}

func (dao ArticleInsert) SQL(entity *entities.Article) (string, error) {
	return "INSERT INTO Article (Content, Title) VALUES (" + strconv.Quote(*entity.Content) + ", " + strconv.Quote(*entity.Title) + ")", nil
}
