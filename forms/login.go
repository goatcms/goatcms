package forms

import (
	"github.com/goatcms/goatcms/models"
	"github.com/goatcms/goatcms/services"
)

// LoginForm is structure with login form values
type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

const (
	errNoUser            = "Please give email for account you want to login"
	errWrongLoginAttempt = "E-mail and password you entered are not correct"
)

// Validate validate form and return bool how validation passed
func (r LoginForm) Validate(u models.UserDTO, c services.Crypt) (bool, map[string][]string) {
	validation := true
	errors := make(map[string][]string)
	// do validation
	if u == nil { // if no user with given email
		// log.Println("FAIL LOGIN: no user with given email")
		errors["email"] = []string{errWrongLoginAttempt}
		validation = false
	}
	if u != nil { // if there is user with given email
		passMatch, err2 := c.Compare(u.GetPassHash(), r.Password)
		if err2 != nil { // here error means: hash and pass are not matching
			// log.Println("FAIL LOGIN: password wrong")
			errors["email"] = []string{errWrongLoginAttempt}
			validation = false
		}
		if passMatch == true { // if error == nil and compare == true
			// log.Println("LOGIN OK: password correct")
			validation = true
		}
	}
	if r.Password == "" {
		errors["email"] = []string{errWrongLoginAttempt}
		validation = false
	}
	if r.Email == "" {
		errors["email"] = []string{errNoUser}
		validation = false
	}
	return validation, errors
}
