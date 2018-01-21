package session

import (
	"fmt"
	"strconv"
	"time"

	"github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/varutil"
)

// SessionsManager store sessions in database
type SessionsManager struct {
	deps struct {
		SecretLength string                     `config:"?session.secret_length"`
		Lifetime     string                     `config:"?session.lifetime"`
		Search       dao.SessionCriteriaSearch  `dependency:"SessionCriteriaSearch"`
		Inserter     dao.SessionInsert          `dependency:"SessionInsert"`
		Deleter      dao.SessionCriteriaDeleter `dependency:"SessionCriteriaDeleter"`
	}
	secretLength int
	lifetime     int64
	cache        map[string]*entities.Session
}

// SessionsManagerFactory create a session manager instance
func SessionsManagerFactory(dp dependency.Provider) (interface{}, error) {
	var err error
	instance := &SessionsManager{
		cache: map[string]*entities.Session{},
	}
	if err = dp.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	if instance.deps.SecretLength == "" {
		instance.secretLength = 120
	} else {
		if instance.secretLength, err = strconv.Atoi(instance.deps.SecretLength); err == nil {
			return nil, err
		}
	}
	if instance.deps.Lifetime == "" {
		instance.lifetime = 60 * 60 * 24 * 365
	} else {
		if instance.lifetime, err = strconv.ParseInt(instance.deps.Lifetime, 10, 64); err == nil {
			return nil, err
		}
	}
	return services.SessionManager(instance), nil
}

// Get return session by secret
func (s *SessionsManager) Get(scope app.Scope, secret string) (session *entities.Session, err error) {
	var (
		rows dao.SessionRows
		ok   bool
	)
	if session, ok = s.cache[secret]; ok {
		if *session.Expires < time.Now().Unix() {
			delete(s.cache, secret)
			return nil, fmt.Errorf("Session %v expired or is incorrect", secret)
		}
		return session, nil
	}
	if rows, err = s.deps.Search.Find(scope, &dao.SessionCriteria{
		Fields: entities.SessionAllFieldsAndID,
		Where: dao.SessionCriteriaWhere{
			Secret: &dao.StringFieldCriteria{
				Type:  dao.EQ,
				Value: []string{secret},
			},
		},
		Related: dao.SessionCriteriaRelated{
			User: &dao.UserCriteria{
				Fields: entities.UserMainFieldsAndID,
			},
		},
	}, &dao.Pager{
		Limit:  1,
		Offset: 0,
	}); err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, fmt.Errorf("unknow session by %v", secret)
	}
	if session, err = rows.Get(); err != nil {
		return nil, err
	}
	if *session.Expires < time.Now().Unix() {
		return nil, fmt.Errorf("Session %v expired", secret)
	}
	s.cache[secret] = session
	return session, nil
}

// Create create a new session instance, persist it in database and return
func (s *SessionsManager) Create(scope app.Scope, user *entities.User) (session *entities.Session, err error) {
	secret := varutil.RandString(s.secretLength, varutil.StrongBytes)
	expires := time.Now().Unix() + s.lifetime
	session = &entities.Session{
		Secret:  &secret,
		Expires: &expires,
		UserID:  user.ID,
		User:    user,
	}
	if _, err = s.deps.Inserter.Insert(scope, session); err != nil {
		return nil, err
	}
	return session, nil
}

func (s *SessionsManager) Delete(scope app.Scope, secret string) (err error) {
	if err = s.deps.Deleter.Delete(scope, &dao.SessionCriteria{
		Where: dao.SessionCriteriaWhere{
			Secret: &dao.StringFieldCriteria{
				Type:  dao.EQ,
				Value: []string{secret},
			},
		},
	}); err != nil {
		return err
	}
	return nil
}
