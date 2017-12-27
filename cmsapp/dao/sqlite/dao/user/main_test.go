package userdao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.User {
	return &entities.User{
		Firstname: "rwUfihXRfSpPBC4OY5NM3HzLvLMupnBbytSi8p9A4nAA",
		Lastname:  "16QRNYNeTj5a8WBD0UdZy2DmcaiouWYho0jjldTrZQ11"}
}

func NewMockEntity2() *entities.User {
	return &entities.User{
		Lastname:  "hqKUxjdaOe4EbH7y8RmvZ4o10rDcsobcf4TE9OSrDQgZ",
		Firstname: "MYhVJUYgTp5k46XenfaGVyuv9H7d72gDOOsbwH6GwrqQ"}
}
