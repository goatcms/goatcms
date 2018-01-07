package forms

import (
	"github.com/goatcms/goatcore/messages"
	"github.com/goatcms/goatcore/messages/msgcollection"
	"github.com/goatcms/goatcore/varutil/validator"
)

func ValidSignup(e *Signup) (mm messages.MessageMap, err error) {
	mm = msgcollection.NewMessageMap()
	if err = AddSignupValid("", mm, e); err != nil {
		return nil, err
	}
	return mm, nil
}

func AddSignupValid(basekey string, mm messages.MessageMap, e *Signup) error {
	var messageKey string

	// Lastname field
	messageKey = basekey + "Lastname"
	if e.Lastname == nil {
		mm.Add(messageKey, "required")
	} else if len(*e.Lastname) == 0 {
		mm.Add(messageKey, "required")
	}

	// Password field
	messageKey = basekey + "Password"
	if e.Password == nil {
		mm.Add(messageKey, "required")
	}
	if e.Password != nil {
		if err := validator.MinStringValid(e.Password.First, basekey+"Password", mm, 8); err != nil {
			return err
		}
		if err := validator.MaxStringValid(e.Password.First, basekey+"Password", mm, 255); err != nil {
			return err
		}
		if e.Password.First != e.Password.Second {
			mm.Add(messageKey, "identical_password")
		}
	}

	// Firstname field
	messageKey = basekey + "Firstname"
	if e.Firstname == nil {
		mm.Add(messageKey, "required")
	} else if len(*e.Firstname) == 0 {
		mm.Add(messageKey, "required")
	}

	// Username field
	messageKey = basekey + "Username"
	if e.Username == nil {
		mm.Add(messageKey, "required")
	} else if len(*e.Username) == 0 {
		mm.Add(messageKey, "required")
	}

	// Email field
	messageKey = basekey + "Email"
	if e.Email == nil {
		mm.Add(messageKey, "required")
	} else if len(*e.Email) == 0 {
		mm.Add(messageKey, "required")
	}
	if e.Email != nil {
		if err := validator.EmailValid(*e.Email, messageKey, mm); err != nil {
			return err
		}
	}

	return nil
}
