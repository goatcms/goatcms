package dbc

import (
	"github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcore/app"
)

// RunExec execute db:exec command
func RunExec(a app.App, ctxScope app.Scope) (err error) {
	var (
		deps struct {
			Input    app.Input    `dependency:"InputService"`
			Output   app.Output   `dependency:"OutputService"`
			Database dao.Database `dependency:"db0"`
		}
		command struct {
			SQL string `command:"sql"`
		}
	)
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	if err = ctxScope.InjectTo(&command); err != nil {
		return err
	}
	if err = deps.Database.Exec(nil, command.SQL); err != nil {
		return err
	}
	deps.Output.Printf("OK")
	return nil
}
