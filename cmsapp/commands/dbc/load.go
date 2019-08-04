package dbc

import (
	"os"
	"strings"

	"github.com/goatcms/goatcms/cmsapp/commands"
	"github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/filesystem"
)

// RunLoad execute db:load command
func RunLoad(a app.App, ctxScope app.Scope) (err error) {
	var (
		files []os.FileInfo
		data  []byte
		deps  struct {
			Input     app.Input            `dependency:"InputService"`
			Output    app.Output           `dependency:"OutputService"`
			Database  dao.Database         `dependency:"db0"`
			Filespace filesystem.Filespace `filespace:"root"`
			Path      string               `command:"?path"`
		}
	)
	deps.Path = commands.DefaultFixtureDir
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	if err = ctxScope.InjectTo(&deps); err != nil {
		return err
	}
	if files, err = deps.Filespace.ReadDir(deps.Path); err != nil {
		return err
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".sql") {
			if data, err = deps.Filespace.ReadFile(deps.Path + "/" + file.Name()); err != nil {
				return err
			}
			if err = deps.Database.Exec(ctxScope, string(data)); err != nil {
				return err
			}
			deps.Output.Printf(" loaded %s\n", file.Name())
		}
	}
	deps.Output.Printf("commited... ")
	if err = deps.Database.Commit(ctxScope); err != nil {
		deps.Output.Printf("fail\n")
		return err
	}
	deps.Output.Printf("ok\n")
	return nil
}
