package templatex

import "github.com/goatcms/goatcore/app"

// InitDependencies is init callback to inject dependencies inside module
func InitDependencies(a app.App) error {
	module, err := NewModule(a.DependencyProvider())
	if err != nil {
		return err
	}
	return module.Register()
}
