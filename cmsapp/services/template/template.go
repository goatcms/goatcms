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

// Provider is global template provider
type Provider struct {
	deps struct {
		Filespace filesystem.Filespace `filespace:"template"`
		Cached    string               `config:"template.cached"`
	}
	providerMutex sync.Mutex
	provider      *ghprovider.Provider
	funcs         template.FuncMap
	isCached      bool
}

// ProviderFactory create new template provider
func ProviderFactory(dp dependency.Provider) (i interface{}, err error) {
	provider := &Provider{
		funcs: template.FuncMap{},
	}
	if err = dp.InjectTo(&provider.deps); err != nil {
		return nil, err
	}
	if provider.deps.Cached != "false" {
		provider.isCached = true
	}
	return services.Template(provider), nil
}

// Init initialize template instance
func (t *Provider) init() {
	t.providerMutex.Lock()
	defer t.providerMutex.Unlock()
	if t.provider != nil {
		return
	}
	t.provider = ghprovider.NewProvider(t.deps.Filespace, goathtml.HelpersPath, goathtml.LayoutPath, goathtml.ViewPath, goathtml.FileExtension, t.funcs)
}

// AddFunc adds the elements of the argument map to the template's function map.
func (t *Provider) AddFunc(name string, f interface{}) error {
	if t.provider != nil {
		return fmt.Errorf("Provider.AddFunc: Add functions to template after init template provider")
	}
	if _, ok := t.funcs[name]; ok {
		return fmt.Errorf("Provider.AddFunc: Can not add function for the same name %s twice", name)
	}
	t.funcs[name] = f
	return nil
}

// View execute template with given data and send result to io.Writer
func (t *Provider) View(layoutName, viewName string, eventScope app.EventScope) (*template.Template, error) {
	if t.provider == nil {
		t.init()
	}
	return t.provider.View(layoutName, viewName, eventScope)
}
