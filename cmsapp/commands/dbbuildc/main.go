package dbbuildc

import (
	"fmt"

	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
)

// Run execute create schema command
func Run(a app.App) (err error) {
	var deps struct {
		SchemaCreator services.SchemaCreator `dependency:"SchemaCreator"`
	}
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	if deps.SchemaCreator.CreateSchema(); err != nil {
		return err
	}
	fmt.Printf("\n\nschema created\n")
	fmt.Printf("commited... ")
	if err := a.AppScope().Trigger(app.CommitEvent, nil); err != nil {
		return err
	}
	fmt.Printf("ok\n")
	return nil
}
