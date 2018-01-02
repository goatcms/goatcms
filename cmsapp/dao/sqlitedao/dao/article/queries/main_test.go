package queries

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.Article {
	var (
		content string = "oHs1mf5Iy6lPNJBkp1ZfPDtoOL91FNiVGzVSCmm2ATye"
		title   string = "AzudGr2qJikdhRkwnPXcVWeXgoPJp8qLaSgxqKEaEkyM"
	)
	return &entities.Article{
		Content: &content,
		Title:   &title,
	}
}

func NewMockEntity2() *entities.Article {
	var (
		content string = "R9DRNH61Fxxuz50nLfj425VlC9meBJNPjcYKGNIX5Abj"
		title   string = "d0EzRxpDL7IdgbNMGIK6z8PkN48c0cdMHmB9gLxatlv3"
	)
	return &entities.Article{
		Content: &content,
		Title:   &title,
	}
}
