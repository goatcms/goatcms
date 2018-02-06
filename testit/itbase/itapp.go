package itbase

import (
	"bytes"
	"flag"
	"strings"

	"github.com/goatcms/goatcms/cmsapp"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/bootstrap"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/mockupapp"
	"github.com/goatcms/goatcore/app/modules/terminal"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/goatcms/goatcore/filesystem/filespace/diskfs"
)

func NewITApp(host string) (mapp app.App, err error) {
	var (
		options      mockupapp.MockupOptions
		smtpUsername = flag.String("smtpuser", "", "smtp username for test connection")
		smtpPassword = flag.String("smtpass", "", "smtp password for test connection")
	)
	// Input / Output
	options.Input = gio.NewInput(strings.NewReader(""))
	options.Output = gio.NewOutput(new(bytes.Buffer))
	// Config scope
	options.ConfigScope = scope.NewScope(app.ConfigTagName)
	options.ConfigScope.Set("database.url", ":memory:")
	options.ConfigScope.Set("translate.langs", "en")
	options.ConfigScope.Set("translate.default", "en")
	options.ConfigScope.Set("router.host", host)
	options.ConfigScope.Set("mailer.smtp.addr", "smtp.gmail.com:465")
	options.ConfigScope.Set("mailer.smtp.auth.username", *smtpUsername)
	options.ConfigScope.Set("mailer.smtp.auth.password", *smtpPassword)
	// Filespaces
	if options.RootFilespace, err = diskfs.NewFilespace("../"); err != nil {
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
	bootstrap := bootstrap.NewBootstrap(mapp)
	if err = bootstrap.Register(terminal.NewModule()); err != nil {
		return nil, err
	}
	if err = bootstrap.Register(cmsapp.NewModule()); err != nil {
		return nil, err
	}
	if err := bootstrap.Init(); err != nil {
		return nil, err
	}
	if err := bootstrap.Run(); err != nil {
		return nil, err
	}
	return mapp, nil
}
