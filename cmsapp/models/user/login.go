package user

import (
	"fmt"

	"github.com/goatcms/goatcms/cmsapp/models"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/db"
	"github.com/goatcms/goatcore/dependency"
)

// LoginContext is context for Login function
type LoginContext struct {
	deps struct {
		Table db.Table       `dependency:"UserTable"`
		DSQL  db.DSQL        `dependency:"DSQL"`
		Crypt services.Crypt `dependency:"CryptService"`
	}
	query string
}

// Insert create new record
func (q LoginContext) Login(tx db.TX, name, password string) (*models.User, error) {
	var (
		err          error
		hashPassword string
		row          db.Row
		ok           bool
	)
	if row, err = tx.QueryRowx(q.query, name, hashPassword); err != nil {
		return nil, err
	}
	user := &models.User{}
	if err = row.StructScan(user); err != nil {
		return nil, err
	}
	if ok, err = q.deps.Crypt.Compare(user.PasswordHash, password); err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("Incorrect password")
	}
	return user, err
}

// LoginContext create new CreateTable function
func LoginFactory(dp dependency.Provider) (interface{}, error) {
	var err error
	ctx := &LoginContext{}
	if err = dp.InjectTo(&ctx.deps); err != nil {
		return nil, err
	}
	ctx.query, err = ctx.deps.DSQL.NewSelectWhereSQL(ctx.deps.Table.Name(), ctx.deps.Table.Fields(), "Email = :$1")
	if err != nil {
		return nil, err
	}
	return ctx.Login, nil
}
