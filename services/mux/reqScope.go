package mux

import (
	"net/http"

	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goat-core/scope"
	"github.com/goatcms/goat-core/scope/corescope"
	"github.com/goatcms/goatcms/services"
	"github.com/jmoiron/sqlx"
)

// RequestScope is simple request scope
type RequestScope struct {
	scope.Scope

	req      *http.Request
	res      http.ResponseWriter
	database services.Database
	tx       *sqlx.Tx
}

// NewRequestScope create new instance of scope
func NewRequestScope(dp dependency.Provider, res http.ResponseWriter, req *http.Request) *RequestScope {
	return &RequestScope{
		Scope: corescope.NewScope(dp),
		res:   res,
		req:   req,
		tx:    nil,
	}
}

// TX return current transaction
func (rs *RequestScope) TX() (*sqlx.Tx, error) {
	var err error
	if rs.tx != nil {
		return rs.tx, nil
	}
	rs.tx, err = rs.database.Adapter().Beginx()
	return rs.tx, err
}

// Request return current request
func (rs *RequestScope) Request() *http.Request {
	return rs.req
}

// Response return current response
func (rs *RequestScope) Response() http.ResponseWriter {
	return rs.res
}
