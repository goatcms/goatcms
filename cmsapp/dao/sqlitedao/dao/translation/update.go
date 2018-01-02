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

// TranslationUpdate is a Data Access Object for translation entity
type TranslationUpdate struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewTranslationUpdate(di dependency.Injector) (*TranslationUpdate, error) {
	instance := &TranslationUpdate{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func TranslationUpdateFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewTranslationUpdate(dp)
	if err != nil {
		return nil, err
	}
	return maindef.TranslationUpdate(instance), nil
}

func (dao TranslationUpdate) Update(scope app.Scope, entity *entities.Translation, fields []string) (err error) {
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

func (dao TranslationUpdate) SQL(fields []string, entity *entities.Translation) (string, error) {
	sql := "UPDATE Translation SET "
	for i, row := range fields {
		if i == 0 {
			sql += row + "="
		} else {
			sql += ", " + row + "="
		}
		switch row {
		case "Key":
			sql += strconv.Quote(*entity.Key)
		case "Value":
			sql += strconv.Quote(*entity.Value)
		}
	}
	return sql, nil
}
