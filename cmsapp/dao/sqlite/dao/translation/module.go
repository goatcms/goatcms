package translationdao

import (
	"github.com/goatcms/goatcore/dependency"
)

func RegisterDependencies(dp dependency.Provider) error {
	if err := dp.AddDefaultFactory("TranslationCreateTable", TranslationCreateTableFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("TranslationDelete", TranslationDeleteFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("TranslationDropTable", TranslationDropTableFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("TranslationFindAll", TranslationFindAllFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("TranslationFindByID", TranslationFindByIDFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("TranslationInsert", TranslationInsertFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("TranslationUpdate", TranslationUpdateFactory); err != nil {
		return err
	}
	return nil
}
