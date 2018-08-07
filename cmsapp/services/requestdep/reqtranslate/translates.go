package reqtranslate

import (
	"net/http"
	"strings"

	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/varutil"
)

// TranslateService provide translate service
type TranslateService struct {
	prefix    string
	translate services.Translate
}

// Translate a key for current lang
func (ts *TranslateService) Translate(key string, values ...interface{}) (string, error) {
	return ts.translate.Translate(ts.prefix+key, values...)
}

// Lang return current lang id
func (ts *TranslateService) Lang() string {
	return ts.prefix[:len(ts.prefix)-1]
}

// isLangSepChar check if character is lang separate character
// (different browser can use different separators)
func isLangSepChar(r rune) bool {
	return r == ';' || r == ','
}

// Factory create new Translate service instance
func Factory(dp dependency.Provider) (interface{}, error) {
	var deps struct {
		Translate services.Translate `dependency:"TranslateService"`
		Logger    services.Logger    `dependency:"LoggerService"`
		Request   *http.Request      `request:"Request"`
	}
	if err := dp.InjectTo(&deps); err != nil {
		return nil, err
	}
	prefix := deps.Translate.Default()
	langs := deps.Translate.Langs()
	acceptLanguage := deps.Request.Header.Get("Accept-Language")
	accepts := strings.FieldsFunc(acceptLanguage, isLangSepChar)
	for _, name := range accepts {
		if varutil.IsArrContainStr(langs, name) {
			prefix = name
			break
		}
	}
	deps.Logger.TestLog("reqtranslate.Factory: Current language is %v", prefix)
	prefix = prefix + "."
	service := &TranslateService{
		translate: deps.Translate,
		prefix:    prefix,
	}
	return requestdep.Translate(service), nil
}
