package translationdao

import (
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	sqlitebase "github.com/goatcms/goatcms/cmsapp/dao/sqlite"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/jmoiron/sqlx"
	"strconv"
)

// TranslationFindByID is a Data Access Object for translation entity
type TranslationFindByID struct {
	deps struct {
		DB *sqlx.DB `dependency:"sqlitedb"`
	}
}

func NewTranslationFindByID(di dependency.Injector) (*TranslationFindByID, error) {
	instance := &TranslationFindByID{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func TranslationFindByIDFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewTranslationFindByID(dp)
	if err != nil {
		return nil, err
	}
	return maindef.FindByID(instance), nil
}

func (dao TranslationFindByID) Find(scope app.Scope, fields []string, id int64) (row maindef.Row, err error) {
	var (
		sql string
		tx  *sqlx.Tx
	)
	if tx, err = sqlitebase.TX(scope, dao.deps.DB); err != nil {
		return nil, err
	}
	if sql, err = dao.SQL(fields, id); err != nil {
		return nil, err
	}
	row = tx.QueryRowx(sql)
	return row, nil
}

func (dao TranslationFindByID) SQL(fields []string, id int64) (string, error) {
	sql := "SELECT "
	i := 0
	for _, row := range fields {
		if i > 0 {
			sql += ", "
		}
		sql += row
		i++
	}
	return sql + " FROM Translation WHERE id=" + strconv.FormatInt(id, 10) + " LIMIT 1", nil
}
