package cmsapp

import (
	"github.com/goatcms/goatcms/cmsapp/commands"
	"github.com/goatcms/goatcms/cmsapp/commands/dbbuildc"
	"github.com/goatcms/goatcms/cmsapp/commands/dbloadc"
	"github.com/goatcms/goatcms/cmsapp/commands/dbsexportc"
	"github.com/goatcms/goatcms/cmsapp/commands/servec"
	"github.com/goatcms/goatcore/app"
)

func RegisterCommands(a app.App) error {
	commandScope := a.CommandScope()
	// serve
	commandScope.Set("help.command.run", commands.RunHelp)
	commandScope.Set("command.run", servec.Run)
	// dbbuild
	commandScope.Set("help.command.dbbuild", commands.DBBuildHelp)
	commandScope.Set("command.dbbuild", dbbuildc.Run)
	// dbbuild
	commandScope.Set("help.command.dbsexport", commands.DBExportHelp)
	commandScope.Set("command.db_schema_export", dbsexportc.Run)
	// dbload
	commandScope.Set("help.command.dbload", commands.DBLoadHelp)
	commandScope.Set("command.dbload", dbloadc.Run)
	// arguments
	commandScope.Set("help.argument.env", commands.EnvArg)
	commandScope.Set("help.argument.loglvl", commands.LoglvlArg)
	commandScope.Set("help.argument.host", commands.HostArg)
	return nil
}
