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
)

// FragmentInsert is a Data Access Object for fragment entity
type FragmentInsert struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewFragmentInsert(di dependency.Injector) (*FragmentInsert, error) {
	instance := &FragmentInsert{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func FragmentInsertFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewFragmentInsert(dp)
	if err != nil {
		return nil, err
	}
	return maindef.FragmentInsert(instance), nil
}

func (dao FragmentInsert) Insert(scope app.Scope, entity *entities.Fragment) (id int64, err error) {
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
	entity.ID = &id
	return id, nil
}

func (dao FragmentInsert) SQL(entity *entities.Fragment) (string, error) {
	sql := "INSERT INTO Fragment ("
	if entity.ID != nil {
		sql += "ID, "
	}
	sql += "Lang, Name, Content, ID, ID, ID) VALUES ("
	if entity.ID != nil {
		sql += strconv.FormatInt(*entity.ID, 10) + ", "
	}
	sql += "" + helpers.Quote(entity.Lang) + ", " + helpers.Quote(entity.Name) + ", " + helpers.Quote(entity.Content) + ", " + helpers.FormatInt(entity.ID, 10) + ", " + helpers.FormatInt(entity.ID, 10) + ", " + helpers.FormatInt(entity.ID, 10) + ")"
	return sql, nil
}
