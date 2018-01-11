package queries

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.Session {
	var (
		secret string = "2gDuxXwhGXrRzIxsgoUa9wZtue6sxu4WWihUrqRYqQIb"
	)
	return &entities.Session{
		Secret: &secret,
	}
}

func NewMockEntity2() *entities.Session {
	var (
		secret string = "v9u8NvqPxo2ewhscsJPSECiZwhNx4WQWJmXaHBDhM1cP"
	)
	return &entities.Session{
		Secret: &secret,
	}
}
