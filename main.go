package main

import (
	"fmt"
	"os"

	"github.com/goatcms/goat-core/filesystem/filespace/diskfs"
	"github.com/goatcms/goatcms/commands"
	"github.com/goatcms/goatcms/controllers/articles"
	"github.com/goatcms/goatcms/controllers/home"
	"github.com/goatcms/goatcms/controllers/users"
	"github.com/goatcms/goatcms/models/article"
	"github.com/goatcms/goatcms/models/user"
	"github.com/goatcms/goatcms/services"
	"github.com/goatcms/goatcms/services/auth"
	"github.com/goatcms/goatcms/services/crypt"
	"github.com/goatcms/goatcms/services/database"
	"github.com/goatcms/goatcms/services/mux"
	"github.com/goatcms/goatcms/services/session"
	"github.com/goatcms/goatcms/services/template"
	"github.com/urfave/cli"
)

const (
	uplaodPath = "./uploads/"
)

// App represents an application
type App struct {
	dp  services.Provider
	cli *cli.App
}

// NewApp create new instance of application
func NewApp() *App {
	return &App{
		dp:  services.NewProvider(),
		cli: cli.NewApp(),
	}
}

func (app *App) initDeps() error {
	if err := crypt.InitDep(app.dp); err != nil {
		return err
	}
	if err := database.InitDep(app.dp); err != nil {
		return err
	}
	if err := mux.InitDep(app.dp); err != nil {
		return err
	}
	if err := template.InitDep(app.dp); err != nil {
		return err
	}
	if err := auth.InitDep(app.dp); err != nil {
		return err
	}
	if err := session.InitDep(app.dp); err != nil {
		return err
	}
	/*if err := randomid.InitDep(app.dp); err != nil {
		return err
	}*/
	return nil
}

func (app *App) initFilespaces() error {
	uploadFilespaceFactory := diskfs.BuildFilespaceFactory(uplaodPath)
	if err := app.dp.AddService(services.UploadFilespaceID, uploadFilespaceFactory); err != nil {
		return err
	}
	return nil
}

func (app *App) initModels() error {
	if err := articlemodel.InitDep(app.dp); err != nil {
		return err
	}
	if err := usermodel.InitDep(app.dp); err != nil {
		return err
	}
	/*if err := imagemodel.InitDep(app.dp); err != nil {
		return err
	}*/
	return nil
}

func (app *App) initControllers() error {
	if err := articles.Init(app.dp); err != nil {
		return err
	}
	if err := home.Init(app.dp); err != nil {
		return err
	}
	if err := users.Init(app.dp); err != nil {
		return err
	}
	/*if err := images.Init(app.dp); err != nil {
		return err
	}*/
	return nil
}

func (app *App) initDatabase() error {
	db, err := app.dp.Database()
	if err != nil {
		return err
	}
	if err := db.Open(); err != nil {
		return err
	}
	/*if err := db.CreateTables(); err != nil {
		return err
	}*/
	return nil
}

func (app *App) initCLI() error {
	if err := commands.InitCLI(app.cli, app.dp); err != nil {
		return err
	}
	return nil
}

func (app *App) start() error {
	mux, err := app.dp.Mux()
	if err != nil {
		return err
	}
	mux.Start()
	return nil
}

func main() {
	fmt.Println("GoatCMS v0.01@dev")
	app := NewApp()
	if err := app.initFilespaces(); err != nil {
		fmt.Println(err)
		return
	}
	if err := app.initDeps(); err != nil {
		fmt.Println(err)
		return
	}
	if err := app.initModels(); err != nil {
		fmt.Println(err)
		return
	}
	if err := app.initDatabase(); err != nil {
		fmt.Println(err)
		return
	}
	if err := app.initControllers(); err != nil {
		fmt.Println(err)
		return
	}

	/*if err := app.start(); err != nil {
		fmt.Println(err)
		return
	}*/
	if err := app.initCLI(); err != nil {
		fmt.Println(err)
		return
	}
	app.cli.Run(os.Args)
}
