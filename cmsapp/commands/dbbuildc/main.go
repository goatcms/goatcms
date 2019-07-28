package dbbuildc

import (
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
)

// Run execute create schema command
func Run(a app.App, ctxScope app.Scope) (err error) {
	var deps struct {
		Input         app.Input              `dependency:"InputService"`
		Output        app.Output             `dependency:"OutputService"`
		SchemaCreator services.SchemaCreator `dependency:"SchemaCreator"`
	}
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	if deps.SchemaCreator.CreateSchema(); err != nil {
		return err
	}
	deps.Output.Printf("\n\nschema created\n")
	deps.Output.Printf("commited... ")
	if err = a.AppScope().Trigger(app.CommitEvent, nil); err != nil {
		return err
	}
	deps.Output.Printf("ok\n")
	return nil
}
