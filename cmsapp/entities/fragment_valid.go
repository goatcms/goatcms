package entities

import (
	"github.com/goatcms/goatcore/messages"
	"github.com/goatcms/goatcore/messages/msgcollection"
	"github.com/goatcms/goatcore/varutil/validator"
)

func ValidFragment(e *Fragment) (mm messages.MessageMap, err error) {
	mm = msgcollection.NewMessageMap()
	if err = AddFragmentValid("", mm, e); err != nil {
		return nil, err
	}
	return mm, nil
}

func AddFragmentValid(basekey string, mm messages.MessageMap, e *Fragment) error {
	var messageKey string

	// Name field
	messageKey = basekey + "Name"
	if e.Name == nil {
		mm.Add(messageKey, "required")
	} else if len(*e.Name) == 0 {
		mm.Add(messageKey, "required")
	}

	// Content field
	messageKey = basekey + "Content"
	if e.Content == nil {
		mm.Add(messageKey, "required")
	} else if len(*e.Content) == 0 {
		mm.Add(messageKey, "required")
	}

	// Lang field
	messageKey = basekey + "Lang"
	if e.Lang == nil {
		mm.Add(messageKey, "required")
	} else if len(*e.Lang) == 0 {
		mm.Add(messageKey, "required")
	}
	if e.Lang != nil {
		if err := validator.MinStringValid(*e.Lang, basekey+"Lang", mm, 2); err != nil {
			return err
		}
		if err := validator.MaxStringValid(*e.Lang, basekey+"Lang", mm, 3); err != nil {
			return err
		}
	}

	return nil
}
