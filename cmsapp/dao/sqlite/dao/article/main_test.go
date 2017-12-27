package articledao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.Article {
	return &entities.Article{
		Title:   "KQc2VajZLKmaPJlxq02hyqni84myEOUCtjCm2KGVt2jd",
		Content: "JAk0AG3usiBE1NgVkIOFEUW402bnBMVRP9jO7mJkqdyH"}
}

func NewMockEntity2() *entities.Article {
	return &entities.Article{
		Title:   "nJ1NDgo0mMwLmjeTdf7oycGRDfuyy0lJKEQtCJ1X8lKn",
		Content: "zciE6xwQMbL66e7cGDsDTfeI5rZCXhXfaEBLSSw0FK1I"}
}
