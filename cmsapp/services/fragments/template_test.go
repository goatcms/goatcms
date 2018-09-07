package fragments

import (
	"html/template"
	"strconv"
	"strings"
	"testing"

	"github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/scope"
)

func TestTemplateStory(t *testing.T) {
	var (
		mapp app.App
		err  error
		deps struct {
			FragmentInserter dao.FragmentInsert              `dependency:"FragmentInsert"`
			TemplateHelper   services.FragmentTemplateHelper `dependency:"FragmentTemplateHelper"`
		}
		entity        *entities.Fragment
		ExpectedText  = "some content"
		EntityContent = "##" + ExpectedText + "\n"
		ExpectedLang  = "pl"
		ExpectedName  = "first_fragment"
		result        template.HTML
	)
	if mapp, err = NewTestApp(); err != nil {
		t.Error(err)
		return
	}
	if err = mapp.DependencyProvider().InjectTo(&deps); err != nil {
		t.Error(err)
		return
	}
	entity = &entities.Fragment{
		Content: &EntityContent,
		Lang:    &ExpectedLang,
		Name:    &ExpectedName,
	}
	s := scope.NewScope("test")
	if _, err = deps.FragmentInserter.Insert(s, entity); err != nil {
		t.Error(err)
		return
	}
	if err = s.Trigger(app.CommitEvent, nil); err != nil {
		t.Error(err)
		return
	}
	key := ExpectedLang + "." + ExpectedName
	result = deps.TemplateHelper.RenderFragment(key, "")
	// Content can by wrap by HTML elements
	if !strings.Contains(string(result), ExpectedText) {
		t.Errorf("Result don't contains '%s'", ExpectedText)
		return
	}
	idSTR := " g-fragment-id=\"" + strconv.FormatInt(*entity.ID, 10) + "\""
	if !strings.Contains(string(result), idSTR) {
		t.Errorf("Result must contains g-fragment-id attribute")
		return
	}
	keySTR := " g-fragment-key=\"" + key + "\""
	if !strings.Contains(string(result), keySTR) {
		t.Errorf("Result must contains g-fragment-key attribute")
		return
	}

}
