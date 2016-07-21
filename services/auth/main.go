package auth

import (
	"fmt"

	"github.com/goatcms/goatcms/services"
)

const (
	// sessionLoginUserId is a key for user in session
	sessionLoginUserID = "loginUserId"
)

// Auth is global auth provider
type Auth struct {
	sess services.SessionManager
}

// NewAuth create a authentification service instance
func NewAuth(dp services.Provider) (*Auth, error) {
	sess, err := dp.SessionManager()
	if err != nil {
		return nil, err
	}
	return &Auth{
		sess: sess,
	}, nil
}

// GetUserID retrieve user id from session
func (a *Auth) GetUserID(sessid string) (string, error) {
	id, err := a.sess.Get(sessid, sessionLoginUserID)
	if err != nil {
		return "", err
	}
	if id == "" {
		return "", fmt.Errorf("User session expired")
	}
	return id, nil
}

// Auth save a user id into session
func (a *Auth) Auth(sessid string, userid string) error {
	if err := a.sess.Set(sessid, sessionLoginUserID, userid); err != nil {
		return err
	}
	return nil
}

// Clear remove a user id from session
func (a *Auth) Clear(sessid string) error {
	return a.Auth(sessid, "")
}
