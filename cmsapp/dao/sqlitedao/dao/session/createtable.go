package dao

import (
	"database/sql"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

// SessionCreateTable is a Data Access Object for session entity
type SessionCreateTable struct {
	deps struct {
		DB *sql.DB `dependency:"db0.engine"`
	}
}

func NewSessionCreateTable(di dependency.Injector) (*SessionCreateTable, error) {
	instance := &SessionCreateTable{}
	if err := di.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	return instance, nil
}

func SessionCreateTableFactory(dp dependency.Provider) (interface{}, error) {
	instance, err := NewSessionCreateTable(dp)
	if err != nil {
		return nil, err
	}
	return maindef.CreateTable(instance), nil
}

func (dao SessionCreateTable) CreateTable(scope app.Scope) (err error) {
	var tx *sql.Tx
	if tx, err = helpers.TX(scope, dao.deps.DB); err != nil {
		return err
	}
	_, err = tx.Exec(dao.SQL())
	return err
}

func (dao SessionCreateTable) SQL() string {
	return `CREATE TABLE IF NOT EXISTS Session (ID INTEGER PRIMARY KEY, Secret TEXT UNIQUE NOT NULL, UserID INTEGER NOT NULL);`
}

func (dao SessionCreateTable) AlterSQL() string {
	return `
		ALTER TABLE Session ADD CONSTRAINT fk_session_user
			FOREIGN KEY (UserID)
			REFERENCES User(ID)
			;`
}
