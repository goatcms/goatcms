package itbase

import (
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
)

func LoadFixtures(mapp app.App, path string) (err error) {
	var (
		deps struct {
			Fixture services.Fixture `dependency:"FixtureService"`
		}
	)
	if err = mapp.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	if err = deps.Fixture.Load(mapp.DependencyProvider(), mapp.AppScope(), mapp.RootFilespace(), "./database/fixtures/tests/"+path); err != nil {
		return err
	}
	if err = mapp.AppScope().Trigger(app.CommitEvent, nil); err != nil {
		return err
	}
	return nil
}
