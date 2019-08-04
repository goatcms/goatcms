package dbc

import (
	"fmt"
	"strings"

	"github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcore/app"
)

// RunExport execute db:export command
func RunExport(a app.App, ctxScope app.Scope) (err error) {
	var (
		deps struct {
			Input           app.Input  `dependency:"InputService"`
			Output          app.Output `dependency:"OutputService"`
			DependencyScope app.Scope  `dependency:"DependencyScope"`
		}
		keys     []string
		instance interface{}
		creator  dao.CreateTable
		ok       bool
	)
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	if keys, err = deps.DependencyScope.Keys(); err != nil {
		return err
	}
	for _, key := range keys {
		if strings.HasSuffix(key, "CreateTable") {
			if instance, err = deps.DependencyScope.Get(key); err != nil {
				return err
			}
			if creator, ok = instance.(dao.CreateTable); !ok {
				return fmt.Errorf("%s is not instance of dao.CreateTable", key)
			}
			deps.Output.Printf("\n%s\n", creator.SQL())
		}
	}
	return nil
}
