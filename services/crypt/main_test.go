package crypt

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestIfHashAndPassAreEqual(t *testing.T) {
	const testPassword = "foobar"
	cryptService := &Crypt{}
	hashed, err := cryptService.Hash(testPassword)
	if err != nil {
		t.Error(err)
		return
	}
	if bcrypt.CompareHashAndPassword(
		[]byte(hashed),
		[]byte(testPassword),
	) != nil {
		t.Error("comparing hash and CORRECT password: they should be identical")
	}
}

func TestIfHashAndPassNotMatchingAreNotEqual(t *testing.T) {
	const testPassword = "foobar"
	cryptService := &Crypt{}
	hashed, err := cryptService.Hash(testPassword)
	if err != nil {
		t.Error(err)
		return
	}
	if bcrypt.CompareHashAndPassword(
		[]byte(hashed),
		[]byte("foobaz"),
	) == nil {
		t.Error("comparing hash and WRONG password: they should not be identical")
	}
}
