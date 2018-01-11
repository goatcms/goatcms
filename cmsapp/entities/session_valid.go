package entities

import (
	"github.com/goatcms/goatcore/messages"
	"github.com/goatcms/goatcore/messages/msgcollection"
	"github.com/goatcms/goatcore/varutil/validator"
)

func ValidSession(e *Session) (mm messages.MessageMap, err error) {
	mm = msgcollection.NewMessageMap()
	if err = AddSessionValid("", mm, e); err != nil {
		return nil, err
	}
	return mm, nil
}

func AddSessionValid(basekey string, mm messages.MessageMap, e *Session) error {
	var messageKey string

	// Secret field
	messageKey = basekey + "Secret"
	if e.Secret == nil {
		mm.Add(messageKey, validator.FieldIsRequired)
	} else if len(*e.Secret) == 0 {
		mm.Add(messageKey, validator.FieldIsRequired)
	}

	// User relation field
	messageKey = basekey + "User"
	if e.User == nil {
		mm.Add(messageKey, validator.FieldIsRequired)
	}

	return nil
}
