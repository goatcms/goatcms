package user

import (
	"github.com/goatcms/goatcore/db"
	"github.com/goatcms/goatcms/cmsapp/models"
)

// CreateTableContext is context for findByID function
type LoginContext struct {
	query string
}

// Insert create new record
func (q LoginContext) Login(tx db.TX, name, password string) (db.Row, error) {
	row, err := tx.QueryRowx(q.query, name, password)
	return row.(db.Row), err
}

// CreateTableContext create new CreateTable function
func NewLogin(table db.Table, dsql db.DSQL) (models.UserLoginQuery, error) {
	query, err := dsql.NewSelectWhereSQL(table.Name(), table.Fields(), "email = :$1 AND hashPassword = :$2 ")
	if err != nil {
		return nil, err
	}
	context := &LoginContext{
		query: query,
	}
	return context.Login, nil
}
