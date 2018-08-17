package fragments

import (
	"bytes"
	"strings"

	"github.com/goatcms/goatcli/cliapp/services/template"
	"github.com/goatcms/goatcms/cmsapp/dao/sqlitedao"
	"github.com/goatcms/goatcms/cmsapp/services/logger"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/mockupapp"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
)

func NewTestApp(host string) (mapp app.App, err error) {
	var (
		options mockupapp.MockupOptions
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
	if options.RootFilespace, err = memfs.NewFilespace(); err != nil {
		return nil, err
	}
	if err = options.RootFilespace.MkdirAll("tmp", 0766); err != nil {
		return nil, err
	}
	if options.TMPFilespace, err = options.RootFilespace.Filespace("tmp"); err != nil {
		return nil, err
	}
	// build mockup app
	if mapp, err = mockupapp.NewApp(options); err != nil {
		return nil, err
	}
	// bootstrap
	dp := mapp.DependencyProvider()
	if err = sqlitedao.RegisterDependencies(dp); err != nil {
		return nil, err
	}
	if err = logger.RegisterDependencies(dp); err != nil {
		return nil, err
	}
	if err = template.RegisterDependencies(dp); err != nil {
		return nil, err
	}
	return mapp, nil
}
