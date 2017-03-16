package dbsexportc

import (
	"fmt"
	"strings"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/db"
)

func Run(a app.App) error {
	var deps struct {
		DependencyScope app.Scope `dependency:"DependencyScope"`
		DSQL            db.DSQL   `dependency:"DSQL"`
	}
	if err := a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	keys, err := deps.DependencyScope.Keys()
	if err != nil {
		return err
	}
	for _, key := range keys {
		if strings.HasSuffix(key, "Table") {
			tableIns, err := deps.DependencyScope.Get(key)
			if err != nil {
				return err
			}
			table, ok := tableIns.(db.Table)
			if !ok {
				return fmt.Errorf("%s is not instance of db.Table", key)
			}
			query, err := deps.DSQL.NewCreateSQL(table.Name(), table.Types())
			if err != nil {
				return err
			}
			fmt.Printf("\n%s\n", query)
		}
	}
	return nil
}
