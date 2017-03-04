package mailer

import (
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/dependency"
)

// RegisterDependencies is init callback to register module dependencies
func RegisterDependencies(dp dependency.Provider) error {
	if err := dp.AddDefaultFactory(services.MailerService, MailerFactory); err != nil {
		return err
	}
	return nil
}
