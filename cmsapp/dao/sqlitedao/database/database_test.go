package database

import (
	"bytes"
	"strings"
	"testing"

	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/mockupapp"
	"github.com/goatcms/goatcore/app/scope"
)

func TestFactory(t *testing.T) {
	var (
		err  error
		mapp app.App
	)
	t.Parallel()
	// prepare mockup application
	configScope := scope.NewScope(app.ConfigTagName)
	configScope.Set("database.url", ":memory:")
	if mapp, err = mockupapp.NewApp(mockupapp.MockupOptions{
		Input:       gio.NewInput(strings.NewReader("")),
		Output:      gio.NewOutput(new(bytes.Buffer)),
		ConfigScope: configScope,
	}); err != nil {
		t.Error(err)
		return
	}
	if err = mapp.DependencyProvider().AddDefaultFactory("db0.engine", EngineFactory); err != nil {
		t.Error(err)
		return
	}
	if err = mapp.DependencyProvider().AddDefaultFactory("db0", Factory); err != nil {
		t.Error(err)
		return
	}
	// test
	var deps struct {
		DB maindef.Database `dependency:"db0"`
	}
	if err = mapp.DependencyProvider().InjectTo(&deps); err != nil {
		t.Error(err)
		return
	}
	if deps.DB == nil {
		t.Error("should inject database instance to DB property (during injection)")
		return
	}
}
