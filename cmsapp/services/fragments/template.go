package fragments

import (
	"html/template"
	"strconv"
	"strings"

	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/dependency"
)

// TemplateHelper is fragment helper service
type TemplateHelper struct {
	deps struct {
		FragmentStorage services.FragmentStorage `dependency:"FragmentStorage"`
	}
}

// TemplateHelperFactory create new TemplateHelper instance
func TemplateHelperFactory(dp dependency.Provider) (in interface{}, err error) {
	helper := &TemplateHelper{}
	if err = dp.InjectTo(&helper.deps); err != nil {
		return nil, err
	}
	return services.FragmentTemplateHelper(helper), nil
}

// RenderFragment return a HTML content for fragment. It is uset for small block with inline editor
func (helper *TemplateHelper) RenderFragment(key, defaultValue string) (result template.HTML) {
	var row *services.Fragment
	if row = helper.deps.FragmentStorage.Get(key); row == nil {
		row = &services.Fragment{
			ID:   0,
			HTML: defaultValue,
		}
	}
	return template.HTML(strings.Join([]string{
		`<div class="fragment" g-small-fragment g-fragment-key="`,
		key,
		`" g-fragment-id="`,
		strconv.FormatInt(row.ID, 10),
		`">`,
		row.HTML,
		`</div>`,
	}, ""))
}
