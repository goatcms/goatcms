package dao

import (
	"bytes"
	"strings"
	"testing"

	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	database "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/database"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/mockupapp"
	"github.com/goatcms/goatcore/app/scope"
)

func TestModule(t *testing.T) {
	var (
		err  error
		mapp app.App
	)
	t.Parallel()
	// prepare mockup application
	configScope := scope.NewScope(app.ConfigTagName)
	configScope.Set("database.url", ":memory:")
	if mapp, err = mockupapp.NewApp(mockupapp.MockupOptions{
		Input:       gio.NewInput(strings.NewReader("")),
		Output:      gio.NewOutput(new(bytes.Buffer)),
		ConfigScope: configScope,
	}); err != nil {
		t.Error(err)
		return
	}
	if err = mapp.DependencyProvider().AddDefaultFactory("db0.engine", database.EngineFactory); err != nil {
		t.Error(err)
		return
	}
	if err = RegisterDependencies(mapp.DependencyProvider()); err != nil {
		t.Error(err)
		return
	}
	// test
	var deps struct {
		TranslationCreateTable maindef.CreateTable         `dependency:"TranslationCreateTable"`
		TranslationDelete      maindef.Delete              `dependency:"TranslationDelete"`
		TranslationDropTable   maindef.DropTable           `dependency:"TranslationDropTable"`
		TranslationFindAll     maindef.TranslationFindAll  `dependency:"TranslationFindAll"`
		TranslationFindByID    maindef.TranslationFindByID `dependency:"TranslationFindByID"`
		TranslationInsert      maindef.TranslationInsert   `dependency:"TranslationInsert"`
		TranslationUpdate      maindef.TranslationUpdate   `dependency:"TranslationUpdate"`
		TranslationSearch      maindef.TranslationSearch   `dependency:"TranslationSearch"`
	}
	if err = mapp.DependencyProvider().InjectTo(&deps); err != nil {
		t.Error(err)
		return
	}
}
