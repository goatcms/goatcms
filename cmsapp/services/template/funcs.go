package template

import (
	"html/template"
	"strings"

	"github.com/goatcms/goatcore/messages"
	"github.com/goatcms/goatcms/cmsapp/services"
)

// CutText cut text to max length
func CutText(max int, text string) string {
	text = strings.Trim(text, " \t")
	if len(text) < max {
		return text
	}
	cuter := text[0:max]
	last := strings.LastIndexAny(cuter, ". ")
	if last > 0 {
		return cuter[0:last-1] + "..."
	}
	return cuter + "..."
}

// Messages view a messages list
func Messages(messages messages.MessageMap, class string, fields ...string) template.HTML {
	if messages == nil || len(messages.GetAll()) == 0 {
		return ""
	}
	out := template.HTML("<div class=\"messages " + class + "\">")
	if fields == nil || len(fields) == 0 {
		for _, list := range messages.GetAll() {
			out += messageLists(list, "")
		}
	} else {
		for _, fieldName := range fields {
			list := messages.Get(fieldName)
			out += messageLists(list, "")
		}
	}
	return out + "</ul>"
}

func messageLists(list messages.MessageList, class string) template.HTML {
	if list == nil || len(list.GetAll()) == 0 {
		return ""
	}
	out := template.HTML("<ul class=\"messages " + class + "\">")
	for _, msg := range list.GetAll() {
		out += template.HTML("<li>" + msg + "</li>")
	}
	return out + "</ul>"
}

// Init inicjalize default teplate function
func AddDefaultTemplateFunctions(t services.Template) {
	t.AddFunc(services.CutTextTF, CutText)
	t.AddFunc(services.MessagesTF, Messages)
}
