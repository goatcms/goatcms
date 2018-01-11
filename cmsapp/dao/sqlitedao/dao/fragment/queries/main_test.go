package queries

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.Fragment {
	var (
		lang    string = "cn"
		name    string = "gZaWvBHSkfMsJHfQ88Ew4IbZDSw1EL7aAHorgY3yg647"
		content string = "e wsjyhSgfKJo7VnXYJ5ftTGQ1 yZsJH7a1dYIRsYXD8KWZU8uz8c9lUjxKtDxgHfMXioCRQs2S8aNpZBKC59Gx OtrLd0w6xXB  Bp3sN9uV90xqs6LbYRDblqcncimg M ybWhCySmfT26yHkIHG9T5wgIu0NZKAxZT128MGsIlv5ltOceE9sBrABzw7TMSlGfkphqMT0mElKLSISNg65N8i5t MkGA kbpT9en8YUh4d hTdG72zWMeEPPC80NSxJCnsxSkKG5EB MiUM8LHBcZQEgWls4kuiNjNcxaRfqCHS 0Af1ERitSKYmCbzxRoIVxHGXu48WduRgLbWTBSn057wBC6k7sbFJgkX 9DMzUnnhCaeoDXCd2U8ZlkEgf2vLpTBJ7GZvZeo"
	)
	return &entities.Fragment{
		Lang:    &lang,
		Name:    &name,
		Content: &content,
	}
}

func NewMockEntity2() *entities.Fragment {
	var (
		lang    string = "fr"
		name    string = "Pn4ZXOukpsSJcIFzwuyNnEgZk9fkevNZidovqxYKJw5G"
		content string = "UlSx2OOCFRTJ8stRYXg61OWMjeY0VHTqWm2KVIfgfSPCsNLWkGPcLYtJnLC677Ez97yq gVpupi0kmlqPJNTihqKa3KmyRlkq2s6t1e1GsOojyCw9rGoDMEbydy0PnQ9XwoO4X0Srp64CIdJUMMXljz H9t395gFDZYGqj5dFFSjG33ZDlPT6HaSJ2WJg6bYMQsS YH0FbRyIF8D2SbtyM21ce2jG4WpTyzQ0lSNX0FAPmnsc8951gPDz5O zebiyeAa1r1T5dYeBSMmXNg52y LIWqmyyyJvbYbCdyuej9K 2lzl2QjVA03N5SGcCXXQAjJNYhIKKQ1enOKhybLp87zW2LZjfDKdMxU6UMoFH5XZplKaRadNhaMqafeoWvMenPNeh8gp1N0TlYu"
	)
	return &entities.Fragment{
		Lang:    &lang,
		Name:    &name,
		Content: &content,
	}
}
