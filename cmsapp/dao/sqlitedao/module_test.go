package sqlitedao

import (
	"bytes"
	"strings"
	"testing"

	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/mockupapp"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/jmoiron/sqlx"
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
	if err = RegisterDependencies(mapp.DependencyProvider()); err != nil {
		t.Error(err)
		return
	}
	// test
	var deps struct {
		DB                     *sqlx.DB            `dependency:"db0.engine"`
		UserCreateTable        maindef.CreateTable `dependency:"UserCreateTable"`
		UserDelete             maindef.Delete      `dependency:"UserDelete"`
		UserDropTable          maindef.DropTable   `dependency:"UserDropTable"`
		UserFindAll            maindef.FindAll     `dependency:"UserFindAll"`
		UserFindByID           maindef.FindByID    `dependency:"UserFindByID"`
		UserInsert             maindef.Insert      `dependency:"UserInsert"`
		UserUpdate             maindef.Update      `dependency:"UserUpdate"`
		TranslationCreateTable maindef.CreateTable `dependency:"TranslationCreateTable"`
		TranslationDelete      maindef.Delete      `dependency:"TranslationDelete"`
		TranslationDropTable   maindef.DropTable   `dependency:"TranslationDropTable"`
		TranslationFindAll     maindef.FindAll     `dependency:"TranslationFindAll"`
		TranslationFindByID    maindef.FindByID    `dependency:"TranslationFindByID"`
		TranslationInsert      maindef.Insert      `dependency:"TranslationInsert"`
		TranslationUpdate      maindef.Update      `dependency:"TranslationUpdate"`
		ArticleCreateTable     maindef.CreateTable `dependency:"ArticleCreateTable"`
		ArticleDelete          maindef.Delete      `dependency:"ArticleDelete"`
		ArticleDropTable       maindef.DropTable   `dependency:"ArticleDropTable"`
		ArticleFindAll         maindef.FindAll     `dependency:"ArticleFindAll"`
		ArticleFindByID        maindef.FindByID    `dependency:"ArticleFindByID"`
		ArticleInsert          maindef.Insert      `dependency:"ArticleInsert"`
		ArticleUpdate          maindef.Update      `dependency:"ArticleUpdate"`
	}
	if err = mapp.DependencyProvider().InjectTo(&deps); err != nil {
		t.Error(err)
		return
	}
}
