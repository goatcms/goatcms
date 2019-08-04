package dbc

import (
	"github.com/goatcms/goatcms/cmsapp/commands"
	"github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/filesystem"
)

//Run execute db:fixtures:load command
func RunFixturesLoad(a app.App, ctxScope app.Scope) (err error) {
	var (
		deps struct {
			Input          app.Input            `dependency:"InputService"`
			Output         app.Output           `dependency:"OutputService"`
			Database       dao.Database         `dependency:"db0"`
			FixtureService services.Fixture     `dependency:"FixtureService"`
			Filespace      filesystem.Filespace `filespace:"root"`
			Path           string               `command:"?path"`
		}
	)
	deps.Path = commands.DefaultFixtureDir
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	if err = ctxScope.InjectTo(&deps); err != nil {
		return err
	}
	if err = deps.FixtureService.Load(a.DependencyProvider(), ctxScope, deps.Filespace, deps.Path); err != nil {
		return err
	}
	if err = deps.Database.Commit(ctxScope); err != nil {
		return err
	}
	deps.Output.Printf("ok\n")
	return nil
}
