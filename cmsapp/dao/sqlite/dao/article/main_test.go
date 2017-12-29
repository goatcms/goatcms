package dao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.Article {
	return &entities.Article{
		Title:   "TVBHlBbMYAPcOsexBfPT1wasYSmOpTPaMHLRso9MSxAD",
		Content: "sHawAYvPBl35EBTxPM3oXi3ro2JlXEBOazkl3YsaWynP"}
}

func NewMockEntity2() *entities.Article {
	return &entities.Article{
		Content: "9dfAov9McIwqpMUAxV15C12ZrBf6GmsRB7ua0nnmSmnq",
		Title:   "R5p6xKBwipUqDvdYC8nWGYKMYFgz8BKUNT1THc7xX7As"}
}
