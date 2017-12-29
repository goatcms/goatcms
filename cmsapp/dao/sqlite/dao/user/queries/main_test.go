package queries

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.User {
	return &entities.User{
		Login:     "LCXXGmNCk4IelH8LSgMIoEsHeAasuDRaKeBMnC8dHFGO",
		Email:     "YR5JH6ZpSILWMHrdug5HlM20jjdQhASWZgE66yxoI09B",
		Password:  "IUgXern6VJTPImj4EC9SRdVhhYwudcbZx9YsJXsMKzaR",
		Firstname: "lq2qvxF8oTrhzMGSWdPRZvVF1yRnSoGDtb57rjURtu66"}
}

func NewMockEntity2() *entities.User {
	return &entities.User{
		Firstname: "YwQB4DH9TU2UYUgCffshZflUmyfgfpE7dOPx61GkBOrn",
		Login:     "OBCPyVW0kcIC6Ct3gGHGEv94r4mMO90tE5JpKrxj4NO9",
		Email:     "GVOer7hfguigvSYA5NI8RkMmdprPlUEFjX1c1IgeRjoo",
		Password:  "f93uJ06oXSB9bl3wRGfyc2xRyEN0QFh5sE4Mc3fBhtA2"}
}
