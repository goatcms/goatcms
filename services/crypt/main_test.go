package crypt

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestIfHashAndPassIsIdentical(t *testing.T) {
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
		t.Error("comparing hash and CORRECT pass should be true")
	}
}

func TestIfHashAndPassNotMatchingIsNotIdentical(t *testing.T) {
	cryptService := &Crypt{}
	hashed, err := cryptService.Hash("foobar")
	if err != nil {
		t.Error(err)
		return
	}
	if bcrypt.CompareHashAndPassword(
		[]byte(hashed),
		[]byte("foobaz"),
	) == nil {
		t.Error("comparing hash and WRONG pass pass should not be identical")
	}
}
