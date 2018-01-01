package helpers

import (
	"github.com/goatcms/goatcore/app"
	"github.com/jmoiron/sqlx"
)

const (
	TXKey        = "_dbtx"
	CommitInited = "_dbtx_commit_inited"
)

func TX(scope app.Scope, db *sqlx.DB) (tx *sqlx.Tx, err error) {
	var ins interface{}
	ins, err = scope.Get(TXKey)
	if err != nil || ins == nil {
		tx, err = db.Beginx()
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
		tx = ins.(*sqlx.Tx)
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
	tx := ins.(*sqlx.Tx)
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
	tx := ins.(*sqlx.Tx)
	scope.Set(TXKey, nil)
	return true, tx.Rollback()
}

func NewMemoryDB() (db *sqlx.DB, err error) {
	if db, err = sqlx.Open("sqlite3", ":memory:"); err != nil {
		return nil, err
	}
	return db, nil
}
