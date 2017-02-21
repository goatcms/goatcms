package reqauth

import (
	"fmt"

	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcms/cmsapp/models"
	"github.com/goatcms/goatcms/cmsapp/services"
)

const (
	// UserID is a key for logged user id
	UserID = "dependency.user.auth.id"
)

// RequestAuth is global auth provider
type RequestAuth struct {
	deps struct {
		SessionManager services.SessionManager `request:"SessionManagerService"`
		Database       services.Database       `dependency:"DatabaseService"`
		LoginQuery     models.UserLoginQuery   `dependency:"db.query.user.LoginQuery"`
	}
}

// RequestAuthFactory create an authentification service instance
func RequestAuthFactory(dp dependency.Provider) (interface{}, error) {
	auth := &RequestAuth{}
	if err := dp.InjectTo(&auth.deps); err != nil {
		return nil, err
	}
	return services.RequestAuth(auth), nil
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
func (a *RequestAuth) Login(name, password string) (*models.User, error) {
	tx, err := a.deps.Database.TX()
	if err != nil {
		return nil, err
	}
	row, err := a.deps.LoginQuery(tx, name, password)
	if err != nil {
		return nil, err
	}
	user := &models.User{}
	if err = row.StructScan(user); err != nil {
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
