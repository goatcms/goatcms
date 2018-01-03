package queries

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.Translation {
	var (
		key   string = "qXbsVjhoQg0fcNSo3q6bHb9zvhqHl9kS9lFHQW2MoKKs"
		value string = "GWTIzwW1ANgGBNtmEafEzhUjck2qClDvj4eiTSkqlfDctTvn4jvquhuXjWuj381vZg5vgSjZAw3ZzdvCQDjHSrrvZaOU  0zXZmyRGYpEE204Xf5Mfc1YEJCExRikdEpksc3evnpswzLppLF3UdY4cRSW3aXwqDERhbBihr63uDtqmKoZoEI9yxjoC6h5Qxf6LSG3 iywTboOvQF6ZKfH9DnnqSfV8q7OnGVdyA43Jo86PZrzZmkNuUFowX93m7Yoen42H3eV4dXhRXZ 0uXBvu FRny5qevs5qluyAtmKJj4KK 2hvSFzsU1HA Va8xhDtWE4CZNWbQ1bM pL65CPl0GcmhbjJ5B7V9RNCRijbR7VzDzUkBEraNulM72BkHs79CYt2hYOX5xErc"
	)
	return &entities.Translation{
		Key:   &key,
		Value: &value,
	}
}

func NewMockEntity2() *entities.Translation {
	var (
		value string = "JthrhYyIDZdelYcEaShuh53CsfPxtb942ECeWqbYbaXb77u5GV86Jo42FQSgPBwJYmVxWEPlYQgJez3CjLAfQScO0JCpDvdAZ59SnPiXE51kWhYxi SHaU3oX7ifYcQZ0cX7liL4sQ0HSTSXUJ8Or4xCGdMvEaMxe1V6mG6hQp9j81iwIMxhK8xCQhIPj14iH5Q2AfTL239TlcMQ5kblhQz4otFIX0U3jtNj5PSPB55AkQWk2u3xdj09G9rbFNO0bem1hVxIgq9yHdo8DYA6Ddz2bYyUzDmaEsgBQTNxzzGBjJMeHRNDH8cl23pXpnKXW27ze336q5dHycZ5ybcH8xxSSW5yc1ymeyPThq4PvmNTRRc7fBIbFWhaF6yEuEa351fvMjD gc25S8rs"
		key   string = "De2mtBVoElB48VlDvBb7fh3cEPk901veUWEAU4Max5kC"
	)
	return &entities.Translation{
		Value: &value,
		Key:   &key,
	}
}
