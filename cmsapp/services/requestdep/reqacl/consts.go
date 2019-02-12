package reqacl

const (
	// SuperAdminRole is super administrator role name (it take all system provilage)
	SuperAdminRole = "superadmin"
)

var (
	// DefaultAnonymousRoles contains roles for unauthorized guests
	DefaultAnonymousRoles = []string{"anonymous"}
	// anonymousRoles contains roles for unauthorized (anonymous) users
	cache *Cache
	// UserAdditionalRoles contains additional roles for authorized users
	UserAdditionalRoles = []string{"loggedin"}
)
