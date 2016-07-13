package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"

	"github.com/goatcms/goat-core/dependency"
)

// Auth is global auth provider
type Auth struct {
}

// NewAuth create a authentification service instance
func NewAuth(dp dependency.Provider) (*Auth, error) {
	return &Auth{}, nil
}

// GetCode create HMAC for given string
func (a *Auth) GetCode(data string) string {
	salt := "hereshouldbesomekey"
	h := hmac.New(sha256.New, []byte(salt))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
