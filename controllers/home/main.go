package home

import "github.com/goatcms/goatcms/services"

// Dependency is default set of dependency
type Dependency struct {
	DP       services.Provider
	Template services.Template
	Mux      services.Mux
}

const (
	// HomeURL is url of homepage
	HomeURL = "/"
)

// NewDependency is default set of dependency
func NewDependency(dp services.Provider) (*Dependency, error) {
	var err error
	d := &Dependency{
		DP: dp,
	}
	if d.Template, err = dp.Template(); err != nil {
		return nil, err
	}
	if d.Mux, err = dp.Mux(); err != nil {
		return nil, err
	}
	return d, nil
}

// Init initialize the homepage controller package
func Init(dp services.Provider) error {
	d, err := NewDependency(dp)
	if err != nil {
		return err
	}
	home := NewHomeCTRL(d)
	d.Mux.Get(HomeURL, home.Get)
	return nil
}
