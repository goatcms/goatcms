package home

import "github.com/goatcms/goatcms/services"

// MainCTRL is main page endpoint
type MainCTRL struct {
	d *Dependency
}

// NewMainCTRL create instance of a home controller
func NewMainCTRL(d *Dependency) *MainCTRL {
	return &MainCTRL{
		d: d,
	}
}

// Get is http get endpoint to serve home page
func (c *MainCTRL) Get(scope services.RequestScope) {
	err := c.d.Template.ExecuteTemplate(scope.Response(), "home/index", nil)
	if err != nil {
		scope.Error(err)
		return
	}
}
