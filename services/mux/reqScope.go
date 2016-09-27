package mux

import (
	"log"
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

// Commit persist scope data
func (rs *RequestScope) Commit() error {
	if err := rs.Scope.Trigger(scope.CommitEvent); err != nil {
		return err
	}
	if rs.tx != nil {
		return rs.tx.Commit()
	}
	return nil
}

// Rollback remove changes
func (rs *RequestScope) Rollback() error {
	if err := rs.Scope.Trigger(scope.RollbackEvent); err != nil {
		return err
	}
	if rs.tx != nil {
		return rs.tx.Rollback()
	}
	return nil
}

// Error notice error and show user error page
func (rs *RequestScope) Error(err error) {
	rs.Set(scope.Error, err)
	rs.Trigger(scope.ErrorEvent)
	log.Println(err)
	http.Error(rs.Response(), err.Error(), http.StatusInternalServerError)
	return
}

// Fatal notice error and kill scope
func (rs *RequestScope) Fatal(err error) {
	rs.Set(scope.Error, err)
	rs.Trigger(scope.ErrorEvent)
	log.Fatal(err)
	http.Error(rs.Response(), err.Error(), http.StatusInternalServerError)
	rs.Trigger(scope.KillEvent)
}
