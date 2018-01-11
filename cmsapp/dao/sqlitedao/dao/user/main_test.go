package dao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.User {
	var (
		firstname string = "W1F53lSKlmQIEn8OTUYiESSQcDypHZX6FyzGBqC4dCJD"
		lastname  string = "vNAimHGTdqnyhtosgvu5nTXuJabbkmnoegysYwxuGCf1"
		email     string = "Jessica.Robinson@mms.pl"
		password  string = "F017wSE5K7NOr8GsS3ENyAzD3JD3K3pRmtheJTy1mpfh"
		roles     string = "GGPYEW8Y60uZZGYYdoLDq3SEzJ3GullsxvSyw1EWBCUH"
		username  string = "HGTPxFrb7n7PeHJki142vryJ69Tfw7KQLhvKERjSqhZx"
	)
	return &entities.User{
		Firstname: &firstname,
		Lastname:  &lastname,
		Email:     &email,
		Password:  &password,
		Roles:     &roles,
		Username:  &username,
	}
}

func NewMockEntity2() *entities.User {
	var (
		firstname string = "xTd6U4yCn7h9qHWoAnHYfjXukSam0EcDYYbSq1N8DIl9"
		lastname  string = "ZL3oj6UV0qCduBMKzvXZaF4si76uQx3zn8Mmj8t3xjbz"
		email     string = "Alfie.Thomas@qjk.fi"
		password  string = "Sv91RsxPc4Vjli7lxh051AYXuY5IXGgLI4usU9fZn2LB"
		roles     string = "YEVgLGH7G4sj2PKvCuR9AFSy6ZV1MakRt5RzpisD9qQX"
		username  string = "gFIBFr9GtsAchhu0TxTmjEKhbBf0B1RQNs3O80PC3jlo"
	)
	return &entities.User{
		Firstname: &firstname,
		Lastname:  &lastname,
		Email:     &email,
		Password:  &password,
		Roles:     &roles,
		Username:  &username,
	}
}
