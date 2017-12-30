package dao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.User {
	return &entities.User{
		Firstname: "vQ61BF5B5cfnzBkoq7tCBrqasgIr50YgMKIJglVuwqi3",
		Password:  "PXegHK8CP0aJhOD5sCfzbk3dgkn9MrHg9sJQ8PYC0aaO",
		Login:     "pTbj0mLuRshUZXA2K7ViQ18e8L5YgmFJKbwBT7pxAr8b",
		Email:     "i0JPUNLokXsOijULZpIHEk3xakQAKQUgdJVabOePXAPN"}
}

func NewMockEntity2() *entities.User {
	return &entities.User{
		Email:     "zBEFGsOv9PqgKfTniaRUV4t7oAlv8JKRKlYgtF2li91i",
		Firstname: "cPzUWgSBKoGd6aCyLMvkHbDJI8C4vvlLzAzmkd1wyD57",
		Password:  "ek7VHP4jZpRCXo06UALk630QeyTrHgKtUhJYErtoQzGX",
		Login:     "nnXWhGgLYcc546fyem11OM2o3fb5v78t7J6dObH4vrf2"}
}
