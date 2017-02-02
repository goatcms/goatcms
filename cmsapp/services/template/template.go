package template

import (
	"fmt"
	gotemplate "html/template"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goatcms/cmsapp/services"
)

// TemplateProvider is global template provider
type TemplateProvider struct {
	deps struct {
		Path string `config:"?template.path"`
	}
	tmpl    *gotemplate.Template
	funcMap gotemplate.FuncMap
	inited  bool
}

// TemplateProviderFactory create new template provider
func TemplateProviderFactory(dp dependency.Provider) (interface{}, error) {
	t := &TemplateProvider{
		funcMap: gotemplate.FuncMap{},
		inited:  false,
	}
	if err := dp.InjectTo(&t.deps); err != nil {
		return nil, err
	}
	if t.deps.Path == "" {
		t.deps.Path = services.DefaultTemplatePath
	}
	return services.Template(t), nil
}

// Init initialize template instance
func (t *TemplateProvider) Init() error {
	if t.inited == true {
		return fmt.Errorf("template service can be inited only once")
	}
	t.tmpl = gotemplate.New("template")
	t.tmpl.Funcs(t.funcMap)
	if err := filepath.Walk(t.deps.Path, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			if _, err := t.tmpl.ParseFiles(path); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}
	t.inited = true
	return nil
}

// Funcs adds the elements of the argument map to the template's function map.
func (t *TemplateProvider) AddFunc(name string, f interface{}) error {
	if t.inited == true {
		return fmt.Errorf("TemplateProvider.AddFunc: Add functions to template after init is illegal")
	}
	if _, ok := t.funcMap[name]; ok {
		return fmt.Errorf("TemplateProvider.AddFunc: Can not add function for name %s twice", name)
	}
	t.funcMap[name] = f
	return nil
}

// ExecuteTemplate execute template with given data and send result to io.Writer
func (t *TemplateProvider) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	if !t.inited {
		if err := t.Init(); err != nil {
			return err
		}
	}
	return t.tmpl.ExecuteTemplate(wr, name, data)
}
