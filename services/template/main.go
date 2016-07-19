package template

import (
	gotemplate "html/template"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/goatcms/goatcms/services/template/funcs"
)

// Template is global template provider
type Template struct {
	tmpl *gotemplate.Template
}

// NewTemplate create a template service instance
func NewTemplate() (*Template, error) {
	t := &Template{}
	t.Init("templates")
	return t, nil
}

// Init initialize template instance
func (t *Template) Init(path string) error {
	t.tmpl = gotemplate.New("template")
	funcs.Init(t)
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			if _, err := t.tmpl.ParseFiles(path); err != nil {
				return err
			}
		}
		return nil
	})
	return nil
}

// Funcs adds the elements of the argument map to the template's function map.
func (t *Template) Funcs(funcMap gotemplate.FuncMap) error {
	t.tmpl.Funcs(funcMap)
	return nil
}

// ExecuteTemplate execute template with given data and send result to io.Writer
func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	t.tmpl.ExecuteTemplate(wr, name, data)
	return nil
}
