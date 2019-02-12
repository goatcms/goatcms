package reqacl

import (
	"testing"
)

func TestCacheAnonymousRolesForEmpty(t *testing.T) {
	t.Parallel()
	cache := NewCache("")
	result := cache.AnonRoles()
	if len(result) != 1 {
		t.Errorf("require one result element")
	}
	if result[0] != "anonymous" {
		t.Errorf("require anonymous role for unknow user")
	}
}

func TestCacheAnonymousRolesWhitespaces(t *testing.T) {
	t.Parallel()
	cache := NewCache("   \t \t \n\n \t")
	result := cache.AnonRoles()
	if len(result) != 1 {
		t.Errorf("require one result element")
	}
	if result[0] != "anonymous" {
		t.Errorf("require anonymous role for unknow user")
	}
}

func TestCacheAnonymousCustomRoles(t *testing.T) {
	t.Parallel()
	cache := NewCache("   \t \t \n\nrole1 \t role2")
	result := cache.AnonRoles()
	if len(result) != 3 {
		t.Errorf("require one result element")
	}
	if result[0] != "role1" {
		t.Errorf("require role1")
	}
	if result[1] != "role2" {
		t.Errorf("require role2")
	}
	if result[2] != "anonymous" {
		t.Errorf("require anonymous role for unknow user")
	}
}
