package user

import (
	"github.com/goatcms/goatcms/cmsapp/models"
	"github.com/goatcms/goatcore/db"
)

// LoginContext is context for Login function
type LoginContext struct {
	query string
}

// Insert create new record
func (q LoginContext) Login(tx db.TX, name, password string) (db.Row, error) {
	row, err := tx.QueryRowx(q.query, name, password)
	return row.(db.Row), err
}

// LoginContext create new CreateTable function
func NewLogin(table db.Table, dsql db.DSQL) (models.UserLogin, error) {
	query, err := dsql.NewSelectWhereSQL(table.Name(), table.Fields(), "email = :$1 AND hashPassword = :$2 ")
	if err != nil {
		return nil, err
	}
	context := &LoginContext{
		query: query,
	}
	return context.Login, nil
}
