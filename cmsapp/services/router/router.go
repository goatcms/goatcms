package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/disk"
	"github.com/gorilla/mux"
)

// Router is request router system
type Router struct {
	deps struct {
		EventScope   app.Scope            `dependency:"EngineScope"`
		AppScope     app.Scope            `dependency:"AppScope"`
		Logger       services.Logger      `dependency:"LoggerService"`
		Host         string               `config:"?router.host"`
		StaticPrefix string               `config:"?router.static.prefix"`
		StaticPath   string               `config:"?router.static.path"`
		SecurityMode string               `config:"?router.security.mode"`
		SecurityCert string               `config:"?router.security.cert"`
		SecurityKey  string               `config:"?router.security.key"`
		ArgHost      string               `argument:"?host"`
		TmpFilespace filesystem.Filespace `filespace:"tmp"`
	}
	dependencyFactories map[string]dependency.Factory
	grouter             *mux.Router
	dp                  dependency.Provider
}

// RouterFactory is Router instance builder
func RouterFactory(dp dependency.Provider) (interface{}, error) {
	router := &Router{
		dependencyFactories: map[string]dependency.Factory{},
		grouter:             mux.NewRouter(),
		dp:                  dp,
	}
	if err := dp.InjectTo(&router.deps); err != nil {
		return nil, err
	}
	if router.deps.ArgHost != "" {
		router.deps.Host = router.deps.ArgHost
	}
	if router.deps.Host == "" {
		router.deps.Host = DefaultHost
	}
	if router.deps.StaticPath == "" {
		router.deps.StaticPath = DefaultStaticPath
	}
	if router.deps.StaticPrefix == "" {
		router.deps.StaticPrefix = DefaultStaticPrefix
	}
	if router.deps.SecurityMode == "" {
		router.deps.SecurityMode = TLSSecurityMode
	}
	if router.deps.SecurityCert == "" {
		router.deps.SecurityCert = "./data/certs/fullchain.pem"
	}
	if router.deps.SecurityKey == "" {
		router.deps.SecurityKey = "./data/certs/privkey.pem"
	}
	router.deps.SecurityMode = strings.ToUpper(router.deps.SecurityMode)
	if router.deps.SecurityMode != TLSSecurityMode && router.deps.SecurityMode != HTTPSecurityMode {
		router.deps.SecurityMode = TLSSecurityMode
	}
	fs := http.FileServer(http.Dir(router.deps.StaticPath))
	s := http.StripPrefix(router.deps.StaticPrefix, fs)
	router.grouter.PathPrefix(router.deps.StaticPrefix).Handler(s)
	return services.Router(router), nil
}

// OnGet append http get routing to global pool
func (router *Router) OnGet(path string, handler services.ScopeHandler) {
	router.grouter.HandleFunc(path, router.scopeHandlerToMuxHandler(handler)).Methods("GET")
}

// OnPost append http post routing to global pool
func (router *Router) OnPost(path string, handler services.ScopeHandler) {
	router.grouter.HandleFunc(path, router.scopeHandlerToMuxHandler(handler)).Methods("POST")
}

// OnPut append http put routing to global pool
func (router *Router) OnPut(path string, handler services.ScopeHandler) {
	router.grouter.HandleFunc(path, router.scopeHandlerToMuxHandler(handler)).Methods("PUT")
}

// OnDelete append http delete routing to global pool
func (router *Router) OnDelete(path string, handler services.ScopeHandler) {
	router.grouter.HandleFunc(path, router.scopeHandlerToMuxHandler(handler)).Methods("DELETE")
}

// On append http delete routing to global pool
func (router *Router) On(methods []string, path string, handler services.ScopeHandler) {
	router.grouter.HandleFunc(path, router.scopeHandlerToMuxHandler(handler)).Methods(methods...)
}

// Host return current host value
func (router *Router) Host() string {
	return router.deps.Host
}

// Start add http server
func (router *Router) Start() (err error) {
	srv := &http.Server{
		Addr:    router.deps.Host,
		Handler: router.grouter,
	}
	if router.deps.SecurityMode == TLSSecurityMode {
		if !disk.IsFile(router.deps.SecurityCert) {
			panic(fmt.Sprintf("'%v' [config: router.security.cert] certificate file is not exist.", router.deps.SecurityCert))
		}
		if !disk.IsFile(router.deps.SecurityKey) {
			panic(fmt.Sprintf("'%v' [config: router.security.cert] certificate file is not exist.", router.deps.SecurityKey))
		}
		if err = srv.ListenAndServeTLS(router.deps.SecurityCert, router.deps.SecurityKey); err != nil {
			return err
		}
	} else {
		if err = srv.ListenAndServe(); err != nil {
			return err
		}
	}
	router.deps.AppScope.On(app.CloseEvent, func(interface{}) (err error) {
		return srv.Shutdown(nil)
	})
	return nil
}

// AddFactory add routing to global pool
func (router *Router) AddFactory(name string, factory dependency.Factory) error {
	if _, ok := router.dependencyFactories[name]; ok {
		return fmt.Errorf("router.AddFactory: Add %s dependency twice", name)
	}
	router.dependencyFactories[name] = factory
	return nil
}

func (router *Router) newRequestScope(w http.ResponseWriter, r *http.Request) app.Scope {
	rs := NewRequestScope(router.dp, router.deps.TmpFilespace, router.dependencyFactories, w, r)
	router.deps.EventScope.Trigger(services.CreateRequestScope, rs)
	return rs
}

func (router *Router) scopeHandlerToMuxHandler(handler services.ScopeHandler) services.MuxHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			deps struct {
				RequestError requestdep.Error          `request:"ErrorService"`
				Session      requestdep.SessionManager `request:"SessionService"`
				Logger       services.Logger           `dependency:"LoggerService"`
			}
			err   error
			scope = router.newRequestScope(w, r)
		)
		if err = scope.InjectTo(&deps); err != nil {
			panic(err)
		}
		if err = deps.Session.LoadSession(); err != nil {
			deps.Logger.TestLog("%v", err)
		}
		if err = handler(scope); err != nil {
			deps.RequestError.DO(err)
			return
		}
		if err := scope.Trigger(app.CloseEvent, nil); err != nil {
			deps.RequestError.DO(err)
			return
		}
	}
}
