package userdao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.User {
	return &entities.User{
		Firstname: "cvbP36HeWTgkdbPlwXltcdX9dE7nrRxfAZSRi2TgQdaA",
		Lastname:  "sMH2TgAnUW34Jc8hcJdVh3C5ghPHIt8tEg92v3ofFDCi"}
}

func NewMockEntity2() *entities.User {
	return &entities.User{
		Lastname:  "KzKcC0D8CtInC8wZMOo3vZehGVHt5hDfOaAjxLyJay3z",
		Firstname: "v4tkIc2mQ0j9cPBW9wE8QqyHKU8xvcFNXwR27J1JHl7J"}
}
