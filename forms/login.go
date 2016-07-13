package forms

import (
	"log"

	"github.com/goatcms/goatcms/models"
	"github.com/goatcms/goatcms/services"
)

// LoginForm is structure with login form values
type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

var (
	errNoSuchUser        = "There is no user with this email"
	errWrongLoginAttempt = "E-mail and password you entered are not correct"
)

// Validate validate form and return bool how validation passed
func (r LoginForm) Validate(u models.UserDTO, c services.Crypt) (bool, []string) {
	validation := true
	errors := []string{}
	// do validation
	if r.Email == "" {
		errors = append(errors, errNoSuchUser)
		validation = false
	}
	if u == nil { // if no user with given email
		log.Println("FAIL LOGIN: no user with given email")
		errors = []string{errWrongLoginAttempt}
		validation = false
	}
	if u != nil { // if there is user with given email
		passMatch, err2 := c.Compare(u.GetPassHash(), r.Password)
		if err2 != nil { // here error means: hash and pass are not matching
			log.Println("FAIL LOGIN: password wrong")
			errors = []string{errWrongLoginAttempt}
			validation = false
		}
		if passMatch == true { // if error == nil and compare == true
			log.Println("FAIL LOGIN: password correct")
			validation = true
		}
	}
	if r.Password == "" {
		errors = []string{errWrongLoginAttempt}
		validation = false
	}
	return validation, errors
}
