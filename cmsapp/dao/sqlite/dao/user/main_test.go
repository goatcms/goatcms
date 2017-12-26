package userdao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.User {
	return &entities.User{
		Lastname:  "joOMF0e5xMP7xwBh723AT8gDIGyHI6OBrq0y4Cid2sjs",
		Firstname: "gMTKB2IJkXNNGgz7WF0uw8IqnxkhCQQ3MBxu3hXTtIDI"}
}

func NewMockEntity2() *entities.User {
	return &entities.User{
		Lastname:  "ziuIimEvamYacEGB9r6QLQclqj74jjv7Vr7Y69aew7u3",
		Firstname: "RCQA43mhTy569mnQqvsfgwPZ3VkZLihKJgLgfMVmjqEM"}
}
