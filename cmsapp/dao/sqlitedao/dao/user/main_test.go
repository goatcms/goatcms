package dao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.User {
	var (
		firstname string = "X6PejhrKtMYKBt2oigvcnAaoKhj5H9QoCFHTacSln0if"
		password  string = "0sPxWuYwtOTWh1n9F51lufCJXc8PIyJDGawqjhrdCloY"
		email     string = "Chloe.Thomas@iev.fi"
		lastname  string = "5JZo9C7qkhbYtBXyPGwkzwM2Y7CSDcJqPt9490Gqi025"
		roles     string = "L77VqP13nuvQFlMf7eEqDdwZv7vgLW1g7vFsaVYenarO"
		username  string = "cLH3HdVeFDvk1PDaXEXVknNgpURjTpAGQY7qMdzENDrq"
	)
	return &entities.User{
		Firstname: &firstname,
		Password:  &password,
		Email:     &email,
		Lastname:  &lastname,
		Roles:     &roles,
		Username:  &username,
	}
}

func NewMockEntity2() *entities.User {
	var (
		password  string = "QjsZ2jKNDQ3hkV2LfZVGkkMXmzrLpTrNBJFhw6okg2CU"
		email     string = "Ruby.Taylor@xdv.cn"
		lastname  string = "MTsBfcTVDuqDky0fhv5qjIZoG6yOEoQb3vCiZJTvXxca"
		roles     string = "JeNSVlsPZEfFVgAuACd1bYROlpWmfPg9h41pu3BK0ZtI"
		firstname string = "DOaN3V0gemtoFeOcioBJAIjIAdplsxXcpqww7E0PVNTh"
		username  string = "X0eKFTURdx2Y6Z75XbW28ZPqzl0xkYWirIxTD4XACaDu"
	)
	return &entities.User{
		Password:  &password,
		Email:     &email,
		Lastname:  &lastname,
		Roles:     &roles,
		Firstname: &firstname,
		Username:  &username,
	}
}
