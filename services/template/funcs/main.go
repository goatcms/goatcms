package funcs

import (
	"html/template"
	"strings"

	"github.com/goatcms/goatcms/services"
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

// Init inicjalize default teplate function
func Init(t services.Template) {
	t.Funcs(template.FuncMap{
		"CutText": CutText,
	})
}
