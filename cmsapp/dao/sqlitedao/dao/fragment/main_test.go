package dao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.Fragment {
	var (
		content string = "XDuFNzA1iMw8MTPLMHM9gMdQSG6pbG1HJghTuHrnFb6rgFuk5yBH1K9CgVCFYBCbded6FxFGJHX6idXW1xMpaV7iAVZnai0uI3hiYvNergdPf mwcgoZF3RxznR7zsX1LWNDG0PaxCQmy0A8HMfLF1lisxVFA9Z3nSyGLMglp9qYd4tmKDccGcmuKkClciSFkPakb9CO9WsYWwhcTQla1u106XKLyiwRpXN06 cH7JVc0kWzMVFygnclJZcXe9dfobTO9aQycJrM6Tb7VENXdZQZ G RIW1K0GOw9nuDnizC6duvCojiBgcv3QVipBd3xC6bV1P3RZWaBblJfTpbfr5M xyE4WLDyxUtyw noktCE5HzXOQHsF6S2jBMaMuhCt1Ktpnr9eToT5 M"
		lang    string = "us"
		name    string = "Gz8pQpECfYFMqqphfM5FCmMhvcmtobICtGMBkp5ucU3M"
	)
	return &entities.Fragment{
		Content: &content,
		Lang:    &lang,
		Name:    &name,
	}
}

func NewMockEntity2() *entities.Fragment {
	var (
		lang    string = "fi"
		name    string = "EWcnpI4N3IfkAe1deYZKPQFKO7iTRl55FddOonjWXumr"
		content string = "XB8nTtDUvnu5RmskMB eKcaOOYWuMMZzY9aTeTptUcKgN6WBZwQTTxMTdn1cW4M5ovCOIxGXttuxvbzk6JEvdbfNYSAjMLP5JoYjLg7t1AfKEXZWH2o0FoZ76zpydteCZddm7Csa5izZyzYBHgoet4e0M8adLcjvMyMuyx9pQR8DZLnVguzsup0lt5jjsPhSU0lV4yefRcnLNzDKhK  TCItcI7hbQFSBGDz5WOkeBEJudsHpwafb6tXo1JQe C9k6PfR4HpRbpoVyLm8KCzFQoqt9zD09Qy839dQLbx3qJurRoVJalRKtyE4t82SZ0B8bhp p3wYgtDH3Vz4iALd1ZmpDvwl8d7QabPewiy1ebgdc3iX9Lrf6LiwZ74JmkE9zyj7rRMRaIkWmfW"
	)
	return &entities.Fragment{
		Lang:    &lang,
		Name:    &name,
		Content: &content,
	}
}
