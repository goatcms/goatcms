package reqtranslate

import (
	"net/http"
	"strings"

	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcms/cmsapp/services/requestdep"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/varutil"
)

// TranslateService provide translate system
type TranslateService struct {
	prefix    string
	translate services.Translate
}

func (ts *TranslateService) Translate(key string, values ...interface{}) (string, error) {
	return ts.translate.Translate(ts.prefix+key, values...)
}

func (ts *TranslateService) Lang() string {
	return ts.prefix[:len(ts.prefix)-1]
}

func isLangSepChar(r rune) bool {
	return r == ';' || r == ','
}

// TranslateFactory create new Translate provider
func TranslateFactory(dp dependency.Provider) (interface{}, error) {
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
	prefix = prefix + "."
	service := &TranslateService{
		translate: deps.Translate,
		prefix:    prefix,
	}
	return requestdep.Translate(service), nil
}
