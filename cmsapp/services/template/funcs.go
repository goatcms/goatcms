package template

import (
	"strings"

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

// Init inicjalize default teplate function
func AddDefaultTemplateFunctions(t services.Template) {
	t.AddFunc(services.CutTextTF, CutText)
}
