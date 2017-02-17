package cmsapp

import (
	"github.com/goatcms/goat-core/app"
	"github.com/goatcms/goat-core/db/dsql/sqliteDSQL"
	"github.com/goatcms/goatcms/cmsapp/commands"
	"github.com/goatcms/goatcms/cmsapp/commands/dbbuildc"
	"github.com/goatcms/goatcms/cmsapp/commands/dbloadc"
	"github.com/goatcms/goatcms/cmsapp/commands/servec"
	"github.com/goatcms/goatcms/cmsapp/controllers/articlesctrl"
	"github.com/goatcms/goatcms/cmsapp/controllers/homectrl"
	"github.com/goatcms/goatcms/cmsapp/controllers/userctrl"
	"github.com/goatcms/goatcms/cmsapp/models/article"
	"github.com/goatcms/goatcms/cmsapp/models/user"
	"github.com/goatcms/goatcms/cmsapp/services/crypt"
	"github.com/goatcms/goatcms/cmsapp/services/database"
	"github.com/goatcms/goatcms/cmsapp/services/request/reqauth"
	"github.com/goatcms/goatcms/cmsapp/services/request/reqdb"
	"github.com/goatcms/goatcms/cmsapp/services/request/reqerror"
	"github.com/goatcms/goatcms/cmsapp/services/request/reqsession"
	"github.com/goatcms/goatcms/cmsapp/services/router"
	"github.com/goatcms/goatcms/cmsapp/services/session"
	"github.com/goatcms/goatcms/cmsapp/services/template"
)

const (
	TemplateFilespace = "template"
	TemplatePath      = "templates"
)

// CMSAppModule is module contains all services
type CMSAppModule struct {
}

// NewCMSAppModule create new module instance
func NewModule() app.Module {
	return &CMSAppModule{}
}

// RegisterDependency is init callback to register module dependencies
func (m *CMSAppModule) RegisterDependencies(a app.App) error {
	// filespaces
	if err := m.registerFilesystems(a); err != nil {
		return err
	}
	// commands
	if err := m.registerCommands(a); err != nil {
		return err
	}
	// services
	dp := a.DependencyProvider()
	if err := database.RegisterDependencies(dp); err != nil {
		return err
	}
	if err := session.RegisterDependencies(dp); err != nil {
		return err
	}
	if err := template.RegisterDependencies(dp); err != nil {
		return err
	}
	if err := crypt.RegisterDependencies(dp); err != nil {
		return err
	}
	if err := router.RegisterDependencies(dp); err != nil {
		return err
	}
	// models
	dsql := sqliteDSQL.NewDSQL()
	if err := user.RegisterDependencies(dp, dsql); err != nil {
		return err
	}
	if err := article.RegisterDependencies(dp, dsql); err != nil {
		return err
	}
	return nil
}

func (m *CMSAppModule) registerCommands(a app.App) error {
	commandScope := a.CommandScope()
	// serve
	commandScope.Set("help.serve", commands.ServeHelp)
	commandScope.Set("command.serve", servec.Run)
	commandScope.Set("command.s", servec.Run)
	// dbbuild
	commandScope.Set("help.dbbuild", commands.DBBuildHelp)
	commandScope.Set("command.dbbuild", dbbuildc.Run)
	commandScope.Set("command.dbb", dbbuildc.Run)
	// dbload
	commandScope.Set("help.dbbuild", commands.DBLoadHelp)
	commandScope.Set("command.dbload", dbloadc.Run)
	commandScope.Set("command.dbl", dbloadc.Run)
	return nil
}

func (m *CMSAppModule) registerFilesystems(a app.App) error {
	root := a.RootFilespace()
	templateFS, err := root.Filespace(TemplatePath)
	if err != nil {
		return err
	}
	a.FilespaceScope().Set(TemplateFilespace, templateFS)
	return nil
}

// InitDependency is init callback to inject dependencies inside module
func (m *CMSAppModule) InitDependencies(a app.App) error {
	// services
	if err := reqsession.InitDependencies(a); err != nil {
		return err
	}
	if err := template.InitDependencies(a); err != nil {
		return err
	}
	// request services
	if err := reqerror.InitDependencies(a); err != nil {
		return err
	}
	if err := reqdb.InitDependencies(a); err != nil {
		return err
	}
	if err := reqauth.InitDependencies(a); err != nil {
		return err
	}
	// controllers
	if err := userctrl.InitDependencies(a); err != nil {
		return err
	}
	if err := articlectrl.InitDependencies(a); err != nil {
		return err
	}
	if err := homectrl.InitDependencies(a); err != nil {
		return err
	}
	return nil
}

// Run is run event callback
func (m *CMSAppModule) Run() error {
	return nil
}
