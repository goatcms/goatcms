package dbbuildc

import (
	"fmt"
	"strings"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/db"
	"github.com/goatcms/goatcms/cmsapp/services"
)

func Run(a app.App) error {
	var deps struct {
		Database        services.Database `dependency:"DatabaseService"`
		DependencyScope app.Scope         `dependency:"DependencyScope"`
	}
	if err := a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	tx, err := deps.Database.TX()
	if err != nil {
		return err
	}
	keys, err := deps.DependencyScope.Keys()
	if err != nil {
		return err
	}
	for _, key := range keys {
		if strings.HasPrefix(key, "db.query.") && strings.HasSuffix(key, ".CreateTable") {
			createTableIns, err := deps.DependencyScope.Get(key)
			if err != nil {
				return err
			}
			createTable := createTableIns.(db.CreateTable)
			err = createTable(tx)
			if err != nil {
				return err
			}
			fmt.Printf("%s runed \n", key)
		}
	}
	fmt.Printf("created all tables\n")
	fmt.Printf("commited... ")
	if err := tx.Commit(); err != nil {
		fmt.Printf("fail\n")
		return err
	}
	fmt.Printf("ok\n")
	return nil
}
