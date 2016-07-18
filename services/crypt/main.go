package crypt

import (
	"github.com/goatcms/goat-core/dependency"
	"golang.org/x/crypto/bcrypt"
)

// Crypt is global encrypting provider
type Crypt struct{}

// NewCrypt create a encrypting instance
func NewCrypt(dp dependency.Provider) (*Crypt, error) {
	return &Crypt{}, nil
}

// Hash take input string (f.e. password) and return bcrypted string
func (c *Crypt) Hash(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// Compare take pass and hashedPass and return boolean if it's match or not
func (c *Crypt) Compare(hashedPass, pass string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
	if err != nil {
		return false, err
	}
	return true, nil
}
