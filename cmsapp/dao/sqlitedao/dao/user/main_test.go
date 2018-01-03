package dao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.User {
	var (
		email     string = "4nJvJbJDA84ppNbDgIkBkQRZ1ZRFhCl6u9P3jqt7Yb9w"
		password  string = "p5FsUi7NEUGkO8gy27kuzDvdvaQBiwzwHvGfLEqnbsPV"
		firstname string = "zHHoKO74mlHQHtVQMv6GIC1seaaycH3rlh4ZcUwpmaQb"
		login     string = "MNGvJiBXXZ5p3QDTHm4Oxwx9iAu3AyJveAD04ml7pXjG"
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
		firstname string = "0gVOZLBwDxc4DQtqvcRPCvXKUbx7vIjxnyfyYgICYB5H"
		password  string = "REf4s9ddV99dPhmCxaj0eoUrNqW867i1lFQbniFhkl68"
		login     string = "bT8gTl5rQ9xzmlCS0WCzTkMPZ5eMvt2IVlzNkgbr7swT"
		email     string = "rUAFh8XPXUdQkpQRttSoIzbg2n6VdTQHlernYrnm8JzB"
	)
	return &entities.User{
		Firstname: &firstname,
		Password:  &password,
		Login:     &login,
		Email:     &email,
	}
}
