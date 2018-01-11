package dao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.Fragment {
	var (
		lang    string = "us"
		name    string = "H36Idjnja9GmTWZjRiftAsjPsUOaOOKX0qDIepVrNbZk"
		content string = "UzZGuHT6ZMREf0zxLRKyPxBcydjz7qn4Y7xyp7olFu4h4v0lXcvHXXDyrjL94tthRSUgZJJDN7izz0UdFLzJbhQ6bm4cVkSncGmkoexGn81DlvFY8NrdFC 1p0iDQ2lJghJzbOCsotD1iw2GXwUQLAxg87XCCQOgH3YyjcBF2IMslmoE4XvVRZwFeulfiY15eqUqvgzQxmXexxFWlXh7p7GceIQpOQUjnecPwu1dWIlIf 2j7iiTk4p7XoXgve3oQkh5kqkDjyw8Z h1JBBJmgOKbWSh4ZQ Fry71SddayhsJSUxEd2MdgFOXOIhPZG9 NNR2ktQiQ2NHmX3SRvotMVfUz8OebMqj6dwxMSAdaruvh0Bm9g nPW5SBSspxxmwxWWV8wMvlzIjmxQ"
	)
	return &entities.Fragment{
		Lang:    &lang,
		Name:    &name,
		Content: &content,
	}
}

func NewMockEntity2() *entities.Fragment {
	var (
		lang    string = "pl"
		name    string = "FdBB1FvHXiZPw0Xk7jCKbPcsa7KXPl0uJp3KoADaFaMO"
		content string = "O2Z2Dwdc1V3uegwtx8e9LBswjceo7tVUO5SsH8F6bYmviKuMWKapJ15xFbmJj2vNHPgwGhzlP3X41eywhuyc34Di3lj3Kxx85UVNFCRe8ZmYOAxo5GI5MTHIVWudBzCFqjzp hC2YE3TPjnEsNLGRbMVXH9gwKPYsD3wHTHSBa7CDgqDUVcUvdrpZ72mQaikrqm7tWRN39FZZHwI1sR3CwSjJ V8pf2AcqxD4urdT2JYBnyJBCYzAOvzcEr4h6wAZLFbCu2ZCT9c8XIZqAFyxw5pTFTuiJhU7sJuRSjFYZL8xKvJJF F5JysngNUS9Qjls4SpYoEhgc1QqM5bb24KGNJMMUZkGKZM9MAx8vp46pfyWvYRlSuU8I6tjg6DQjoTYHA cbNCpnPQOpe"
	)
	return &entities.Fragment{
		Lang:    &lang,
		Name:    &name,
		Content: &content,
	}
}
