package forms

import (
	"github.com/goatcms/goatcore/messages"
	"github.com/goatcms/goatcore/messages/msgcollection"
)

func ValidSignin(e *Signin) (mm messages.MessageMap, err error) {
	mm = msgcollection.NewMessageMap()
	if err = AddSigninValid("", mm, e); err != nil {
		return nil, err
	}
	return mm, nil
}

func AddSigninValid(basekey string, mm messages.MessageMap, e *Signin) error {
	var messageKey string

	// Username field
	messageKey = basekey + "Username"
	if e.Username == nil {
		mm.Add(messageKey, "required")
	} else if len(*e.Username) == 0 {
		mm.Add(messageKey, "required")
	}

	// Password field
	messageKey = basekey + "Password"
	if e.Password == nil {
		mm.Add(messageKey, "required")
	} else if len(*e.Password) == 0 {
		mm.Add(messageKey, "required")
	}

	return nil
}
