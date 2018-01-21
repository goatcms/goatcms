package templatex

import (
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/dependency"
)

// Module group template functions
type Module struct {
	deps struct {
		Translate services.Translate `dependency:"TranslateService"`
		Template  services.Template  `dependency:"TemplateService"`
		Logger    services.Logger    `dependency:"LoggerService"`
	}
}

// NewModule create new template helpers module instance
func NewModule(di dependency.Injector) (*Module, error) {
	instane := &Module{}
	if err := di.InjectTo(&instane.deps); err != nil {
		return nil, err
	}
	return instane, nil
}

// Register template functions
func (module *Module) Register() (err error) {
	template := module.deps.Template
	template.AddFunc("lengthLimit", module.LengthLimit)
	template.AddFunc("messages", module.Messages)
	template.AddFunc("dict", Dict)
	template.AddFunc("contains", Contains)
	return nil
}
