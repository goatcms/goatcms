package session

import (
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/services"
)

// Factory is a session manager dependency builder
func Factory(dp dependency.Provider) (dependency.Instance, error) {
	return NewSessionManager()
}

// InitDep initialize a new session manager dependency
func InitDep(prov dependency.Provider) error {
	if err := prov.AddService(services.SessionManagerID, Factory); err != nil {
		return err
	}
	return nil
}
