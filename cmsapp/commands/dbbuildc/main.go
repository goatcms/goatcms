package dbbuildc

import (
	"fmt"
	"strings"

	"github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcore/app"
)

func Run(a app.App) error {
	var deps struct {
		//Database        services.Database `dependency:"DatabaseService"`
		DependencyScope app.Scope `dependency:"DependencyScope"`
		AppScope        app.Scope `dependency:"AppScope"`
	}
	if err := a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	keys, err := deps.DependencyScope.Keys()
	if err != nil {
		return err
	}
	for _, key := range keys {
		if strings.HasSuffix(key, "CreateTable") {
			tableIns, err := deps.DependencyScope.Get(key)
			if err != nil {
				return err
			}
			creator, ok := tableIns.(dao.CreateTable)
			if !ok {
				return fmt.Errorf("%s is not instance of db.Table", key)
			}
			if err := creator.CreateTable(deps.AppScope); err != nil {
				return err
			}
			fmt.Printf("\n %v... success", key)
		}
	}
	fmt.Printf("\n\ncreated all tables\n")
	fmt.Printf("commited... ")
	if err := deps.AppScope.Trigger(app.CommitEvent, nil); err != nil {
		fmt.Printf("fail %v\n", err)
		return nil
	}
	fmt.Printf("ok\n")
	return nil
}
