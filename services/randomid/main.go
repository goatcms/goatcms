package randomid

import (
	"crypto/rand"
	"fmt"

	"github.com/goatcms/goat-core/dependency"
)

// RandomID is global random id provider
type RandomID struct{}

// NewRandomID create a random id instance
func NewRandomID(dp dependency.Provider) (*RandomID, error) {
	return &RandomID{}, nil
}

// Source string used to generate a random id from
const idSource = "0123456789" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"abcdefghijklmnopqrstuvwxyz"

// Save length in a constant so we don't have to look it up each time
const idSourceLen = byte(len(idSource))

// GenerateID creates a prefixed random id
func (r *RandomID) GenerateID(prefix string, length int) (string, error) {
	// Create an array with the correct capacity
	id := make([]byte, length)
	// Fill our array with random numbers
	rand.Read(id)
	// Replace each random number with alphanumeric value
	for i, b := range id {
		id[i] = idSource[b%idSourceLen]
	}
	// Returns the formatted id
	return fmt.Sprintf("%s_%s", prefix, string(id)), nil
}
