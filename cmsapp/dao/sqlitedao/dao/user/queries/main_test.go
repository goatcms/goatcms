package queries

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.User {
	var (
		login     string = "nLZrJsjBSIQpzTdoue3r5ozRYcnIHlOzVTJ3sH1biq1M"
		firstname string = "6uXLZo7mW4vAkXeahkM941c6QHJ2QyUGP0r3OsrZ6JSm"
		lastname  string = "Fztnri0TEp7RbCqbLUhfYDY7spZ7CmfHSBwLuSOJrx0E"
		email     string = "William.Evans@ywi.en"
		roles     string = "5R2SbaR3oCvnPKqtWYRn5EwfV3Z4VRqLCYF6LlhtYBL7"
		password  string = "uiklZOZEY07jpGJScxMEvwdhbpTiRzjRi06a6eauPPVt"
	)
	return &entities.User{
		Login:     &login,
		Firstname: &firstname,
		Lastname:  &lastname,
		Email:     &email,
		Roles:     &roles,
		Password:  &password,
	}
}

func NewMockEntity2() *entities.User {
	var (
		lastname  string = "QG25qSthy6l6gbsDbH1YxMx6Ol95PBr1PGjtVgTIWJ04"
		password  string = "HTpNu4MfiUoSeLxEatnsyy1jCPU40Sa1mqssafytZu4D"
		login     string = "x8YbPkMMW9XcFz5IX46jpI3s3kjG4rjeAhjGr0KiCXXA"
		email     string = "Ella.Brown@swk.fi"
		firstname string = "z4QblWLp36eMSqzE61mGiI6lvHZCMD2ZvlCtMfQ3znnN"
		roles     string = "XHkh1oG9Nzd5WQVYXECyJRKykMM09vmeXZv2V770TDJJ"
	)
	return &entities.User{
		Lastname:  &lastname,
		Password:  &password,
		Login:     &login,
		Email:     &email,
		Firstname: &firstname,
		Roles:     &roles,
	}
}
