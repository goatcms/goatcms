package dao

import (
	"database/sql"
	"fmt"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlite/helpers"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/jmoiron/sqlx"
)

// ArticleDelete is a Data Access Object for article entity
type ArticleDelete struct {
	deps struct {
		DB *sqlx.DB `dependency:"sqlitedb0"`
	}
}

func NewArticleDelete(di dependency.Injector) (*ArticleDelete, error) {
	instance := &ArticleDelete{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func ArticleDeleteFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewArticleDelete(dp)
	if err != nil {
		return nil, err
	}
	return maindef.Delete(instance), nil
}

func (dao ArticleDelete) Delete(scope app.Scope, id int64) error {
	var (
		res         sql.Result
		err         error
		count       int64
		idContainer struct {
			ID int64 `db:"id"`
		}
		tx *sqlx.Tx
	)
	idContainer.ID = id
	if tx, err = helpers.TX(scope, dao.deps.DB); err != nil {
		return err
	}
	sql := "DELETE FROM Article WHERE id=:id"
	if res, err = tx.NamedExec(sql, &idContainer); err != nil {
		return fmt.Errorf("%s: %s", err.Error(), sql)
	}
	if count, err = res.RowsAffected(); err != nil {
		return fmt.Errorf("%s: %s", err.Error(), sql)
	}
	if count != 1 {
		return fmt.Errorf("Delete more than one record (%v records deleted)", count)
	}
	return nil
}

func (dao ArticleDelete) SQL(where string) string {
	return "DELETE FROM Article WHERE " + where
}