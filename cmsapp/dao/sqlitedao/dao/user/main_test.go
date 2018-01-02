package dao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.User {
	var (
		email     string = "Xqw1BWEmRKshIVfir27nBuixhHZJrm3XsDZJUD5PmiZz"
		password  string = "W94EmM84BgqvX98rLINfuE1qvztSaN9YWHwLvtWq55Kh"
		firstname string = "JF5mMUgdjolW1m5EyfZ6Phh6oH4V98Vt70dluexfspaU"
		login     string = "bfJgj1u0TWvXgZwyWTc6o8ls8Q4c3Ue1RYjD16qYWnWx"
	)
	return &entities.User{
		Email:     &email,
		Password:  &password,
		Firstname: &firstname,
		Login:     &login,
	}
}

func NewMockEntity2() *entities.User {
	var (
		login     string = "VPI6rHeDmJsckwo8z0cfuwmnGUywzJBLibiX8iWjxkwC"
		password  string = "DQPsXF9C1r3zLQgekZO4ZxUaHoLksfzQdY6BB3RneAvW"
		firstname string = "8A1ZuFaz9IlAuxoWeGBlf1uxTTvzOmzBhsPuA8AZuOu9"
		email     string = "awSYN8Yg4ThoCyur97rIlAiwAusrXxKi1bQxzvDfN6LO"
	)
	return &entities.User{
		Login:     &login,
		Password:  &password,
		Firstname: &firstname,
		Email:     &email,
	}
}
