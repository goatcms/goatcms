package db

import (
	"github.com/urfave/cli"

	"github.com/goatcms/goatcms/services"
)

// InitCLI init set of default commands
func InitCLI(c *cli.App, dp services.Provider) error {
	buildCommand := NewBuildCommand(dp)
	c.Commands = append(c.Commands, cli.Command{
		Name:    "database:build",
		Aliases: []string{"db:b"},
		Usage:   "build a database schema",
		Action:  buildCommand.Run,
	})
	return nil
}
