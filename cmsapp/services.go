package cmsapp

import (
	"github.com/goatcms/goatcms/cmsapp/services/crypt"
	"github.com/goatcms/goatcms/cmsapp/services/databases"
	"github.com/goatcms/goatcms/cmsapp/services/genservices/fixture"
	"github.com/goatcms/goatcms/cmsapp/services/logger"
	"github.com/goatcms/goatcms/cmsapp/services/mailer"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep/reqacl"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep/reqauth"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep/reqerror"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep/reqresponser"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep/reqsession"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep/reqtranslate"
	"github.com/goatcms/goatcms/cmsapp/services/router"
	"github.com/goatcms/goatcms/cmsapp/services/session"
	"github.com/goatcms/goatcms/cmsapp/services/template"
	"github.com/goatcms/goatcms/cmsapp/services/translate"
	"github.com/goatcms/goatcms/cmsapp/services/user/signup"
	"github.com/goatcms/goatcore/app"
)

func RegisterServices(a app.App) error {
	// services
	dp := a.DependencyProvider()
	if err := session.RegisterDependencies(dp); err != nil {
		return err
	}
	if err := template.RegisterDependencies(dp); err != nil {
		return err
	}
	if err := crypt.RegisterDependencies(dp); err != nil {
		return err
	}
	if err := router.RegisterDependencies(dp); err != nil {
		return err
	}
	if err := mailer.RegisterDependencies(dp); err != nil {
		return err
	}
	if err := logger.RegisterDependencies(dp); err != nil {
		return err
	}
	if err := translate.RegisterDependencies(dp); err != nil {
		return err
	}
	if err := signup.RegisterDependencies(dp); err != nil {
		return err
	}
	if err := fixture.RegisterDependencies(dp); err != nil {
		return err
	}
	if err := databases.RegisterDependencies(dp); err != nil {
		return err
	}
	return nil
}

func InitServices(a app.App) error {
	// services
	if err := reqsession.InitDependencies(a); err != nil {
		return err
	}
	if err := template.InitDependencies(a); err != nil {
		return err
	}
	if err := translate.InitDependencies(a); err != nil {
		return err
	}
	// request services
	if err := reqerror.InitDependencies(a); err != nil {
		return err
	}
	if err := reqauth.InitDependencies(a); err != nil {
		return err
	}
	if err := reqtranslate.InitDependencies(a); err != nil {
		return err
	}
	if err := reqresponser.InitDependencies(a); err != nil {
		return err
	}
	if err := reqacl.InitDependencies(a); err != nil {
		return err
	}
	return nil
}
