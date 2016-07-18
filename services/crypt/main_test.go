package crypt

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestIfHashAndGoodPassDoesMatch(t *testing.T) {
	const testPassword = "goodpassword"
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

func TestIfHashAndWrongPassDontMatch(t *testing.T) {
	cryptService := &Crypt{}
	hashed, err := cryptService.Hash("goodpassword")
	if err != nil {
		t.Error(err)
		return
	}
	if bcrypt.CompareHashAndPassword(
		[]byte(hashed),
		[]byte("wrongpassword"),
	) == nil {
		t.Error("comparing hash and WRONG password: they should not be identical")
	}
}

func TestCompareFuncWorksIfPassAndHashMatch(t *testing.T) {
	const testPassword = "goodpassword"
	cryptService := &Crypt{}
	hashed, err := cryptService.Hash(testPassword)
	if err != nil {
		t.Error(err)
		return
	}
	eval, err := cryptService.Compare(hashed, testPassword)
	if err != nil {
		t.Error(err)
	}
	if eval != true {
		t.Error("Compare func should be true ")
	}
}

func TestCompareFuncWorksIfPassAndHashDoesntMatch(t *testing.T) {
	cryptService := &Crypt{}
	hashed, err := cryptService.Hash("wrongpassword")
	if err != nil {
		t.Error(err)
		return
	}
	eval, err := cryptService.Compare(hashed, "goodpassword")
	if err == nil {
		t.Error(err)
	}
	if eval == true {
		t.Error("Compare func should not be true ")
	}
}
