package articledao

import (
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	sqlitebase "github.com/goatcms/goatcms/cmsapp/dao/sqlite"
	"github.com/goatcms/goatcore/app"
)

// ArticleFindByID is a Data Access Object for article entity
type ArticleFindByID struct {
	deps struct {
		DB *sql.DB `dependency:"sqlitedb"`
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
	return maindef.FindById(instance), nil
}

func (dao ArticleFindByID) Find(scope app.Scope, fields []string, id int64) (row maindef.Row, err error) {
	var (
		sql string
		tx  *sql.Tx
	)
	if tx, err = sqlitebase.TX(scope, dao.deps.DB); err != nil {
		return err
	}
	if sql, err := dao.SQL(fields, id); err != nil {
		return nil, err
	}
	if row, err := tx.QueryRowx(sql); err != nil {
		return nil, err
	}
	return row.(maindef.Row), nil
}

func (dao ArticleFindByID) SQL(fields []string, id int) (string, error) {
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
