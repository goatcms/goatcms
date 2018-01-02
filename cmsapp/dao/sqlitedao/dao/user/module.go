package dao

import (
	"github.com/goatcms/goatcore/dependency"
)

func RegisterDependencies(dp dependency.Provider) error {
	if err := dp.AddDefaultFactory("UserCreateTable", UserCreateTableFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("UserDelete", UserDeleteFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("UserDropTable", UserDropTableFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("UserFindAll", UserFindAllFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("UserFindByID", UserFindByIDFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("UserInsert", UserInsertFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("UserUpdate", UserUpdateFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("UserSearch", UserSearchFactory); err != nil {
		return err
	}
	return nil
}
