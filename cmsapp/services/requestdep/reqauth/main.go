package reqauth

import (
	"fmt"

	"github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcms2/cmsapp/models"
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
		Database       services.Database       `dependency:"DatabaseService"`
		Login          models.UserLogin        `dependency:"UserLogin"`
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

// RequestAuth save a user id into session
func (a *RequestAuth) Login(name, password string) (*entities.User, error) {
	tx, err := a.deps.Database.TX()
	if err != nil {
		return nil, err
	}
	user, err := a.deps.Login(tx, name, password)
	if err != nil {
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
