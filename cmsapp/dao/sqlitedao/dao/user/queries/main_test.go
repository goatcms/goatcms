package queries

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.User {
	var (
		email     string = "G8VBQoI2ijurOTQJJGpZzOgtKlk4McqhIzGsROrDqE5O"
		firstname string = "P0KFeogJm4FLQlB0T8ilodd0rmDW2RIU5r5yWVuP0RA4"
		password  string = "aZotKCGNSEDBb2q6kRvybeBANFQZX1iOufcPefylr8k9"
		login     string = "0e9iSk9PbFPBxADFki9kzmjiTkd9OKikaECb4I7ZuC7Q"
	)
	return &entities.User{
		Email:     &email,
		Firstname: &firstname,
		Password:  &password,
		Login:     &login,
	}
}

func NewMockEntity2() *entities.User {
	var (
		email     string = "fs6uGSH7YSWIIaK89nO8G4d5covlkuP8pmFVJGwPM112"
		firstname string = "k1NcwMTESh1hhSQKLFQsab5psPSBmUwm3YWUfdXe0rqY"
		password  string = "Rr20H0hpojWtX83OKJPurMSTmapNkXIImuJ4IpCUktzm"
		login     string = "VIU00E01mwSeutzkt0qw81TQ0LP2yrqrgGDEXPka3YeO"
	)
	return &entities.User{
		Email:     &email,
		Firstname: &firstname,
		Password:  &password,
		Login:     &login,
	}
}
