package dao

import (
	"github.com/goatcms/goatcore/dependency"
)

func RegisterDependencies(dp dependency.Provider) error {
	if err := dp.AddDefaultFactory("FragmentCreateTable", FragmentCreateTableFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("FragmentDelete", FragmentDeleteFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("FragmentDropTable", FragmentDropTableFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("FragmentFindAll", FragmentFindAllFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("FragmentFindByID", FragmentFindByIDFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("FragmentInsert", FragmentInsertFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("FragmentUpdate", FragmentUpdateFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("FragmentSearch", FragmentSearchFactory); err != nil {
		return err
	}
	return nil
}
