package dao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.Article {
	return &entities.Article{
		Content: "gKHz7DtuOMfJf7cxHMTLedAUa8iICUpm5Fg3GG6odLPt",
		Title:   "ktiuxJVB0ugsB69Bi0yzqZ9OcUW2PESYgHXZEL75Jno2"}
}

func NewMockEntity2() *entities.Article {
	return &entities.Article{
		Content: "evMtMAMl8fqHQKsIdchyILpcDutu1HKsO10Dl07osg1B",
		Title:   "V1DbqJTRJH9CApIJd0zduRVYoQIueSnOqhBc7dN3OXOq"}
}
