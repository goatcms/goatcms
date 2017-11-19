package translationdao

import (
	"fmt"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	sqlitebase "github.com/goatcms/goatcms/cmsapp/dao/sqlite"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/varutil"
	"github.com/jmoiron/sqlx"
	"math/rand"
)

// TranslationInsert is a Data Access Object for translation entity
type TranslationInsert struct {
	deps struct {
		DB *sqlx.DB `dependency:"sqlitedb"`
	}
}

func NewTranslationInsert(di dependency.Injector) (*TranslationInsert, error) {
	instance := &TranslationInsert{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func TranslationInsertFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewTranslationInsert(dp)
	if err != nil {
		return nil, err
	}
	return maindef.Insert(instance), nil
}

func (dao TranslationInsert) Insert(scope app.Scope, entity interface{}, fields []string) (id int64, err error) {
	var (
		sql string
		tx  *sqlx.Tx
	)
	if tx, err = sqlitebase.TX(scope, dao.deps.DB); err != nil {
		return -1, err
	}
	if sql, err = dao.SQL(fields); err != nil {
		return -1, err
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

func (dao TranslationInsert) SQL(fields []string) (string, error) {
	sql := "INSERT INTO Translation ("
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
