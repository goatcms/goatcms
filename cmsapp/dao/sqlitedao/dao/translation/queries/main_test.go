package queries

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.Translation {
	var (
		key   string = "voJJ7TeCjaJ2aZZHqc2FRcqO3Z7BS4f3ZkHMRdXL0U6Y"
		value string = "d4ae1wGamu4p8w0o RbCbP8AewVttN9C2KOlzwWXjhDTTtE51MbkFut2AES0bJCkMZ4WwX0pEzpueEZ2QVHwT c1uAcfoXG579NpNDZDUvbHYxiy D5i XkQcjSTOMgx6FEEV2DI4V4tOB3SRMm20Ycl7ZAu3q0aDLaff8rIPqtKZNMVjlE4O qiJQNPxgTFTah0viTIC GiDF9yNxeWdl5JTHoqYzHu4ytar94pA0bhfJ2VNIN8Y2pmCd1XOjzSAJwSTzul68rbBJKWMQgIsjDTtERKPIUz3SML9Orc51drCbkzFTb7n5JcnjMaqSyhUfMKUNEHulDZb1umypHRwg4uv541pNkk0p6DipH2wcFlzoGFNUa5rMCrdVKJpNlWERmpr9mCBif4geHL"
	)
	return &entities.Translation{
		Key:   &key,
		Value: &value,
	}
}

func NewMockEntity2() *entities.Translation {
	var (
		key   string = "SJ2mEBrmWKSMMNb7qyTxyPIxitKAeDZWENZKMjOwyVAA"
		value string = "nC6lIPv4sqV8A5L6WWF JByb081mDYoTU239nIpZUzN4wFvEmgTHrhnQaeu8NfCFmiYbe3eEUJ57Yb5rJAWWAKHDVbQCTqGVkPqbAv3pdRRQyUUqdNQWAVnBF8kwpBMPjhYTtE3Cq3XJMoCRzlkgI6PwyOKFBNurupoVGzHfjYInRqm9G C jOKLAkMxB0TNT5l4LSOSGLNaMHovOHd1VSp3HK7T6sCshewiY4tGAakrDAGLan1tLce6GzYk 4KpfwYVjiUqXZW1HgXmAg5ybQMWL2 3JkRb eoizyBYNnALZoalcem8 MfNxIjN2Qs2xItJ7EVWU0mIfqeTgUPlJf SvZ8FaFsWizE8t2xjlPL7JHbj88793SntGHFc14TCIzXb4ipctT9gLY0h"
	)
	return &entities.Translation{
		Key:   &key,
		Value: &value,
	}
}
