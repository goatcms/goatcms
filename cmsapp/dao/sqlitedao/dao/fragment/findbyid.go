package dao

import (
	"database/sql"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	entities "github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"strconv"
)

// FragmentFindByID is a Data Access Object for fragment entity
type FragmentFindByID struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewFragmentFindByID(di dependency.Injector) (*FragmentFindByID, error) {
	instance := &FragmentFindByID{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func FragmentFindByIDFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewFragmentFindByID(dp)
	if err != nil {
		return nil, err
	}
	return maindef.FragmentFindByID(instance), nil
}

func (dao FragmentFindByID) Find(scope app.Scope, fields []string, id int64) (*entities.Fragment, error) {
	var (
		err   error
		query string
		tx    *sql.Tx
		row   *FragmentRow
	)
	if tx, err = helpers.TX(scope, dao.deps.DB); err != nil {
		return nil, err
	}
	if query, err = dao.SQL(fields, id); err != nil {
		return nil, err
	}
	row = NewFragmentRow(tx.QueryRow(query), fields)
	return row.Get()
}

func (dao FragmentFindByID) SQL(fields []string, id int64) (string, error) {
	sql := "SELECT "
	i := 0
	for _, row := range fields {
		if i > 0 {
			sql += ", "
		}
		sql += row
		i++
	}
	return sql + " FROM Fragment WHERE id=" + strconv.FormatInt(id, 10) + " LIMIT 1", nil
}
