package dbc

import (
	"github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/varutil"
)

// RunQuery execute db:query command
func RunQuery(a app.App, ctxScope app.Scope) (err error) {
	var (
		deps struct {
			Input    app.Input    `dependency:"InputService"`
			Output   app.Output   `dependency:"OutputService"`
			Database dao.Database `dependency:"db0"`
		}
		command struct {
			SQL string `command:"sql"`
		}
		response struct {
			model  []string
			values [][]interface{}
		}
		rows   dao.Rows
		values []interface{}
		json   string
	)
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	if err = ctxScope.InjectTo(&command); err != nil {
		return err
	}
	response.values = make([][]interface{}, 0)
	if rows, err = deps.Database.Query(nil, command.SQL); err != nil {
		return err
	}
	if response.model, err = rows.Columns(); err != nil {
		return err
	}
	for rows.Next() {
		if values, err = rows.GetValues(); err != nil {
			return err
		}
		response.values = append(response.values, values)
	}
	if json, err = varutil.ObjectToJSON(response); err != nil {
		return err
	}
	deps.Output.Printf("%s", json)
	return nil
}
