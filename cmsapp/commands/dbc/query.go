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
			Model  []string
			Values []map[string]interface{}
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
	response.Values = make([]map[string]interface{}, 0)
	if rows, err = deps.Database.Query(nil, command.SQL); err != nil {
		return err
	}
	if response.Model, err = rows.Columns(); err != nil {
		return err
	}
	for rows.Next() {
		row := map[string]interface{}{}
		if values, err = rows.GetValues(); err != nil {
			return err
		}
		for i := 0; i < len(response.Model); i++ {
			key := response.Model[i]
			value := values[i]
			if vptr, ok := value.(*interface{}); ok {
				fmt.Printf("map: %v %v\n", value, *vptr)
				value = *vptr
			}
			switch v := value.(type) {
			case string:
			case []byte:
				row[key] = fmt.Sprintf("%s", v)
			case *string:
			case *[]byte:
				row[key] = fmt.Sprintf("%s", *v)
			case int64:
			case int32:
			case int:
				row[key] = fmt.Sprintf("%d", v)
			case *int64:
			case *int32:
			case *int:
				row[key] = fmt.Sprintf("%d", *v)
			default:
				row[key] = fmt.Sprintf("%v", v)
			}
		}
		response.Values = append(response.Values, row)
	}
	if json, err = varutil.ObjectToJSON(response.Values); err != nil {
		return err
	}
	deps.Output.Printf("%s", json)
	return nil
}
