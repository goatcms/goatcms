package forms

import (
	"regexp"

	"github.com/goatcms/goatcms/models"
)

// RegisterForm is structure with register form values
type RegisterForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

var (
	errNoEmail        = "You must supply an email"
	errNoPassword     = "You must supply a password"
	errBadEmailFormat = "Email is not correct"
	errPassTooShort   = "Password should be at least 6 characters"
	errEmailOccupied  = "There is already a user with given email"
)

const (
	passwordLength = 6
)

// Validate validate form and return bool how validation passed
func (r RegisterForm) Validate(u models.UserDTO) (bool, []string) {
	validation := true
	errors := []string{}
	// do validation
	if u != nil {
		errors = append(errors, errEmailOccupied)
		validation = false
	}
	if r.Email == "" {
		errors = append(errors, errNoEmail)
		validation = false
	}
	if validateEmail(r.Email) == false {
		errors = append(errors, errBadEmailFormat)
		validation = false
	}
	if r.Password == "" {
		errors = append(errors, errNoPassword)
		validation = false
	}
	if len(r.Password) < passwordLength {
		errors = append(errors, errPassTooShort)
		validation = false
	}
	return validation, errors
}

func validateEmail(email string) bool {
	result := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return result.MatchString(email)
}
