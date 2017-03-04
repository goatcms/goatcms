package translate

import (
	"regexp"
	"strings"

	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/i18n/fsi18loader"
	"github.com/goatcms/goatcore/i18n/i18mem"
)

// TranslateService provide translate system
type TranslateService struct {
	langs       []string
	defaultLang string
	i18         *i18mem.I18Mem
}

func (ts *TranslateService) Langs() []string {
	return ts.langs
}

func (ts *TranslateService) Default() string {
	return ts.defaultLang
}

func (ts *TranslateService) Translate(key string, values ...interface{}) (string, error) {
	return ts.i18.Translate(key, values...)
}

func (ts *TranslateService) TranslateFor(key, prefix string, values ...interface{}) (string, error) {
	return ts.i18.Translate(prefix+key, values...)
}

// TranslateFactory create new Translate provider
func TranslateFactory(dp dependency.Provider) (interface{}, error) {
	var deps struct {
		Filespace   filesystem.Filespace `filespace:"translate"`
		EngineScope app.Scope            `dependency:"EngineScope"`
		Logger      services.Logger      `dependency:"LoggerService"`
		Langs       string               `config:"translate.langs"`
		Default     string               `config:"translate.default"`
	}
	if err := dp.InjectTo(&deps); err != nil {
		return nil, err
	}
	re := regexp.MustCompile("[^a-zA-Z,]*")
	deps.Langs = re.ReplaceAllString(deps.Langs, "")
	langs := strings.Split(deps.Langs, ",")
	deps.Logger.DevLog("tranlsate.TranslateFactory Lang list: %v", langs)
	if err := dp.InjectTo(&deps); err != nil {
		return nil, err
	}
	i18 := i18mem.NewI18Mem()
	if err := fsi18loader.Load(deps.Filespace, "./", i18, deps.EngineScope); err != nil {
		return nil, err
	}
	return services.Translate(&TranslateService{
		i18:         i18,
		defaultLang: deps.Default,
		langs:       langs,
	}), nil
}
