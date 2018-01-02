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
		ArticleCreateTable maindef.CreateTable     `dependency:"ArticleCreateTable"`
		ArticleDelete      maindef.Delete          `dependency:"ArticleDelete"`
		ArticleDropTable   maindef.DropTable       `dependency:"ArticleDropTable"`
		ArticleFindAll     maindef.ArticleFindAll  `dependency:"ArticleFindAll"`
		ArticleFindByID    maindef.ArticleFindByID `dependency:"ArticleFindByID"`
		ArticleInsert      maindef.ArticleInsert   `dependency:"ArticleInsert"`
		ArticleUpdate      maindef.ArticleUpdate   `dependency:"ArticleUpdate"`
		ArticleSearch      maindef.ArticleSearch   `dependency:"ArticleSearch"`
	}
	if err = mapp.DependencyProvider().InjectTo(&deps); err != nil {
		t.Error(err)
		return
	}
}
