package templatex

import "strings"

// LengthLimit cut text to max length
func (module *Module) LengthLimit(max int, text string) string {
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
