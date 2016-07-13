package auth

import (
	"fmt"
	"testing"
)

func TestCreateNewCode(t *testing.T) {
	authService := &Auth{}
	hmac := authService.GetCode("somedata")
	fmt.Println(hmac)
	if hmac == "" {
		t.Error("Should create some HMAC")
	}
}
