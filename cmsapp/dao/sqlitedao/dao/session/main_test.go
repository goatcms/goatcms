package dao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.Session {
	var (
		secret string = "wEnI5zKMCoKBTSJ9BoV0JXzcI6CQyGTPLhoW5KbF7uwl"
		user   int64  = 1
	)
	return &entities.Session{
		Secret: &secret,
		UserID: &user,
	}
}

func NewMockEntity2() *entities.Session {
	var (
		secret string = "rqjXmpWGpegmLEWZzu9JCC5BzlwnX40Gr3p7sSVIeEV5"
		user   int64  = 1
	)
	return &entities.Session{
		Secret: &secret,
		UserID: &user,
	}
}
