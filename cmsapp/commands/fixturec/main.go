package fixturec

import (
	"fmt"

	"github.com/goatcms/goatcms/cmsapp/commands"
	"github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/filesystem"
)

// Deps contains db:fixtures:load comamnd context dependencies
type Deps struct {
	Path           string               `command:"?path"`
	Filespace      filesystem.Filespace `filespace:"root"`
	AppScope       app.Scope            `dependency:"AppScope"`
	Database       dao.Database         `dependency:"db0"`
	FixtureService services.Fixture     `dependency:"FixtureService"`
}

//Run execute db:fixtures:load command
func Run(a app.App, ctxScope app.Scope) (err error) {
	deps := &Deps{
		Path: commands.DefaultFixtureDir,
	}
	if err = a.DependencyProvider().InjectTo(deps); err != nil {
		return err
	}
	if err = deps.FixtureService.Load(a.DependencyProvider(), deps.AppScope, deps.Filespace, deps.Path); err != nil {
		return err
	}
	if err = deps.Database.Commit(deps.AppScope); err != nil {
		return err
	}
	fmt.Printf("ok\n")
	return nil
}
