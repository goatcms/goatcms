package dao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.User {
	var (
		roles     string = "KiYKD8UffbCBbSUYoBWAb8CpO9bhZwNrIeUoNHviX7ur"
		firstname string = "FvDrmzBFO6yHXGj9m5OVjfn8aztO1O6EhJ0KQ375y44G"
		lastname  string = "FWONPVT2n6AEgrtg2zS5GHNDpMc6pK40GTcMFNakRxhj"
		password  string = "kQbm1mIIDuUXzA1mNnV7lL2esbb8NZlLT3mGPjOv5OFL"
		login     string = "6DscMWKbkDm9kgf9l3wMAFzpNqMt5ZwNRQUKPPkatQce"
		email     string = "Jessica.Robinson@jwy.pl"
	)
	return &entities.User{
		Roles:     &roles,
		Firstname: &firstname,
		Lastname:  &lastname,
		Password:  &password,
		Login:     &login,
		Email:     &email,
	}
}

func NewMockEntity2() *entities.User {
	var (
		lastname  string = "FOJVBU6J1ky5GAQfWaWaDJVKIcBFOtSpNHVBUl6aCvMJ"
		email     string = "Alfie.Thomas@ods.fi"
		password  string = "TcoA3goYPEqaw5aFHsH79oyiUq5Zozq5grNSVLc7jloZ"
		firstname string = "r9QrFALVwfu2XIfbfUMh1S5YlUbex3zXr5XiRNybes2U"
		login     string = "haYMKsuM8rxYqvSEKSfRUiaxsv8eCOJf4C96z1OCIE1Q"
		roles     string = "4M9rRWjOKHxRYJwegVA4VazRjUDSuQPsm1h2zJdre8LY"
	)
	return &entities.User{
		Lastname:  &lastname,
		Email:     &email,
		Password:  &password,
		Firstname: &firstname,
		Login:     &login,
		Roles:     &roles,
	}
}
