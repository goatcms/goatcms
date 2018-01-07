package queries

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.User {
	var (
		email     string = "Thomas.Roberts@eoc.cn"
		lastname  string = "a5qxZ8MlEwmZvaTgiC0C53w7mNq8qk7OAZWVxFMNyJJt"
		roles     string = "kobNHpZId7KQ1mWQuurEZbnWvAMbgmrRQtsFtIOHm2Oa"
		firstname string = "Fa6pMOYVdqA83lBxmwQb8ObvDBaEdHtfQRVUY3kwIulj"
		username  string = "yW5YNkaRuqRRG0HI4RiHQxJ5eCCpsTTLn76jibdZjlWi"
		password  string = "4dzuOKmgVvJ1UDUf1TKpyssaUhLwEZ3AkDOehGeLtPhe"
	)
	return &entities.User{
		Email:     &email,
		Lastname:  &lastname,
		Roles:     &roles,
		Firstname: &firstname,
		Username:  &username,
		Password:  &password,
	}
}

func NewMockEntity2() *entities.User {
	var (
		password  string = "MYzY0e5y0N8JZtU4MqwwKcXMs4SRtWpeQgHIUvDbZDNx"
		firstname string = "5jZIx8T9Wm9Wx91ehjrubNP51LkwMzTMjZY7jAtKeUsJ"
		email     string = "Amelia.Robinson@asi.pl"
		lastname  string = "jZjf9qClWaAbC9NLkdI83zI1uV8Yq8krSDRmbeu7DG3m"
		roles     string = "4kRO2hzMzA9iAMTRcR7vkqFxkvd77ogdNdvJbKKEy1km"
		username  string = "aQ7fQqmhlBuNgqFodPuPdfkEA0jRNzodvbY1uxKqGofQ"
	)
	return &entities.User{
		Password:  &password,
		Firstname: &firstname,
		Email:     &email,
		Lastname:  &lastname,
		Roles:     &roles,
		Username:  &username,
	}
}
