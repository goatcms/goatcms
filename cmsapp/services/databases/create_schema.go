package databases

import (
	"fmt"
	"strings"

	"github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

// SchemaCreator is object to create database
type SchemaCreator struct {
	dp dependency.Provider
}

// DatabaseCreatorFactory create n
func SchemaCreatorFactory(dp dependency.Provider) (interface{}, error) {
	return services.SchemaCreator(&SchemaCreator{
		dp: dp,
	}), nil
}

func (creator *SchemaCreator) CreateSchema() (err error) {
	var (
		ok           bool
		tableIns     interface{}
		tableCreator dao.CreateTable
		keys         []string
		deps         struct {
			DependencyScope app.Scope `dependency:"DependencyScope"`
			AppScope        app.Scope `dependency:"AppScope"`
		}
	)
	if err := creator.dp.InjectTo(&deps); err != nil {
		return err
	}
	if keys, err = deps.DependencyScope.Keys(); err != nil {
		return err
	}
	for _, key := range keys {
		if strings.HasSuffix(key, "CreateTable") {
			if tableIns, err = deps.DependencyScope.Get(key); err != nil {
				return err
			}
			if tableCreator, ok = tableIns.(dao.CreateTable); !ok {
				return fmt.Errorf("%s is not instance of db.Table", key)
			}
			if err = tableCreator.CreateTable(deps.AppScope); err != nil {
				return err
			}
		}
	}
	if err = deps.AppScope.Trigger(app.CommitEvent, nil); err != nil {
		return err
	}
	return nil
}
