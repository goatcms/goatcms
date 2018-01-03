package dao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.Article {
	var (
		content string = "cwL151F3YPC100DLcV0TVDd9jQc8d2bsZhjVcBX7W47S"
		title   string = "p35CHyV5oAXJGAIjjGLFGZckATv5MHZObUgiB658xyml"
	)
	return &entities.Article{
		Content: &content,
		Title:   &title,
	}
}

func NewMockEntity2() *entities.Article {
	var (
		content string = "309Uiew5d6b1lvjQatmlislR7wByc6bIUw2m1sRk1nN8"
		title   string = "s4NFjHtRRwKPkN5XIw58xRc0lAls5UoDwWRxaf9oKERi"
	)
	return &entities.Article{
		Content: &content,
		Title:   &title,
	}
}
