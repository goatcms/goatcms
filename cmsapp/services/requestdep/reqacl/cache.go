package reqacl

import (
	"regexp"
	"strings"
)

// Cache is boject storage cached data
type Cache struct {
	// anonymousRoles contains roles for unauthorized (anonymous) users
	anonymousRoles []string
}

// NewCache create new Cache instance
func NewCache(anonRolesFromConfig string) (cache *Cache) {
	cache = &Cache{}
	// match aanonymous roles
	space := regexp.MustCompile(`\s+`)
	anonRolesFromConfig = space.ReplaceAllString(anonRolesFromConfig, " ")
	anonRolesFromConfig = strings.TrimSpace(anonRolesFromConfig)
	if anonRolesFromConfig != "" {
		cache.anonymousRoles = strings.Split(anonRolesFromConfig, " ")
		cache.anonymousRoles = append(cache.anonymousRoles, DefaultAnonymousRoles...)
	} else {
		cache.anonymousRoles = DefaultAnonymousRoles
	}
	return cache
}

// AnonRoles return roles for anonymous user
func (cache *Cache) AnonRoles() []string {
	return cache.anonymousRoles
}
