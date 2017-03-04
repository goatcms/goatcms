package cmsapp

import (
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
	"github.com/goatcms/goatcms/cmsapp/services/logger"
	"github.com/goatcms/goatcms/cmsapp/services/mailer"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep/reqauth"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep/reqdb"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep/reqerror"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep/reqresponser"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep/reqsession"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep/reqtranslate"
	"github.com/goatcms/goatcms/cmsapp/services/router"
	"github.com/goatcms/goatcms/cmsapp/services/session"
	"github.com/goatcms/goatcms/cmsapp/services/template"
	"github.com/goatcms/goatcms/cmsapp/services/translate"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/db/dsql/sqliteDSQL"
)

const (
	TemplateFilespace  = "template"
	TemplatePath       = "templates"
	TranslateFilespace = "translate"
	TranslatePath      = "translates"
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
	if err := mailer.RegisterDependencies(dp); err != nil {
		return err
	}
	if err := logger.RegisterDependencies(dp); err != nil {
		return err
	}
	if err := translate.RegisterDependencies(dp); err != nil {
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
	commandScope.Set("help.command.run", commands.RunHelp)
	commandScope.Set("command.run", servec.Run)
	// dbbuild
	commandScope.Set("help.command.dbbuild", commands.DBBuildHelp)
	commandScope.Set("command.dbbuild", dbbuildc.Run)
	// dbload
	commandScope.Set("help.command.dbload", commands.DBLoadHelp)
	commandScope.Set("command.dbload", dbloadc.Run)
	// arguments
	commandScope.Set("help.argument.env", commands.EnvArg)
	commandScope.Set("help.argument.loglvl", commands.LoglvlArg)
	commandScope.Set("help.argument.host", commands.HostArg)
	return nil
}

func (m *CMSAppModule) registerFilesystems(a app.App) error {
	root := a.RootFilespace()
	// templates
	templateFS, err := root.Filespace(TemplatePath)
	if err != nil {
		return err
	}
	a.FilespaceScope().Set(TemplateFilespace, templateFS)
	// translates
	translateFS, err := root.Filespace(TranslatePath)
	if err != nil {
		return err
	}
	a.FilespaceScope().Set(TranslateFilespace, translateFS)
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
	if err := translate.InitDependencies(a); err != nil {
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
	if err := reqtranslate.InitDependencies(a); err != nil {
		return err
	}
	if err := reqresponser.InitDependencies(a); err != nil {
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
