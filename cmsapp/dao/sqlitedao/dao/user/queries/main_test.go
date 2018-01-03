package queries

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.User {
	var (
		password  string = "VBxyijcJIrV0bj8Qd9sOL25bwjSXfDOimyv4RDagEKn7"
		login     string = "BrjErwkNOZHrM4SofHs4QvWgKSUjyONrLNeYzFu1xLcE"
		email     string = "q92k2K3V25BJr7Lf5H4Gj8uwDpxixNf6wuZqEZXE98bu"
		firstname string = "yDisyViH7rsE804RenDuw4MRGoezVm7sLaIMHOuN6onY"
	)
	return &entities.User{
		Password:  &password,
		Login:     &login,
		Email:     &email,
		Firstname: &firstname,
	}
}

func NewMockEntity2() *entities.User {
	var (
		email     string = "BY4yVxG6JsosFKamz7yq90ovxbDD4AI95Srk4sJjyxX4"
		password  string = "NFaVOsVGe40Q2UNIJc3QcuKpCmywLmlmDEwomEwyCidQ"
		firstname string = "6E4Gd9QOJPFgJlpIl8iW69HQ4MxQww27YSFuhBqI7Mbx"
		login     string = "ChIxphakvBkVmwRKymMZgenKdzBqhMw3CdpEqFi8nWVH"
	)
	return &entities.User{
		Email:     &email,
		Password:  &password,
		Firstname: &firstname,
		Login:     &login,
	}
}
