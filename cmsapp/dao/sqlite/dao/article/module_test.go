package articledao

import (
	"bytes"
	"strings"
	"testing"

	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	database "github.com/goatcms/goatcms/cmsapp/dao/sqlite/database"
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
	if err = mapp.DependencyProvider().AddDefaultFactory("sqlitedb0", database.Factory); err != nil {
		t.Error(err)
		return
	}
	if err = RegisterDependencies(mapp.DependencyProvider()); err != nil {
		t.Error(err)
		return
	}
	// test
	var deps struct {
		ArticleCreateTable maindef.CreateTable `dependency:"ArticleCreateTable"`
		ArticleDelete      maindef.Delete      `dependency:"ArticleDelete"`
		ArticleDropTable   maindef.DropTable   `dependency:"ArticleDropTable"`
		ArticleFindAll     maindef.FindAll     `dependency:"ArticleFindAll"`
		ArticleFindByID    maindef.FindByID    `dependency:"ArticleFindByID"`
		ArticleInsert      maindef.Insert      `dependency:"ArticleInsert"`
		ArticleUpdate      maindef.Update      `dependency:"ArticleUpdate"`
	}
	if err = mapp.DependencyProvider().InjectTo(&deps); err != nil {
		t.Error(err)
		return
	}
}
