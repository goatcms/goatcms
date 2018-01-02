package dao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.Article {
	var (
		content string = "eOr7W1Z40kE6fyMyKKb8E2CezUUmYhVmbPuKXxGyifJi"
		title   string = "jHxgvK0HA5Plv9hvgysp5dXEQp2EPcHtjARDbLUTSCdz"
	)
	return &entities.Article{
		Content: &content,
		Title:   &title,
	}
}

func NewMockEntity2() *entities.Article {
	var (
		title   string = "4cxndrwTVqBjhOPzhLvTBDzaxfDrYMKOlqJMFbp4xUg0"
		content string = "vDiphKeJCyZt0Zcdy4yOOcy8pFzPvD9yqoBFe4bVjtaW"
	)
	return &entities.Article{
		Title:   &title,
		Content: &content,
	}
}
