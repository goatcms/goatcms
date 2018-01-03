package dao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.Translation {
	var (
		key   string = "jPKpCtO41NZ0A1340W0hnrnDa2J9Ln72u3Wj9ULzvHRg"
		value string = "w8lhegz3o0hzF3T4dGi6eGc77N2cvA8q614HE6OlpWYV7WkMKxtKcWd5VMNy52DDkXj4bNk3LPTVGBFh4dDnXmcjD844tsKdU0yFtt7W7UJJT0Podmyt8clsRTWqWiOX8BSlunF8 oOUaaPFx8WBZ5R6LMP39kDvtHFUH64Q9qN8PX84HBjKEt17jOZTrdFr456Yxjpz5hv2MQ44g3Gowi7I s5gVO4 UXAcPM9T9IWHUfpxGxo0KrPEQwqg7LY1QVFfrGVmu8UnBbVRDvX8rTL81jB fOCPJCraRPctcflrIJZT4ZkxKJWhcyW3E2 SC6mSe 26h0a62Ul18VVgDJnJOenHazvWU7srtaOXaXsleGH2sqq1n4tMG1yyt Gneu9LJi49R1bwJRa7"
	)
	return &entities.Translation{
		Key:   &key,
		Value: &value,
	}
}

func NewMockEntity2() *entities.Translation {
	var (
		key   string = "NsBjshvRRRusRIiT0ub7Yw3T1aYMmoFEwThYGXcum83a"
		value string = "BoIlLiCoYke41fbat9zSX7RbLFZPydkWIHWLJh2C03y8b23zKItTuLcZw5bdHoIu2IBvQfcsdIRwylH0AZtXWBZ3AUcItVpwjerU ZqsAdurSqVy3bihKZQKEkMYCj4VPpbncOUO0iN1sTyTgmMy8t1vCcNSK2Lmdd344Xo9zG5H6Qa2TKPcCi5E Y8TJjueZEzyWuyQxQ0q6YALZN4dZEZ1jCFdhLRulg3iFPj0tbqX0E2gv6sIpL7Qx UQasnrtATqychNa7ZH886kRaFF85xZUcAsIDuT1sm6aAtGxrzLMo foI1LCjTnow1V3iHyu3qUgNIHX3EQnCKRJV6OgCYDu4D495hBTbriY5AX8VMCXWm 8h Em5atIgiOSSTPuc0B4EBgFRxLLzip"
	)
	return &entities.Translation{
		Key:   &key,
		Value: &value,
	}
}
