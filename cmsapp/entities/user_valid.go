package entities

import (
	"github.com/goatcms/goatcore/messages"
	"github.com/goatcms/goatcore/messages/msgcollection"
	"github.com/goatcms/goatcore/varutil/validator"
)

func ValidUser(e *User) (mm messages.MessageMap, err error) {
	mm = msgcollection.NewMessageMap()
	if err = AddUserValid("", mm, e); err != nil {
		return nil, err
	}
	return mm, nil
}

func AddUserValid(basekey string, mm messages.MessageMap, e *User) error {
	var messageKey string

	// Username field
	messageKey = basekey + "Username"
	if e.Username == nil {
		mm.Add(messageKey, "required")
	} else if len(*e.Username) == 0 {
		mm.Add(messageKey, "required")
	}

	// Firstname field
	messageKey = basekey + "Firstname"
	if e.Firstname == nil {
		mm.Add(messageKey, "required")
	} else if len(*e.Firstname) == 0 {
		mm.Add(messageKey, "required")
	}

	// Roles field
	messageKey = basekey + "Roles"
	if e.Roles == nil {
		mm.Add(messageKey, "required")
	} else if len(*e.Roles) == 0 {
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

	// Password field
	messageKey = basekey + "Password"
	if e.Password == nil {
		mm.Add(messageKey, "required")
	} else if len(*e.Password) == 0 {
		mm.Add(messageKey, "required")
	}

	// Lastname field
	messageKey = basekey + "Lastname"
	if e.Lastname == nil {
		mm.Add(messageKey, "required")
	} else if len(*e.Lastname) == 0 {
		mm.Add(messageKey, "required")
	}

	return nil
}
