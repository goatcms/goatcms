package router

import (
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcms/cmsapp/services"
)

// RegisterDependency is init callback to register module dependencies
func RegisterDependencies(dp dependency.Provider) error {
	if err := dp.AddDefaultFactory(services.RouterService, RouterFactory); err != nil {
		return err
	}
	return nil
}
