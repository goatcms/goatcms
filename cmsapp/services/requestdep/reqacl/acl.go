package reqacl

import (
	"strings"

	"github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/varutil"
)

const (
	// SuperAdminRole is super administrator role name (it take all system provilage)
	SuperAdminRole = "superadmin"
)

var (
	// DefaultRoles contains roles for unauthorized guests
	DefaultRoles = []string{"anonymous"}
	// UserAdditionalRoles contains additional roles for authorized users
	UserAdditionalRoles = []string{"loggedin"}
)

// ACL is Access Control List object
type ACL struct {
	deps struct {
		Logger         services.Logger           `dependency:"LoggerService"`
		SessionManager requestdep.SessionManager `request:"SessionService"`
		Scope          app.Scope                 `request:"RequestScope"`
	}
	isSuperAdmin bool
	roles        []string
}

// ACLFactory create an Access Control List service instance
func ACLFactory(dp dependency.Provider) (interface{}, error) {
	var (
		err     error
		session *entities.Session
	)
	instance := &ACL{}
	if err = dp.InjectTo(&instance.deps); err != nil {
		return nil, err
	}
	if session, err = instance.deps.SessionManager.Get(); err != nil {
		instance.deps.Logger.DevLog("%v", err)
		instance.roles = DefaultRoles
	} else {
		if session != nil && session.User != nil && session.User.Roles != nil {
			instance.roles = strings.Split(*session.User.Roles, " ")
		} else {
			instance.roles = []string{}
		}
		instance.roles = append(instance.roles, UserAdditionalRoles...)
	}
	instance.isSuperAdmin = varutil.IsArrContainStr(instance.roles, SuperAdminRole)
	instance.deps.Logger.DevLog("ACL.ACLFactory: Is superadmin %v and has %v roles", instance.isSuperAdmin, instance.roles)
	return requestdep.ACL(instance), nil
}

// HasAnyRole return true if has any of roles
func (acl *ACL) HasAnyRole(roles []string) bool {
	if acl.isSuperAdmin {
		acl.deps.Logger.DevLog("ACL.HasAnyRole: User is superadmin. Superadmin has all privileges. (asked for roles: %v)", roles)
		return true
	}
	for _, role := range roles {
		if varutil.IsArrContainStr(acl.roles, role) {
			acl.deps.Logger.DevLog("ACL.HasAnyRole: User must have one role of %v. And the user has %v (%v) role", roles, role, acl.roles)
			return true
		}
	}
	return false
}
