package queries

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.Fragment {
	var (
		content string = "YHL4t9WMrezQHScuEUfxgwH aXqZndc14iDzcqPKTSxGsT 7mXv9gRVUma805rkt6yATpcWLKJvqDdVHlDNJZhELxxb0xViuTkvtcjnpuDZ1pKeUfYVpXK4Q472mjhcG8XbbATcfvogNGsQJgU9XJ01p7neLNuAVv3TKDXhhLNtVmKzWrkEUO8o8sD2eaJUCn2v0FgmMSdReIcLz9Gv7DSWdgo5cD2KnpnHFPJ9jRLW6laGYwM8v8QQtGYts4NI9cbgOszudoOolPWTVIeL4UZQKkfyMUTRq8a2klB6UuCbzyewfIkDf709FlpgZOjw1LaNZ Agtqaitl6QBI7l1V9SByXQd2hpbSQ URtDCY7oiEsYjo V0A19zQHklWxdDIAOS qD0N4WbBdex"
		lang    string = "cn"
		name    string = "BeoJ2QVjvaj3sjIpaTtCXzW5XP4zwgYslZlh9rsS9xPG"
	)
	return &entities.Fragment{
		Content: &content,
		Lang:    &lang,
		Name:    &name,
	}
}

func NewMockEntity2() *entities.Fragment {
	var (
		lang    string = "fr"
		content string = "MQOQ1EpSFaSjhRD4MTTCBru8SaJ47hzgZHrxMJNUSShknbux 07lEOFBM65EOaY2DqNypMYENTGigmg Pn4szFDc1uememv4eSfexjhdcaKO jH3H81r6Nun7Wl85xWkVeRsrSTkWBpgU5wczjmAtaeELEMg5IE4O LUt1aaKc273B7hFJ97JUdnK2caFVBUyxQAKqziTCMPDaJrAnd56iUHBzjxwYqxtH6YNJopgjkKyDTKdkj9Vq6hTA5v9khruB AHLYD6rrtexMEiTTTiMpWsvD2047wEr4w3ZbuiBEDiB3xhxb4W2PsfhsUp7GEZDeExMsdvk3ssapfQG7SSSQ6xNFJRqr8LR1Bxgu6k1FmJrHTJwwUPh9owTfvv1gy2tWlnCosLmLgaujc"
		name    string = "5oqsbYFoyw6yU9Jjc1WyE9yqLm51ClsPCyGhU06jrTvP"
	)
	return &entities.Fragment{
		Lang:    &lang,
		Content: &content,
		Name:    &name,
	}
}
