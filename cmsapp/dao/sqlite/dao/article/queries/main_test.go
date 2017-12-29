package queries

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.Article {
	return &entities.Article{
		Content: "io4GVNKVS5c0EA2QpzvDzPLBn8QNnHccae3fObd2pTkW",
		Title:   "71VP84NHsldayeAsu3uBNFwVRY0CXDisxHiRq0C0LmRK"}
}

func NewMockEntity2() *entities.Article {
	return &entities.Article{
		Title:   "yDKnfL22QVbWU8rJYZFP59hE0Ghh7xTmwbP6CuJQucZ2",
		Content: "VYtf2WA8IEiFAfrYtOtQl1OmQkmJSoQ7DNdcVQUJ10Yt"}
}
