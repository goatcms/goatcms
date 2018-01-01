package dao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.Article {
	return &entities.Article{
		Title:   "KTfaPv2lRcdecK5ddWuA4PaxC1IVxHy3iwy1Itvl8vDP",
		Content: "6DV4WqqwiAs36apeHhdXJ7AVXCMJTgbXPxlrkxD3YTGc"}
}

func NewMockEntity2() *entities.Article {
	return &entities.Article{
		Content: "S4XFk1u5M9NER7Apq1f3pyYJZRgNXXx4gCO0mBx6gXpb",
		Title:   "6ZKlNLlMqwvO1kAQilV8uJPD6tevDdEIGIBGXePgI5Sl"}
}
