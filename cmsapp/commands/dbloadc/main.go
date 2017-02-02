package dbloadc

import (
	"fmt"
	"strings"

	"github.com/goatcms/goat-core/app"
	"github.com/goatcms/goat-core/filesystem"
	"github.com/goatcms/goatcms/cmsapp/commands"
	"github.com/goatcms/goatcms/cmsapp/services"
)

type Deps struct {
	Path      string               `argument:"?path"`
	Filespace filesystem.Filespace `filespace:"root"`
	Database  services.Database    `dependency:"DatabaseService"`
}

func Run(a app.App) error {
	deps := &Deps{
		Path: commands.DefaultFixtureDir,
	}
	if err := a.DependencyProvider().InjectTo(deps); err != nil {
		return err
	}
	tx, err := deps.Database.TX()
	if err != nil {
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
			tx.MustExec(string(data))
			fmt.Printf(" loaded %s\n", file.Name())
		}
	}
	fmt.Printf("commited... ")
	if err := tx.Commit(); err != nil {
		fmt.Printf("fail\n")
		return err
	}
	fmt.Printf("ok\n")
	return nil
}
