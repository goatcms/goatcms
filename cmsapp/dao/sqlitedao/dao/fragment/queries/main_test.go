package queries

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.Fragment {
	var (
		name    string = "mBzCQBWHBEJU0sPvPjVa2RzlhbKiucYnUFBK0MxQOtBz"
		content string = "XPOQsLJt5noKIhoSDmojPI66y5iJBC4YgfOpRbKzpNwNvHP3Iot3s6kbnqL3SGZeNenF 2PSM0OFnd8juGr6X5PjB6Gmt8yQvApGhc0OiB E wUmztwklSy9sq9DXPASng yXSGij1ypim1E47RHfQ9zGitgc76bTJboNEXpmRcFXxQgzFg8vN9M8Et4vi7lrUf8GpgaC18dfbvbeGHW9wPMd8v5LAIzAuLT68XqfSZx6euQIltvlGjrviFfGr9HhAE6kMKH6VBmYlwSdR1IGWguHGbDJXkeYyHeUtwzNFW9hPAgZ I fTGeAKQmsRUqnlrRJKKh4AMUoUFTGTQVuCCY0pusYwyH75QEHsiKlHYqLnhc8JQQ47WJD4CbxO3F650U7KXg2n4XEgzS"
		lang    string = "en"
	)
	return &entities.Fragment{
		Name:    &name,
		Content: &content,
		Lang:    &lang,
	}
}

func NewMockEntity2() *entities.Fragment {
	var (
		lang    string = "fi"
		name    string = "Vgi8xqfkNw1SWErrxC4FQv9qbHY4YZeX6NjJlmEiq0mj"
		content string = "fXuHVQNhqYZxBWO0G1GGEvfSwZVJPMMOaoDdWJlT6gpiUR5KdpgdE3cFcSb4M8gVM05aqYvrTRt dxpqIeblDwaoaG6KKKFMhuOx6A5ZSrXRfDYe4kQS7Y1asey seMMSyMxALqPuSrorxtX4MX8zZ1lCKxsaWxfMvFzkWNCG2A2UARRM5PfnoKCm5ORZqQmpcdyEkztEHUWmCJ0KjFOiEDRqqyrL63wZ8dxfjGi1yQFZMXIFzNMTSNIAfwDnXhtIHTXtJzeU0LyTk6R6LHsRssrpyIg8GwmUEoQ3FKuiNwWt7Fyvl79q8x1 QBztQ6wNiT7AmsvTYkzBQNmfF8HTxnQk549 1E1hiKaTMdxNWXuhebGG ff0nZsbtgPMRapqc4t1vcGi91qSpuY"
	)
	return &entities.Fragment{
		Lang:    &lang,
		Name:    &name,
		Content: &content,
	}
}
