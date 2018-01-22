package reqsession

import (
	"fmt"
	"net/http"
	"time"

	"github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
)

const X_AUTH_TOKEN_HEADER = "X-Auth-Token"

// SessionManager provide session manager for reques (current user)
type SessionManager struct {
	deps struct {
		Request         *http.Request           `request:"Request"`
		RequestScope    app.Scope               `request:"RequestScope"`
		Response        http.ResponseWriter     `request:"Response"`
		Logger          services.Logger         `dependency:"LoggerService"`
		Manager         services.SessionManager `dependency:"SessionManager"`
		SessionCookieID string                  `config:"?session.cookie.id"`
	}
	session *entities.Session
}

// SessionFactory create a session manager instance
func SessionFactory(dp dependency.Provider) (i interface{}, err error) {
	s := &SessionManager{}
	s.deps.SessionCookieID = services.SessionCookieID
	if err = dp.InjectTo(&s.deps); err != nil {
		return nil, err
	}
	return requestdep.SessionManager(s), nil
}

// LoadSession read session secret from X-Auth-Token header or cookie
func (s *SessionManager) LoadSession() (err error) {
	var (
		secret string
	)
	if secret = s.deps.Request.Header.Get(X_AUTH_TOKEN_HEADER); secret == "" {
		var cookie *http.Cookie
		if cookie, err = s.deps.Request.Cookie(s.deps.SessionCookieID); err != nil {
			return err
		}
		secret = cookie.Value
	}
	if s.session, err = s.deps.Manager.Get(s.deps.RequestScope, secret); err != nil {
		s.deps.Logger.DevLog("%v: remove fail session", err)
		if err = s.DestroySession(); err != nil {
			return err
		}
		return nil
	}
	return nil
}

// Get return request session
func (s *SessionManager) Get() (session *entities.Session, err error) {
	if s.session == nil {
		return nil, fmt.Errorf("session does not inited")
	}
	return s.session, nil
}

// CreateSession build new session for user
func (s *SessionManager) CreateSession(user *entities.User) (err error) {
	if s.session, err = s.deps.Manager.Create(s.deps.RequestScope, user); err != nil {
		return err
	}
	expires := time.Unix(*s.session.Expires, 0)
	cookie := http.Cookie{
		Name:     s.deps.SessionCookieID,
		Value:    *s.session.Secret,
		Expires:  expires,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(s.deps.Response, &cookie)
	return nil
}

// DestroySession remove session cookie
func (s *SessionManager) DestroySession() (err error) {
	var cookie *http.Cookie
	if cookie, err = s.deps.Request.Cookie(s.deps.SessionCookieID); err != nil {
		return err
	}
	cookie = &http.Cookie{
		Name:    s.deps.SessionCookieID,
		Value:   "",
		Expires: time.Unix(0, 0),
		Path:    "/",
	}
	http.SetCookie(s.deps.Response, cookie)
	if err = s.deps.Manager.Delete(s.deps.RequestScope, cookie.Value); err != nil {
		return err
	}
	return nil
}
