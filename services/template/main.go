package template

import (
	gotemplate "html/template"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/goatcms/goatcms/services/template/funcs"
)

// TemplateService is global template provider
type TemplateService struct {
	tmpl *gotemplate.Template
}

// NewTemplate create a template service instance
func NewTemplate() (*TemplateService, error) {
	t := &TemplateService{}
	if err := t.Init("templates"); err != nil {
		return nil, err
	}
	return t, nil
}

// Init initialize template instance
func (t *TemplateService) Init(path string) error {
	t.tmpl = gotemplate.New("template")
	funcs.Init(t)
	return filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			if _, err := t.tmpl.ParseFiles(path); err != nil {
				return err
			}
		}
		return nil
	})
}

// Funcs adds the elements of the argument map to the template's function map.
func (t *TemplateService) Funcs(funcMap gotemplate.FuncMap) error {
	t.tmpl.Funcs(funcMap)
	return nil
}

// ExecuteTemplate execute template with given data and send result to io.Writer
func (t *TemplateService) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	t.tmpl.ExecuteTemplate(wr, name, data)
	return nil
}
