package queries

import (
	entities "github.com/goatcms/goatcms/cmsapp/entities"
)

func NewMockEntity1() *entities.User {
	var (
		firstname string = "ycxjSmPAp22VvPMIRhh3HcQymmuZVrK41Ov3WdL0Ri9x"
		lastname  string = "0v45tDtubXhCLDg29qxWxI97mBAwfzA58Icww6dvPy6I"
		email     string = "William.Evans@saq.en"
		password  string = "rPRvBfHh0t299oB4AP6qUKQtDi5IXS9d19OHpyZoVTcl"
		roles     string = "SnNlZKoyJNOKpbF0aD6Jh5AsYGzU9UQUpBbjPdxZJHLa"
		username  string = "wmpIY1L1OgUsWehPJdTRHpKXVYjvlOdhcJoqYNLRwuvj"
	)
	return &entities.User{
		Firstname: &firstname,
		Lastname:  &lastname,
		Email:     &email,
		Password:  &password,
		Roles:     &roles,
		Username:  &username,
	}
}

func NewMockEntity2() *entities.User {
	var (
		firstname string = "rffUtQCezdQU8haG01N3Q4Pp3CrRQj7Biky36y5LMTx7"
		lastname  string = "TOTYn2gg661XvG0aAtGcFYeASHUClyusA30cVoB4x9Ho"
		email     string = "Ella.Brown@gbz.fi"
		password  string = "uNGZJSMVdA383lehHkj68lg3t7CfpBwqGSsCN1gTMVms"
		roles     string = "EBXVURsEMxdRliRsYLy90Fr4Z4j3vkc1OaLHFcldZaOg"
		username  string = "x1eFEFYKKxc7gZvK1SSskkPBaImzJSSN8JcIfsmsCpNJ"
	)
	return &entities.User{
		Firstname: &firstname,
		Lastname:  &lastname,
		Email:     &email,
		Password:  &password,
		Roles:     &roles,
		Username:  &username,
	}
}
