package fragment

import (
	"html/template"
	"net/http"

	"github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcms/cmsapp/entities"
	httpfragment "github.com/goatcms/goatcms/cmsapp/http/httpmodel/fragment"

	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/goathtml"
	"github.com/goatcms/goatcore/messages"
	"github.com/goatcms/goatcore/messages/msgcollection"
)

// Insert is a controler to create new article
type Insert struct {
	deps struct {
		Template services.Template  `dependency:"TemplateService"`
		Logger   services.Logger    `dependency:"LoggerService"`
		Inserter dao.FragmentInsert `dependency:"FragmentInsert"`
	}
	view *template.Template
}

// NewInsert create instance of a insert articles controller
func NewInsert(dp dependency.Provider) (ctrl *Insert, err error) {
	ctrl = &Insert{}
	if err = dp.InjectTo(&ctrl.deps); err != nil {
		return nil, err
	}
	if ctrl.view, err = ctrl.deps.Template.View(goathtml.DefaultLayout, "admin/model/fragment/insert", nil); err != nil {
		return nil, err
	}
	return ctrl, nil
}

// Get is handler to serve template where one can add new article
func (c *Insert) Get(scope app.Scope) {
	var deps struct {
		RequestError requestdep.Error     `request:"ErrorService"`
		Responser    requestdep.Responser `request:"ResponserService"`
	}
	if err := scope.InjectTo(&deps); err != nil {
		c.deps.Logger.ErrorLog("%v", err)
		deps.RequestError.Error(312, err)
		return
	}
	if err := deps.Responser.Execute(c.view, map[string]interface{}{
		"Valid":  msgcollection.NewMessageMap(),
		"Fields": entities.FragmentMainFields,
		"Form":   map[string]interface{}{},
	}); err != nil {
		c.deps.Logger.ErrorLog("%v", err)
		deps.RequestError.Error(312, err)
		return
	}
}

// Post is handler to save article from form obtained data
func (c *Insert) Post(scope app.Scope) {
	var (
		msgs messages.MessageMap
		deps struct {
			RequestError requestdep.Error     `request:"ErrorService"`
			Responser    requestdep.Responser `request:"ResponserService"`
			Request      *http.Request        `request:"Request"`
		}
	)
	if err := scope.InjectTo(&deps); err != nil {
		c.deps.Logger.ErrorLog("%v", err)
		deps.RequestError.Error(312, err)
		return
	}
	entity, err := httpfragment.NewForm(scope, entities.FragmentMainFields)
	if err != nil {
		c.deps.Logger.ErrorLog("%v", err)
		deps.RequestError.Error(312, err)
		return
	}
	if msgs, err = entities.ValidFragment(entity); err != nil {
		c.deps.Logger.ErrorLog("%v", err)
		deps.RequestError.Error(312, err)
		return
	}
	if len(msgs.GetAll()) != 0 {
		// valid error
		if err = deps.Responser.Execute(c.view, map[string]interface{}{
			"Valid":  msgs,
			"Form":   map[string]interface{}{},
			"Fields": entities.FragmentMainFields,
		}); err != nil {
			c.deps.Logger.ErrorLog("%v", err)
			deps.RequestError.Error(312, err)
			return
		}
		return
	}
	// valid success
	if _, err = c.deps.Inserter.Insert(scope, entity); err != nil {
		c.deps.Logger.ErrorLog("%v", err)
		deps.RequestError.Error(312, err)
		return
	}
	if err = scope.Trigger(app.CommitEvent, nil); err != nil {
		c.deps.Logger.ErrorLog("%v", err)
		deps.RequestError.Error(312, err)
		return
	}
	deps.Responser.Redirect("/admin/fragment")
}