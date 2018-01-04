package user

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/goathtml"
)

// List is a controler to show a list of article
type List struct {
	deps struct {
		Template services.Template 					`dependency:"TemplateService"`
		Finder   dao.UserSearch  `dependency:"UserSearch"`
	}
	view *template.Template
}

// NewList create instance of a list user controller
func NewList(dp dependency.Provider) (*List, error) {
	var err error
	ctrl := &List{}
	if err = dp.InjectTo(&ctrl.deps); err != nil {
		return nil, err
	}
	if ctrl.view, err = ctrl.deps.Template.View(goathtml.DefaultLayout, "admin/model/user/list", nil); err != nil {
		return nil, err
	}
	return ctrl, nil
}

// Get is handler to serve template where one can add new article
func (c *List) Get(requestScope app.Scope) {
	var (
		rows        dao.UserRows
		err         error
		entity      *entities.User
		requestDeps struct {
			RequestError requestdep.Error     `request:"ErrorService"`
			Responser    requestdep.Responser `request:"ResponserService"`
			Request      *http.Request        `request:"Request"`
		}
		searchParams dao.UserSearchParams
	)
	if err = requestScope.InjectTo(&requestDeps); err != nil {
		fmt.Println(err)
		return
	}
	if rows, err = c.deps.Finder.Search(requestScope, entities.UserMainFields, &searchParams); err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
	collection := []*entities.User{}
	for rows.Next() {
		if entity, err = rows.Get(); err != nil {
			requestDeps.RequestError.Error(312, err)
			return
		}
		collection = append(collection, entity)
	}
	requestScope.On(app.ErrorEvent, func(erri interface{}) error {
		scopeErr := erri.(error)
		requestDeps.RequestError.Errorf(403, "%s", scopeErr.Error())
		return nil
	})
	if err = requestDeps.Responser.Execute(c.view, map[string]interface{}{
		"Collection": collection,
		"Labels":     entities.UserMainFields,
	}); err != nil {
		requestDeps.RequestError.Error(312, err)
		return
	}
}