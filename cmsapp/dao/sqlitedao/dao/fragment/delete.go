package dao

import (
	"database/sql"
	"fmt"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"strconv"
)

// FragmentDelete is a Data Access Object for fragment entity
type FragmentDelete struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewFragmentDelete(di dependency.Injector) (*FragmentDelete, error) {
	instance := &FragmentDelete{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func FragmentDeleteFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewFragmentDelete(dp)
	if err != nil {
		return nil, err
	}
	return maindef.Delete(instance), nil
}

func (dao FragmentDelete) Delete(scope app.Scope, id int64) error {
	var (
		res   sql.Result
		err   error
		count int64
		tx    *sql.Tx
	)
	if tx, err = helpers.TX(scope, dao.deps.DB); err != nil {
		return err
	}
	sql := dao.SQL(id)
	if res, err = tx.Exec(sql); err != nil {
		return fmt.Errorf("%s: %s", err.Error(), sql)
	}
	if count, err = res.RowsAffected(); err != nil {
		return fmt.Errorf("%s: %s", err.Error(), sql)
	}
	if count != 1 {
		return fmt.Errorf("Delete more than one record (%v records deleted)", count)
	}
	return nil
}

func (dao FragmentDelete) SQL(id int64) string {
	return "DELETE FROM Fragment WHERE " + strconv.FormatInt(id, 10)
}
