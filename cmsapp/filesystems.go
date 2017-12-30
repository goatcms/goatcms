package cmsapp

import "github.com/goatcms/goatcore/app"

func RegisterFilesystems(a app.App) error {
	root := a.RootFilespace()
	// templates
	templateFS, err := root.Filespace(TemplatePath)
	if err != nil {
		return err
	}
	a.FilespaceScope().Set(TemplateFilespace, templateFS)
	// translates
	translateFS, err := root.Filespace(TranslatePath)
	if err != nil {
		return err
	}
	a.FilespaceScope().Set(TranslateFilespace, translateFS)
	return nil
}
