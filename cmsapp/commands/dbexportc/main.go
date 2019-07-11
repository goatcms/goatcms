package dbexportc

import (
	"fmt"
	"strings"

	"github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcore/app"
)

// Run execute db:export command
func Run(a app.App, ctxScope app.Scope) error {
	var deps struct {
		DependencyScope app.Scope `dependency:"DependencyScope"`
	}
	if err := a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	keys, err := deps.DependencyScope.Keys()
	if err != nil {
		return err
	}
	for _, key := range keys {
		if strings.HasSuffix(key, "CreateTable") {
			instance, err := deps.DependencyScope.Get(key)
			if err != nil {
				return err
			}
			creator, ok := instance.(dao.CreateTable)
			if !ok {
				return fmt.Errorf("%s is not instance of dao.CreateTable", key)
			}
			query := creator.SQL()
			fmt.Printf("\n%s\n", query)
		}
	}
	return nil
}
