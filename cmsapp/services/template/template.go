package template

import (
	"fmt"
	"html/template"
	"sync"

	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/goathtml"
	"github.com/goatcms/goatcore/goathtml/ghprovider"
)

// TemplateProvider is global template provider
type TemplateProvider struct {
	deps struct {
		Filespace filesystem.Filespace `filespace:"template"`
	}
	providerMutex sync.Mutex
	provider      *ghprovider.Provider
	funcs         template.FuncMap
}

// TemplateProviderFactory create new template provider
func TemplateProviderFactory(dp dependency.Provider) (interface{}, error) {
	t := &TemplateProvider{
		funcs: template.FuncMap{},
	}
	if err := dp.InjectTo(&t.deps); err != nil {
		return nil, err
	}
	return services.Template(t), nil
}

// Init initialize template instance
func (t *TemplateProvider) init() {
	t.providerMutex.Lock()
	defer t.providerMutex.Unlock()
	if t.provider != nil {
		return
	}
	t.provider = ghprovider.NewProvider(t.deps.Filespace, goathtml.LayoutPath, goathtml.ViewPath, goathtml.FileExtension, t.funcs)
}

// Funcs adds the elements of the argument map to the template's function map.
func (t *TemplateProvider) AddFunc(name string, f interface{}) error {
	if t.provider != nil {
		return fmt.Errorf("TemplateProvider.AddFunc: Add functions to template after init template provider")
	}
	if _, ok := t.funcs[name]; ok {
		return fmt.Errorf("TemplateProvider.AddFunc: Can not add function for the same name %s twice", name)
	}
	t.funcs[name] = f
	return nil
}

// Execute execute template with given data and send result to io.Writer
func (t *TemplateProvider) View(layoutName, viewName string, eventScope app.EventScope) (*template.Template, error) {
	if t.provider == nil {
		t.init()
	}
	return t.provider.View(layoutName, viewName, eventScope)
}
