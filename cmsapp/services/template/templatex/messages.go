package templatex

import (
	"html/template"

	"github.com/goatcms/goatcore/messages"
)

// Messages view a messages list
func (module *Module) Messages(messages messages.MessageMap, langPrefix, class string, fields ...string) template.HTML {
	if messages == nil || len(messages.GetAll()) == 0 {
		return ""
	}
	module.deps.Logger.DevLog("template.DefaultFun.Messages render for %v %v %v %v", messages, langPrefix, class, fields)
	out := template.HTML("<div class=\"messages " + class + "\">")
	if fields == nil || len(fields) == 0 {
		for _, list := range messages.GetAll() {
			out += module.messageLists(list, langPrefix, "")
		}
	} else {
		for _, fieldName := range fields {
			list := messages.Get(fieldName)
			out += module.messageLists(list, langPrefix, "")
		}
	}
	return out + "</ul>"
}

func (module *Module) messageLists(list messages.MessageList, langPrefix, class string) template.HTML {
	if list == nil || len(list.GetAll()) == 0 {
		return ""
	}
	out := template.HTML("<ul class=\"messages " + class + "\">")
	for _, msg := range list.GetAll() {
		tmsg, err := module.deps.Translate.Translate(langPrefix + msg)
		if err != nil {
			module.deps.Logger.DevLog("template.DefaultFun.messageLists error: %v", err)
		}
		out += template.HTML("<li>" + tmsg + "</li>")
	}
	return out + "</ul>"
}
