package articleform

import (
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goat-core/messages"
	"github.com/goatcms/goat-core/varutil/validator"
	"github.com/goatcms/goatcms/cmsapp/models"
)

// ArticleForm is structure with register form values
type ArticleForm models.Article

func NewForm(dp dependency.Injector) (*ArticleForm, error) {
	form := &ArticleForm{}
	if err := dp.InjectTo(form); err != nil {
		return nil, err
	}
	return form, nil
}

func (f *ArticleForm) Valid(basekey string, mm messages.MessageMap) error {
	if err := validator.MinStringValid(f.Title, basekey+"Title", mm, 1); err != nil {
		return err
	}
	if err := validator.MinStringValid(f.Content, basekey+"Content", mm, 1); err != nil {
		return err
	}
	return nil
}
