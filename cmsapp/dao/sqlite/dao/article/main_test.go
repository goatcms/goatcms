package articledao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.Article {
	return &entities.Article{
		Title:   "7isUzEpJEbkHtLyoU6yoyGb3Hs940Bx33KsysGc8rMqx",
		Content: "fnuR72mtw0nbAZl27zwPweR9ls3sJVIaX88EjJtyF8jM"}
}

func NewMockEntity2() *entities.Article {
	return &entities.Article{
		Content: "q8z1zOyQNL3h3QzZRBcoVzu03srz9INgri8fPSqSFdzc",
		Title:   "CcT4iFXEpcH9euJq4tWuftpP5TtgyrNlr4FqzpmUYupW"}
}
