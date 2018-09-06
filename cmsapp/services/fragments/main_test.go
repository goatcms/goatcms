package fragments

import (
	"bytes"
	"strings"

	"github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcms/cmsapp/dao/sqlitedao"
	"github.com/goatcms/goatcms/cmsapp/services/logger"
	"github.com/goatcms/goatcms/cmsapp/services/template"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/mockupapp"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
)

func NewTestApp() (mapp app.App, err error) {
	var (
		options    mockupapp.MockupOptions
		templateFS filesystem.Filespace
		deps       struct {
			CreateTable dao.CreateTable `dependency:"FragmentCreateTable"`
		}
	)
	// Input / Output
	options.Input = gio.NewInput(strings.NewReader(""))
	options.Output = gio.NewOutput(new(bytes.Buffer))
	// Config scope
	options.ConfigScope = scope.NewScope(app.ConfigTagName)
	options.ConfigScope.Set("database.url", ":memory:")
	options.ConfigScope.Set("translate.langs", "en")
	options.ConfigScope.Set("translate.default", "en")
	// Filespaces
	if templateFS, err = memfs.NewFilespace(); err != nil {
		return nil, err
	}
	// build mockup app
	if mapp, err = mockupapp.NewApp(options); err != nil {
		return nil, err
	}
	// bootstrap
	dp := mapp.DependencyProvider()
	mapp.FilespaceScope().Set("template", templateFS)
	if err = sqlitedao.RegisterDependencies(dp); err != nil {
		return nil, err
	}
	if err = logger.RegisterDependencies(dp); err != nil {
		return nil, err
	}
	if err = template.RegisterDependencies(dp); err != nil {
		return nil, err
	}
	if err = RegisterDependencies(dp); err != nil {
		return nil, err
	}
	if err = InitDependencies(mapp); err != nil {
		return nil, err
	}
	if err = dp.InjectTo(&deps); err != nil {
		return nil, err
	}
	s := scope.NewScope("test")
	if err = deps.CreateTable.CreateTable(s); err != nil {
		return nil, err
	}
	if err = s.Trigger(app.CommitEvent, nil); err != nil {
		return nil, err
	}
	return mapp, nil
}
