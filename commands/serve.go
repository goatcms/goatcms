package commands

import (
	"github.com/goatcms/goatcms/services"
	"github.com/urfave/cli"
)

// ServeCommand default serve command
type ServeCommand struct {
	dp services.Provider
}

// NewServeCommand create new instance of ServeCommand
func NewServeCommand(dp services.Provider) *ServeCommand {
	return &ServeCommand{dp}
}

// Run run command
func (command *ServeCommand) Run(c *cli.Context) error {
	mux, err := command.dp.Mux()
	if err != nil {
		return err
	}
	mux.Start()
	return nil
}
