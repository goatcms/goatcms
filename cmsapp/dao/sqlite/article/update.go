package articledao

import (
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	sqlitebase "github.com/goatcms/goatcms/cmsapp/dao/sqlite"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/varutil"
)

// ArticleUpdate is a Data Access Object for article entity
type ArticleUpdate struct {
	deps struct {
		DB *sql.DB `dependency:"sqlitedb"`
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
	return maindef.Update(instance), nil
}

func (dao ArticleUpdate) Update(scope app.Scope, entity interface{}, fields []string) (err error) {
	var (
		res   sql.Result
		count int64
		sql   string
		tx    *sql.Tx
	)
	if tx, err = sqlitebase.TX(scope, dao.deps.DB); err != nil {
		return err
	}
	if sql, err = dao.SQL(fields); err != nil {
		return nil, err
	}
	if res, err = tx.NamedExec(sql, entity); err != nil {
		return fmt.Errorf("%s: %s", err.Error(), sql)
	}
	if count, err = res.RowsAffected(); err != nil {
		return fmt.Errorf("%s: %s", err.Error(), sql)
	}
	if count != 1 {
		return fmt.Errorf("Update modified more then one record (%v records modyfieds): %s", count, sql)
	}
	return nil
}

func (dao ArticleUpdate) SQL(fields []string) (string, error) {
	sql := "UPDATE Article SET "
	for i, row := range fields {
		if i == 0 {
			sql += row + " = :" + row
		} else {
			sql += ", " + row + " = :" + row
		}
	}
	return sql, nil
}
