package home

import "github.com/goatcms/goatcms/services"

// HomeCTRL is main page endpoint
type HomeCTRL struct {
	d *Dependency
}

// NewHomeCTRL create instance of a home controller
func NewHomeCTRL(d *Dependency) *HomeCTRL {
	return &HomeCTRL{
		d: d,
	}
}

// Get is http get endpoint to serve home page
func (c *HomeCTRL) Get(scope services.RequestScope) {
	err := c.d.Template.ExecuteTemplate(scope.Response(), "home/index", nil)
	if err != nil {
		scope.Error(err)
		return
	}
}
