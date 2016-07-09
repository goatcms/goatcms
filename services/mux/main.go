package mux

import (
	"net/http"

	gorillamux "github.com/gorilla/mux"
	"github.com/goatcms/goatcms/services"
)

//Mux is global routing provider
type Mux struct {
	router *gorillamux.Router
}

//NewMux create a mux instance
func NewMux() (*Mux, error) {
	return &Mux{
		router: gorillamux.NewRouter(),
	}, nil
}

//Get append http get routing to global pool
func (m *Mux) Get(path string, handler services.MuxHandler) {
	m.router.HandleFunc(path, handler).Methods("GET")
}

//Post append http get routing to global pool
func (m *Mux) Post(path string, handler services.MuxHandler) {
	m.router.HandleFunc(path, handler).Methods("POST")
}

//Put append http get routing to global pool
func (m *Mux) Put(path string, handler services.MuxHandler) {
	m.router.HandleFunc(path, handler).Methods("PUT")
}

//Delete append http get routing to global pool
func (m *Mux) Delete(path string, handler services.MuxHandler) {
	m.router.HandleFunc(path, handler).Methods("DELETE")
}

//Start add routing to globall poll
func (m *Mux) Start() error {
	http.Handle("/", m.router)
	if err := http.ListenAndServe(":5555", nil); err != nil {
		return err
	}
	return nil
}
