package dbloadc

import (
	"fmt"
	"strings"

	"github.com/goatcms/goatcms/cmsapp/commands"
	"github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/filesystem"
)

type Deps struct {
	Path      string               `argument:"?path"`
	Filespace filesystem.Filespace `filespace:"root"`
	AppScope  app.Scope            `dependency:"AppScope"`
	Database  dao.Database         `dependency:"db0"`
}

func Run(a app.App) error {
	deps := &Deps{
		Path: commands.DefaultFixtureDir,
	}
	if err := a.DependencyProvider().InjectTo(deps); err != nil {
		return err
	}
	files, err := deps.Filespace.ReadDir(deps.Path)
	if err != nil {
		return err
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".sql") {
			data, err := deps.Filespace.ReadFile(deps.Path + "/" + file.Name())
			if err != nil {
				return err
			}
			if err = deps.Database.Exec(deps.AppScope, string(data)); err != nil {
				return err
			}
			fmt.Printf(" loaded %s\n", file.Name())
		}
	}
	fmt.Printf("commited... ")
	if err := deps.Database.Commit(deps.AppScope); err != nil {
		fmt.Printf("fail\n")
		return err
	}
	fmt.Printf("ok\n")
	return nil
}
