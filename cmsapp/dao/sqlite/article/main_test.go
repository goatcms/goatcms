package articledao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.Article {
	return &entities.Article{
		Title:   "CTb3stab4tdlARstcMy54NvoWjY59VBC0VGykkqDAGNp",
		Content: "s7OWC7RRL8qRAr3qo4FG9xEjeTimpSLi0XNpHA5GCE9h"}
}

func NewMockEntity2() *entities.Article {
	return &entities.Article{
		Title:   "OHydsSjzrbHF7lOcAe3n0mmwtiijsNortzfHo2QnNW2b",
		Content: "PV1ImfSDcuZJrsZjzryArssS8Yjpk9VbRUDTUZHFKUQQ"}
}
