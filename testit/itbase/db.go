package itbase

import (
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
)

func CreateSchema(mapp app.App) (err error) {
	var deps struct {
		SchemaCreator services.SchemaCreator `dependency:"SchemaCreator"`
	}
	if err = mapp.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	if deps.SchemaCreator.CreateSchema(); err != nil {
		return err
	}
	if err = mapp.AppScope().Trigger(app.CommitEvent, nil); err != nil {
		return err
	}
	return nil
}
