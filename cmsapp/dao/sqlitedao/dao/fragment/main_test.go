package dao

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.Fragment {
	var (
		content string = "sIdfdxAL5kK ExlSdexhsjHGZXhs44miCFpaneStpv2o5z4jaUmOpUrkth9aGSz9XwWdtQvM0qfekrykqHWJTWPXzKQ2Qn4rQwYYnmFCyj71IvmAn38UxCJun6bL0a WUfAPlfo0bxoPJCNJt19EKxXQOeGKD1EfKzFlnC9DkHZnWStJ07wOOhNa8D0mkl3feYtpS7zAGkswjkCmcsbZZliLXytC BmL8rGbploZXP1bn74LKjIcf22SNDjyjma8PBBSDgRRY4y6JFlprlpVUHivzMp7zxwytydoXwHBNB6NaeWD8PV2bBjP7IznXrOmacbs7trdarlCfWFjZYsiZxJX4pXUmN2ugaPfBt1lM7IanhAW8zyVmND9GsbI0osBEeAOxXK3KBA1qkSE"
		lang    string = "us"
		name    string = "moyYVRXHatGbsh7zHPwY73ed9cUVIpfGL73wDChPhax7"
	)
	return &entities.Fragment{
		Content: &content,
		Lang:    &lang,
		Name:    &name,
	}
}

func NewMockEntity2() *entities.Fragment {
	var (
		name    string = "ZBkK3FL2xxtZSw4uTNA1wLGbn5CW1OrYL4X26K9lj1hT"
		content string = "7z bqnYCLbyTqYvDxdrwo5MbzYo55edZE7IwB34LP0KHgGGnbq yh5 nAdiJd4u3WqEmunqblhDGjnkIbJGTg6i9RexPn03Mf9xtvdjp6b2EKuMxwSQdsYEiJxQ FcqGQuNXDJxEDdx9ZvTla9j9uUcXfnO8UylHVllm tIS2ZscHTTrkT4 rK6tXqD0I7dtOaViCjJD8O9yg1dK4lSRQY2i0JC6lFzi7mok74o3KcQsOKvOXOypX51wVcwoCVIcI6sBVffij7VqpZi9W9R xgiuIv1jEq 7QOA9I PDdKoWhnDiNsOS7bltF8xrwA4cQkiRSlUI2ZegIwOon36zU6I2p9fCqNRnmiGoBJCuoshzYmExdH1IbrUQXzAhSZPpAmulmy4Rb9WMEN0j"
		lang    string = "pl"
	)
	return &entities.Fragment{
		Name:    &name,
		Content: &content,
		Lang:    &lang,
	}
}
