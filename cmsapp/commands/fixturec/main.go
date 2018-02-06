package dbloadc

import (
	"fmt"

	"github.com/goatcms/goatcms/cmsapp/commands"
	"github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/filesystem"
)

type Deps struct {
	Path           string               `argument:"?path"`
	Filespace      filesystem.Filespace `filespace:"root"`
	AppScope       app.Scope            `dependency:"AppScope"`
	Database       dao.Database         `dependency:"db0"`
	FixtureService services.Fixture     `dependency:"FixtureService"`
}

func Run(a app.App) (err error) {
	deps := &Deps{
		Path: commands.DefaultFixtureDir,
	}
	if err = a.DependencyProvider().InjectTo(deps); err != nil {
		return err
	}
	if err = deps.FixtureService.Load(deps.AppScope, deps.Filespace, deps.Path); err != nil {
		return err
	}
	if err = deps.Database.Commit(deps.AppScope); err != nil {
		return err
	}
	fmt.Printf("ok\n")
	return nil
}
