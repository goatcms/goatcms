package mux

import (
	"net/http"

	"github.com/goatcms/goat-core/scope/corescope"
	"github.com/goatcms/goatcms/services"
	gorillamux "github.com/gorilla/mux"
)

// RouterHandler function for routing dispatcher
type RouterHandler func(http.ResponseWriter, *http.Request)

// muxToRouterHandler convert handler from MuxHandler to RequestHandler
func muxToRouterHandler(m *Mux, handler services.MuxHandler) RouterHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		scope := &RequestScope{
			Scope:    corescope.NewScope(m.dp),
			res:      w,
			req:      r,
			database: m.database,
			tx:       nil,
		}
		handler(scope)
	}
}

// Mux is global routing provider
type Mux struct {
	dp       services.Provider
	database services.Database
	router   *gorillamux.Router
}

// NewMux create a mux instance
func NewMux(dp services.Provider) (*Mux, error) {
	db, err := dp.Database()
	if err != nil {
		return nil, err
	}
	return &Mux{
		dp:       dp,
		database: db,
		router:   gorillamux.NewRouter(),
	}, nil
}

// Get append http get routing to global pool
func (m *Mux) Get(path string, handler services.MuxHandler) {
	m.router.HandleFunc(path, muxToRouterHandler(m, handler)).Methods("GET")
}

// Post append http post routing to global pool
func (m *Mux) Post(path string, handler services.MuxHandler) {
	m.router.HandleFunc(path, muxToRouterHandler(m, handler)).Methods("POST")
}

// Put append http put routing to global pool
func (m *Mux) Put(path string, handler services.MuxHandler) {
	m.router.HandleFunc(path, muxToRouterHandler(m, handler)).Methods("PUT")
}

// Delete append http delete routing to global pool
func (m *Mux) Delete(path string, handler services.MuxHandler) {
	m.router.HandleFunc(path, muxToRouterHandler(m, handler)).Methods("DELETE")
}

// Start add routing to global pool
func (m *Mux) Start() error {
	http.Handle("/", m.router)
	if err := http.ListenAndServe(":5555", nil); err != nil {
		return err
	}
	return nil
}
