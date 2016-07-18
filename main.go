package main

import (
	"fmt"

	dep "github.com/goatcms/goat-core/dependency"

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
)

// App represents an application
type App struct {
	dp dep.Provider
}

// NewApp create new instance of application
func NewApp() *App {
	return &App{
		dp: dep.NewProvider(),
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
	return nil
}

func (app *App) initModels() error {
	if err := articlemodel.InitDep(app.dp); err != nil {
		return err
	}
	if err := usermodel.InitDep(app.dp); err != nil {
		return err
	}
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
	return nil
}

func (app *App) initDatabase() error {
	dbIns, err := app.dp.Get(services.DBID)
	if err != nil {
		return err
	}
	db := dbIns.(services.Database)
	if err := db.Open(); err != nil {
		return err
	}
	if err := db.CreateTables(); err != nil {
		return err
	}
	return nil
}

func (app *App) start() error {
	muxIns, err := app.dp.Get(services.MuxID)
	if err != nil {
		return err
	}
	mux := muxIns.(services.Mux)
	mux.Start()
	return nil
}

func main() {
	fmt.Println("Starting GoatCMS")
	app := NewApp()
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
	if err := app.start(); err != nil {
		fmt.Println(err)
		return
	}
}
