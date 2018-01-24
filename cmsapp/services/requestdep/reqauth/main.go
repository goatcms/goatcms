package reqauth

import (
	"github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

const (
	// UserID is a key for logged user id
	UserID = "dependency.user.auth.id"
)

// Auth is global auth provider
type Auth struct {
	deps struct {
		SessionManager requestdep.SessionManager `request:"SessionService"`
		Scope          app.Scope                 `request:"RequestScope"`
		SigninQuery    dao.UserSigninQuery       `dependency:"UserSigninQuery"`
		UserFindByID   dao.UserFindByID          `dependency:"UserFindByID"`
	}
	user *entities.User
}

// AuthFactory create an authentification service instance
func AuthFactory(dp dependency.Provider) (interface{}, error) {
	auth := &Auth{}
	if err := dp.InjectTo(&auth.deps); err != nil {
		return nil, err
	}
	return requestdep.Auth(auth), nil
}

// Signin authorize user by username and passwrd and create session if success
func (a *Auth) Signin(name, password string) (session *entities.Session, err error) {
	var user *entities.User
	if user, err = a.deps.SigninQuery.Signin(a.deps.Scope, []string{"ID"}, &dao.UserSigninQueryParams{
		Username: name,
		Email:    name,
	}); err != nil {
		return nil, err
	}
	if session, err = a.deps.SessionManager.CreateSession(user); err != nil {
		return nil, err
	}
	return session, nil
}

// Signout destroy user session (logout current user)
func (a *Auth) Signout() error {
	if err := a.deps.SessionManager.DestroySession(); err != nil {
		return err
	}
	return nil
}
