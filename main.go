package main

import (
	"log"
	"os"

	"github.com/goatcms/goatcms/cmsapp"
	"github.com/goatcms/goatcore/app/bootstrap"
	"github.com/goatcms/goatcore/app/goatapp"
	"github.com/goatcms/goatcore/app/modules/terminal"
)

//go:generate goatcli build
//go:generate dep ensure
//go:generate go generate -x github.com/goatcms/goatcms/web
func main() {
	errLogs := log.New(os.Stderr, "", 0)
	app, err := goatapp.NewGoatApp("GoatCMS", "0.0.1", "./")
	if err != nil {
		errLogs.Println(err)
		os.Exit(1)
		return
	}

	bootstrap := bootstrap.NewBootstrap(app)
	if err = bootstrap.Register(terminal.NewModule()); err != nil {
		errLogs.Println(err)
		os.Exit(1)
		return
	}
	if err = bootstrap.Register(cmsapp.NewModule()); err != nil {
		errLogs.Println(err)
		os.Exit(1)
		return
	}

	if err := bootstrap.Init(); err != nil {
		errLogs.Println(err)
		os.Exit(1)
		return
	}
	if err := bootstrap.Run(); err != nil {
		errLogs.Println(err)
		os.Exit(1)
		return
	}
}
