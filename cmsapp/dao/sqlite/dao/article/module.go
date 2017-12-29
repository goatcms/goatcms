package dao

import (
	"github.com/goatcms/goatcore/dependency"
)

func RegisterDependencies(dp dependency.Provider) error {
	if err := dp.AddDefaultFactory("ArticleCreateTable", ArticleCreateTableFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("ArticleDelete", ArticleDeleteFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("ArticleDropTable", ArticleDropTableFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("ArticleFindAll", ArticleFindAllFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("ArticleFindByID", ArticleFindByIDFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("ArticleInsert", ArticleInsertFactory); err != nil {
		return err
	}
	if err := dp.AddDefaultFactory("ArticleUpdate", ArticleUpdateFactory); err != nil {
		return err
	}
	return nil
}
