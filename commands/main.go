package commands

import (
	"github.com/urfave/cli"

	"github.com/goatcms/goatcms/commands/db"
	"github.com/goatcms/goatcms/services"
)

// InitCLI init set of default commands
func InitCLI(c *cli.App, dp services.Provider) error {
	serveCommand := NewServeCommand(dp)
	c.Commands = append(c.Commands, cli.Command{
		Name:    "serve",
		Aliases: []string{"s"},
		Usage:   "start server",
		Action:  serveCommand.Run,
	})
	err := db.InitCLI(c, dp)
	if err != nil {
		return err
	}
	return nil
}
