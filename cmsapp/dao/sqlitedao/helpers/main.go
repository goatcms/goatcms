package helpers

import (
	"database/sql"
	"github.com/goatcms/goatcore/app"
	"strconv"
)

const (
	TXKey        = "_dbtx"
	CommitInited = "_dbtx_commit_inited"
)

func TX(scope app.Scope, db *sql.DB) (tx *sql.Tx, err error) {
	var ins interface{}
	ins, err = scope.Get(TXKey)
	if err != nil || ins == nil {
		tx, err = db.Begin()
		if err != nil {
			return nil, err
		}
		scope.Set(TXKey, tx)
		// add commit event callback
		ins, err = scope.Get(CommitInited)
		if err != nil || ins == nil {
			scope.On(app.CommitEvent, func(interface{}) error {
				_, err := Commit(scope)
				return err
			})
		}
	} else {
		tx = ins.(*sql.Tx)
	}
	return tx, nil
}

func Commit(scope app.Scope) (isCommited bool, err error) {
	var ins interface{}
	ins, err = scope.Get(TXKey)
	if err != nil || ins == nil {
		// nothing to commit
		return false, nil
	}
	tx := ins.(*sql.Tx)
	scope.Set(TXKey, nil)
	return true, tx.Commit()
}

func Rollback(scope app.Scope) (isRollback bool, err error) {
	var ins interface{}
	ins, err = scope.Get(TXKey)
	if err != nil || ins == nil {
		// nothing to commit
		return false, nil
	}
	tx := ins.(*sql.Tx)
	scope.Set(TXKey, nil)
	return true, tx.Rollback()
}

func NewMemoryDB() (db *sql.DB, err error) {
	if db, err = sql.Open("sqlite3", ":memory:"); err != nil {
		return nil, err
	}
	return db, nil
}

func Quote(s *string) string {
	if s == nil {
		return "null"
	}
	return strconv.Quote(*s)
}

func FormatInt(i *int64, base int) string {
	if i == nil {
		return "null"
	}
	return strconv.FormatInt(*i, base)
}
