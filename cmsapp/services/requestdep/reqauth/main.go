package reqauth

import (
	"fmt"

	"github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

const (
	// UserID is a key for logged user id
	UserID = "dependency.user.auth.id"
)

// RequestAuth is global auth provider
type RequestAuth struct {
	deps struct {
		SessionManager services.SessionManager `request:"SessionService"`
		Scope          app.Scope               `request:"RequestScope"`
		LoginQuery     dao.UserLoginQuery      `dependency:"UserLoginQuery"`
	}
}

// RequestAuthFactory create an authentification service instance
func AuthFactory(dp dependency.Provider) (interface{}, error) {
	auth := &RequestAuth{}
	if err := dp.InjectTo(&auth.deps); err != nil {
		return nil, err
	}
	return requestdep.Auth(auth), nil
}

// UserID get logged user id from current session
func (a *RequestAuth) UserID() (string, error) {
	id, err := a.deps.SessionManager.Get(UserID)
	if err != nil {
		return "", err
	}
	if id == "" {
		return "", fmt.Errorf("User session expired")
	}
	return id.(string), nil
}

func (a *RequestAuth) Login(name, password string) (user *entities.User, err error) {
	var (
		row dao.Row
	)
	if row, err = a.deps.LoginQuery.Login(a.deps.Scope, []string{"id"}, &dao.UserLoginQueryParams{
		Login:    name,
		Email:    name,
		Password: password,
	}); err != nil {
		return nil, err
	}
	user = &entities.User{}
	if err := row.StructScan(user); err != nil {
		return nil, err
	}
	if err := a.deps.SessionManager.Set(UserID, user.ID); err != nil {
		return nil, err
	}
	return user, nil
}

// Clear remove a user id from session
func (a *RequestAuth) Clear() error {
	if err := a.deps.SessionManager.Set(UserID, nil); err != nil {
		return err
	}
	return nil
}
