package template

import (
	"html/template"
	"strings"

	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/messages"
	"github.com/goatcms/goatcore/varutil"
)

type DefaultFuncs struct {
	Translate services.Translate `dependency:"TranslateService"`
	Template  services.Template  `dependency:"TemplateService"`
	Logger    services.Logger    `dependency:"LoggerService"`
}

func NewDefaultFuncs(di dependency.Injector) (*DefaultFuncs, error) {
	df := &DefaultFuncs{}
	if err := di.InjectTo(df); err != nil {
		return nil, err
	}
	return df, nil
}

// Init inicjalize default teplate function
func (df *DefaultFuncs) Register() {
	df.Template.AddFunc(services.CutTextTF, df.CutText)
	df.Template.AddFunc(services.MessagesTF, df.Messages)
	df.Template.AddFunc("contains", varutil.IsArrContainStr)
	df.Template.AddFunc("dict", Dict)
}

func Dict(v ...interface{}) map[string]interface{} {
	dict := map[string]interface{}{}
	lenv := len(v)
	for i := 0; i < lenv; i += 2 {
		key := v[i].(string)
		if i+1 >= lenv {
			dict[key] = ""
			continue
		}
		dict[key] = v[i+1]
	}
	return dict
}

// CutText cut text to max length
func (dt *DefaultFuncs) CutText(max int, text string) string {
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
func (dt *DefaultFuncs) Messages(messages messages.MessageMap, langPrefix, class string, fields ...string) template.HTML {
	if messages == nil || len(messages.GetAll()) == 0 {
		return ""
	}
	dt.Logger.DevLog("template.DefaultFun.Messages render for %v %v %v %v", messages, langPrefix, class, fields)
	out := template.HTML("<div class=\"messages " + class + "\">")
	if fields == nil || len(fields) == 0 {
		for _, list := range messages.GetAll() {
			out += dt.messageLists(list, langPrefix, "")
		}
	} else {
		for _, fieldName := range fields {
			list := messages.Get(fieldName)
			out += dt.messageLists(list, langPrefix, "")
		}
	}
	return out + "</ul>"
}

func (dt *DefaultFuncs) messageLists(list messages.MessageList, langPrefix, class string) template.HTML {
	if list == nil || len(list.GetAll()) == 0 {
		return ""
	}
	out := template.HTML("<ul class=\"messages " + class + "\">")
	for _, msg := range list.GetAll() {
		tmsg, err := dt.Translate.Translate(langPrefix + msg)
		if err != nil {
			dt.Logger.DevLog("template.DefaultFun.messageLists error: %v", err)
		}
		out += template.HTML("<li>" + tmsg + "</li>")
	}
	return out + "</ul>"
}
