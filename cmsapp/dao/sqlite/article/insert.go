package articledao

import (
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	sqlitebase "github.com/goatcms/goatcms/cmsapp/dao/sqlite"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/varutil"
)

// ArticleInsert is a Data Access Object for article entity
type ArticleInsert struct {
	deps struct {
		DB *sql.DB `dependency:"sqlitedb"`
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
	return maindef.Insert(instance), nil
}

func (dao ArticleInsert) Insert(scope app.Scope, entity interface{}, fields []string) (id int64, err error) {
	var (
		sql string
		tx  *sql.Tx
	)
	if tx, err = sqlitebase.TX(scope, dao.deps.DB); err != nil {
		return err
	}
	if sql, err = dao.SQL(fields); err != nil {
		return nil, err
	}
	id = rand.Int63()
	if err = varutil.SetField(entity, "ID", id); err != nil {
		return -1, fmt.Errorf("%s: %s", err.Error(), sql)
	}
	if _, err = tx.NamedExec(sql, entity); err != nil {
		return -1, fmt.Errorf("%s: %s", err.Error(), sql)
	}
	return id, nil
}

func (dao ArticleInsert) SQL(fields []string) (string, error) {
	sql := "INSERT INTO Article ("
	sqlValues := "VALUES ("
	for i, row := range fields {
		if i == 0 {
			sql += "" + row
			sqlValues += ":" + row
		} else {
			sql += ", " + row
			sqlValues += ", :" + row
		}
	}
	return sql + ") " + sqlValues + ")", nil
}
