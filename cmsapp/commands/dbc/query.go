package dbc

import (
	"fmt"

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
			Model []string
		}
		rows   dao.Rows
		values []interface{}
		json   string
		in     int
	)
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	if err = ctxScope.InjectTo(&command); err != nil {
		return err
	}
	if rows, err = deps.Database.Query(nil, command.SQL); err != nil {
		return err
	}
	if response.Model, err = rows.Columns(); err != nil {
		return err
	}
	deps.Output.Printf("[")
	in = 0
	for rows.Next() {
		if in != 0 {
			deps.Output.Printf(", ")
		}
		row := map[string]interface{}{}
		if values, err = rows.GetValues(); err != nil {
			return err
		}
		for i := 0; i < len(response.Model); i++ {
			key := response.Model[i]
			value := values[i]
			if vptr, ok := value.(*interface{}); ok {
				if *vptr != nil {
					value = *vptr
				} else {
					row[key] = nil
					continue
				}
			}
			switch v := value.(type) {
			case []byte:
				row[key] = fmt.Sprintf("%s", v)
			default:
				row[key] = fmt.Sprintf("%v", v)
			}
		}
		if json, err = varutil.ObjectToJSON(row); err != nil {
			return err
		}
		deps.Output.Printf("%s", json)
		in++
	}
	deps.Output.Printf("]")
	return nil
}
