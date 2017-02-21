package session

import (
	"fmt"
	"strconv"
	"time"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/varutil"
	"github.com/goatcms/goatcms/cmsapp/services"
)

// MemorySessionStorage store sessions in memory
type MemorySessionStorage struct {
	deps struct {
		SessionIDLength string `config:"?session.cookie.length"`
		SessionLifetime string `config:"?session.cookie.lifetime"`
	}
	sessions map[string]app.DataScope
}

// NewSessionManager create a session manager instance
func MemorySessionStorageFactory(dp dependency.Provider) (interface{}, error) {
	ss := &MemorySessionStorage{
		sessions: map[string]app.DataScope{},
	}
	if err := dp.InjectTo(&ss.deps); err != nil {
		return nil, err
	}
	return services.SessionStorage(ss), nil
}

// getSession return session map by session id
func (s *MemorySessionStorage) Get(id string) (app.DataScope, error) {
	dataScope, ok := s.sessions[id]
	if !ok {
		dataScope = scope.NewDataScope(map[string]interface{}{})
		s.sessions[id] = dataScope
		return dataScope, nil
	}
	expirationIns, err := dataScope.Get(services.SessionExpireKey)
	if err != nil {
		return nil, err
	}
	expiration := expirationIns.(int64)
	if expiration < time.Now().Unix() {
		delete(s.sessions, id)
		return nil, fmt.Errorf("Session expired")
	}
	return dataScope, nil
}

// getSession return session map by session id
func (s *MemorySessionStorage) Create() (string, app.DataScope, error) {
	var err error
	var sessionLifetime int64
	var sessionIDLength int
	if sessionIDLength, err = s.sessionIDLength(); err != nil {
		return "", nil, err
	}
	if sessionLifetime, err = s.SessionLifetime(); err != nil {
		return "", nil, err
	}
	sessionID := varutil.RandString(sessionIDLength, varutil.StrongBytes)
	dataScope := scope.NewDataScope(map[string]interface{}{})
	s.sessions[sessionID] = dataScope
	expiration := time.Now().Add(time.Duration(sessionLifetime) * time.Hour).Unix()
	dataScope.Set(services.SessionExpireKey, expiration)
	// TODO: remove session after expired time
	return sessionID, dataScope, nil
}

func (s *MemorySessionStorage) sessionIDLength() (int, error) {
	if s.deps.SessionIDLength != "" {
		return strconv.Atoi(s.deps.SessionIDLength)
	}
	return services.SessionIDLength, nil
}

func (s *MemorySessionStorage) SessionLifetime() (int64, error) {
	if s.deps.SessionLifetime != "" {
		return strconv.ParseInt(s.deps.SessionLifetime, 10, 64)
	}
	return services.SessionLifetime, nil
}
