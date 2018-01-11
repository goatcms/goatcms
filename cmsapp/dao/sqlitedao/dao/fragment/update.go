package dao

import (
	"database/sql"
	"fmt"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	entities "github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

// FragmentUpdate is a Data Access Object for fragment entity
type FragmentUpdate struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewFragmentUpdate(di dependency.Injector) (*FragmentUpdate, error) {
	instance := &FragmentUpdate{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func FragmentUpdateFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewFragmentUpdate(dp)
	if err != nil {
		return nil, err
	}
	return maindef.FragmentUpdate(instance), nil
}

func (dao FragmentUpdate) Update(scope app.Scope, entity *entities.Fragment, fields []string) (err error) {
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

func (dao FragmentUpdate) SQL(fields []string, entity *entities.Fragment) (string, error) {
	sql := "UPDATE Fragment SET "
	for i, row := range fields {
		if i == 0 {
			sql += row + "="
		} else {
			sql += ", " + row + "="
		}
		switch row {
		case "Lang":
			sql += helpers.Quote(entity.Lang)
		case "Name":
			sql += helpers.Quote(entity.Name)
		case "Content":
			sql += helpers.Quote(entity.Content)
		}
	}
	return sql, nil
}
