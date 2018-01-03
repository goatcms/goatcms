package queries

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.Article {
	var (
		content string = "lLvHiDmO0Bbo24O0MIHxgxEYh7JLYJu0bzv7tYD84j9f"
		title   string = "fKE9aFPYo5WjKce7Y9VUG9PTOLyaQiU5Fq89xdl5ipSI"
	)
	return &entities.Article{
		Content: &content,
		Title:   &title,
	}
}

func NewMockEntity2() *entities.Article {
	var (
		content string = "k7UY95CNZcslZAA7UbgiqhRczv34YzRSZkBvlje7pJOl"
		title   string = "6uExWHJXfh4STocoSLBkRuUazmsnGepL1KbmqeafD7WW"
	)
	return &entities.Article{
		Content: &content,
		Title:   &title,
	}
}
