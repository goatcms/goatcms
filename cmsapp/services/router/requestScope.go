package router

import (
	"net/http"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/dependency/provider"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcms/cmsapp/services"
)

const _24M = (1 << 20) * 24

type RequestScope struct {
	app.EventScope
	dataScope app.DataScope
	injector  app.Injector

	parent app.Scope
}

func NewRequestScope(dp dependency.Provider, fs filesystem.Filespace, factories map[string]dependency.Factory, w http.ResponseWriter, r *http.Request) app.Scope {
	ds := &scope.DataScope{
		Data: make(map[string]interface{}),
	}
	eventScope := scope.NewEventScope()
	formInjector := FormInjector{
		req:            r,
		filespace:      fs,
		maxMemFileSize: _24M,
		tagname:        services.FormTagName,
		eventScope:     eventScope,
	}
	requestScope := &scope.Scope{
		EventScope: eventScope,
		DataScope:  ds,
		Injector:   nil,
	}
	requestDependencyProvider := provider.NewStaticProvider(services.RequestTagName, factories, map[string]interface{}{
		"Response":     w,
		"Request":      r,
		"RequestScope": app.Scope(requestScope),
	}, []dependency.Injector{
		dp,
		formInjector,
	})
	requestScope.Injector = requestDependencyProvider
	return requestScope
}
