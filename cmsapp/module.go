package cmsapp

import (
	"github.com/goatcms/goatcms/cmsapp/dao/sqlitedao"
	"github.com/goatcms/goatcore/app"
)

type CMSAppModule struct{}

// NewCMSAppModule create new module instance
func NewModule() app.Module {
	return &CMSAppModule{}
}

// RegisterDependencies is init callback to register module dependencies
func (m *CMSAppModule) RegisterDependencies(a app.App) error {
	dp := a.DependencyProvider()
	if err := RegisterFilesystems(a); err != nil {
		return err
	}
	if err := RegisterCommands(a); err != nil {
		return err
	}
	if err := RegisterServices(a); err != nil {
		return err
	}
	if err := sqlitedao.RegisterDependencies(dp); err != nil {
		return err
	}
	return nil
}

// InitDependencies is init callback to inject dependencies inside modules
func (m *CMSAppModule) InitDependencies(a app.App) error {
	if err := InitServices(a); err != nil {
		return err
	}
	if err := InitControllers(a); err != nil {
		return err
	}
	return nil
}

// Run is run event callback
func (m *CMSAppModule) Run() error {
	return nil
}
