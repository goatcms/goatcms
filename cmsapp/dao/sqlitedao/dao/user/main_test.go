package dao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.User {
	return &entities.User{
		Login:     "mZjXA5Q5pfbwj12mDxqTDRkRW1LEYyRvyaTMwvcNYX3s",
		Email:     "5MNhXGyuTniBZuE3HicP7odPrgVU0JACYgKAAcTWCFFb",
		Password:  "ToZPJNcveIaUW2BkIHgxsWNeeW5MmcdS9KsG7x22j4Jx",
		Firstname: "8VACs2FwIW6CiHwSXbJq5xCiUECDxnRfgI4GrlGBza45"}
}

func NewMockEntity2() *entities.User {
	return &entities.User{
		Email:     "eXPP8T9pGo6aMXyKs7E3cdv34iCPnkLzjIBjsqnmMKGa",
		Login:     "5yyoEHjXfsAu40CQHTJFlngQ9CxcBkBb1LLeRkVodTOS",
		Password:  "2LzI0lBB5GXaMCIkipIL8izz1CIC5NKy5XcfwfO2J57J",
		Firstname: "9OQQHBSlSIW9Oj7JuiUFXRBEyqSBKwCdO1ktX0Al55Xs"}
}
