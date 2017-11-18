package translationdao

import (
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	sqlitebase "github.com/goatcms/goatcms/cmsapp/dao/sqlite"
	"github.com/goatcms/goatcore/app"
)

// TranslationFindAll is a Data Access Object for translation entity
type TranslationFindAll struct {
	deps struct {
		DB *sql.DB `dependency:"sqlitedb"`
	}
}

func NewTranslationFindAll(di dependency.Injector) (*TranslationFindAll, error) {
	instance := &TranslationFindAll{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func TranslationFindAllFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewTranslationFindAll(dp)
	if err != nil {
		return nil, err
	}
	return maindef.FindAll(instance), nil
}

func (dao TranslationFindAll) Find(scope app.Scope, fields []string) (rows maindef.Rows, err error) {
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
	if rows, err = tx.Queryx(sql); err != nil {
		return nil, fmt.Errorf("%s: %s", err.Error(), q.query)
	}
	return rows.(maindef.Rows), nil
}

func (dao TranslationFindAll) SQL(fields []string) (string, error) {
	sql := "SELECT "
	i := 0
	for _, row := range fields {
		if i > 0 {
			sql += ", "
		}
		sql += row
		i++
	}
	return sql + " FROM Translation", nil
}
