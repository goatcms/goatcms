package forms

import "regexp"

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
)

const (
	passwordLength = 6
)

// Validate validate form and return bool how validation passed
func (r RegisterForm) Validate() (bool, []string) {
	validation := true
	errors := []string{}
	// do validation
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
