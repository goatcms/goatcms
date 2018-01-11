package dao

import (
	"github.com/goatcms/goatcore/dependency"
)

func RegisterDependencies(dp dependency.Provider) error {
	if err := dp.AddDefaultFactory("SessionCreateTable", SessionCreateTableFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("SessionDelete", SessionDeleteFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("SessionDropTable", SessionDropTableFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("SessionFindAll", SessionFindAllFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("SessionFindByID", SessionFindByIDFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("SessionInsert", SessionInsertFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("SessionUpdate", SessionUpdateFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("SessionSearch", SessionSearchFactory); err != nil {
		return err
	}
	return nil
}
