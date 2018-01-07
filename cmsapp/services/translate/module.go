package translate

import (
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

// RegisterDependencies is init callback to register module dependencies
func RegisterDependencies(dp dependency.Provider) error {
	if err := dp.AddDefaultFactory(services.TranslateService, TranslateFactory); err != nil {
		return err
	}
	return nil
}

// InitDependencies is init callback to inject dependencies inside module
func InitDependencies(a app.App) error {
	var deps struct {
		Template  services.Template  `dependency:"TemplateService"`
		Translate services.Translate `dependency:"TranslateService"`
	}
	if err := a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	if err := deps.Template.AddFunc("Translate", func(key string, values ...interface{}) string {
		t, err := deps.Translate.Translate(key, values...)
		if err != nil {
			return key
		}
		return t
	}); err != nil {
		return err
	}
	return deps.Template.AddFunc("TranslateFor", func(key, prefix string, values ...interface{}) string {
		t, err := deps.Translate.TranslateFor(key, prefix, values...)
		if err != nil {
			return prefix + key
		}
		return t
	})
}
