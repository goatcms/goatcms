package db

import (
	"fmt"
	"strings"

	"github.com/goatcms/goat-core/db"
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/services"
	"github.com/urfave/cli"
)

// BuildCommand default serve command
type BuildCommand struct {
	dp services.Provider
}

// NewBuildCommand create new instance of BuildCommand
func NewBuildCommand(dp services.Provider) *BuildCommand {
	return &BuildCommand{dp}
}

// Run run command
func (command *BuildCommand) Run(c *cli.Context) error {
	var (
		daoIns dependency.Instance
		dao    db.DAO
		err    error
	)
	database, err := command.dp.Database()
	if err != nil {
		return err
	}

	tx, err := database.Adapter().Beginx()
	if err != nil {
		return err
	}

	builders := command.dp.GetAll()
	for name, builder := range builders {
		if strings.HasPrefix(name, "dao.") {
			daoIns, err = builder.Get(command.dp)
			if err != nil {
				return err
			}
			dao = daoIns.(db.DAO)
			err = dao.CreateTable(tx)
			if err != nil {
				return err
			}
			fmt.Println(" created ", name)
		}
	}
	fmt.Println("Create alltables")
	return tx.Commit()
}
