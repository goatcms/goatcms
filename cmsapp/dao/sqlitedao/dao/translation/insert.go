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

// TranslationInsert is a Data Access Object for translation entity
type TranslationInsert struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
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
	return maindef.TranslationInsert(instance), nil
}

func (dao TranslationInsert) Insert(scope app.Scope, entity *entities.Translation) (id int64, err error) {
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

func (dao TranslationInsert) SQL(entity *entities.Translation) (string, error) {
	return "INSERT INTO Translation (Key, Value) VALUES (" + strconv.Quote(*entity.Key) + ", " + strconv.Quote(*entity.Value) + ")", nil
}
