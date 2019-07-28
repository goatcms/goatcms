package fixturec

import (
	"github.com/goatcms/goatcms/cmsapp/commands"
	"github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/filesystem"
)

// Deps contains db:fixtures:load comamnd context dependencies
type Deps struct {
	Input          app.Input            `dependency:"InputService"`
	Output         app.Output           `dependency:"OutputService"`
	Database       dao.Database         `dependency:"db0"`
	FixtureService services.Fixture     `dependency:"FixtureService"`
	Filespace      filesystem.Filespace `filespace:"root"`
	Path           string               `command:"?path"`
}

//Run execute db:fixtures:load command
func Run(a app.App, ctxScope app.Scope) (err error) {
	deps := Deps{
		Path: commands.DefaultFixtureDir,
	}
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
