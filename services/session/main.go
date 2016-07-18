package session

import (
	"fmt"
	"net/http"
	"time"

	"github.com/goatcms/goat-core/varutil"
)

const (
	sessionCookieID       = "session"
	sessionCookieLen      = 128
	sessionCookieLifetime = 365 * 24
)

// SessionData is single record in sessions map (represent single session data)
type SessionData map[string]string

// SessionManager is global session provider
type SessionManager struct {
	sessions map[string]SessionData
}

// NewSessionManager create a session manager instance
func NewSessionManager() (*SessionManager, error) {
	return &SessionManager{
		sessions: map[string]SessionData{},
	}, nil
}

// Create build new session
func (s *SessionManager) create(w http.ResponseWriter) (string, error) {
	sessid := varutil.RandString(sessionCookieLen, varutil.StrongBytes)
	s.sessions[sessid] = SessionData{}
	expiration := time.Now().Add(sessionCookieLifetime * time.Hour)
	cookie := http.Cookie{
		Name:     sessionCookieID,
		Value:    sessid,
		Expires:  expiration,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	return sessid, nil
}

// Init build new session if a session doesn't exist
func (s *SessionManager) Init(w http.ResponseWriter, r *http.Request) (string, error) {
	var (
		cookie *http.Cookie
		err    error
	)
	if cookie, err = r.Cookie(sessionCookieID); err != nil {
		return s.create(w)
	}
	return cookie.Value, nil
}

// Get return value by name for session selected by session id
func (s *SessionManager) Get(id, name string) (string, error) {
	sessionRow, ok := s.sessions[id]
	if !ok {
		return "", fmt.Errorf("Session " + id + " doesn't exist")
	}
	value, ok := sessionRow[name]
	if !ok {
		return "", fmt.Errorf("Session key " + name + " doesn't exist")
	}
	return value, nil
}

// Set return value by name for session selected by session id
func (s *SessionManager) Set(id, name, value string) error {
	sessionRow, ok := s.sessions[id]
	if !ok {
		return fmt.Errorf("Session " + id + " doesn't exist")
	}
	sessionRow[name] = value
	return nil
}
