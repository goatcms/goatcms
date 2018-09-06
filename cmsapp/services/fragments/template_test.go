package fragments

import (
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
			FragmentInserter dao.FragmentInsert       `dependency:"FragmentInsert"`
			FragmentStorage  services.FragmentStorage `dependency:"FragmentStorage"`
		}
		entity          *entities.Fragment
		ExpectedContent = "some content"
		ExpectedLang    = "pl"
		ExpectedName    = "first_fragment"
		result          *services.Fragment
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
		Content: &ExpectedContent,
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
	/// ADD RENDER TEMPLATE HEREs
	// TODO: ...

}
