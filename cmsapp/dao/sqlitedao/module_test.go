package sqlitedao

import (
	"bytes"
	"database/sql"
	"strings"
	"testing"

	maindef "github.com/goatcms/goatcms/cmsapp/dao"
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
	if err = RegisterDependencies(mapp.DependencyProvider()); err != nil {
		t.Error(err)
		return
	}
	// test
	var deps struct {
		DB                     *sql.DB                     `dependency:"db0.engine"`
		TranslationCreateTable maindef.CreateTable         `dependency:"TranslationCreateTable"`
		TranslationDelete      maindef.Delete              `dependency:"TranslationDelete"`
		TranslationDropTable   maindef.DropTable           `dependency:"TranslationDropTable"`
		TranslationFindAll     maindef.TranslationFindAll  `dependency:"TranslationFindAll"`
		TranslationFindByID    maindef.TranslationFindByID `dependency:"TranslationFindByID"`
		TranslationInsert      maindef.TranslationInsert   `dependency:"TranslationInsert"`
		TranslationUpdate      maindef.TranslationUpdate   `dependency:"TranslationUpdate"`
		UserCreateTable        maindef.CreateTable         `dependency:"UserCreateTable"`
		UserDelete             maindef.Delete              `dependency:"UserDelete"`
		UserDropTable          maindef.DropTable           `dependency:"UserDropTable"`
		UserFindAll            maindef.UserFindAll         `dependency:"UserFindAll"`
		UserFindByID           maindef.UserFindByID        `dependency:"UserFindByID"`
		UserInsert             maindef.UserInsert          `dependency:"UserInsert"`
		UserUpdate             maindef.UserUpdate          `dependency:"UserUpdate"`
		ArticleCreateTable     maindef.CreateTable         `dependency:"ArticleCreateTable"`
		ArticleDelete          maindef.Delete              `dependency:"ArticleDelete"`
		ArticleDropTable       maindef.DropTable           `dependency:"ArticleDropTable"`
		ArticleFindAll         maindef.ArticleFindAll      `dependency:"ArticleFindAll"`
		ArticleFindByID        maindef.ArticleFindByID     `dependency:"ArticleFindByID"`
		ArticleInsert          maindef.ArticleInsert       `dependency:"ArticleInsert"`
		ArticleUpdate          maindef.ArticleUpdate       `dependency:"ArticleUpdate"`
	}
	if err = mapp.DependencyProvider().InjectTo(&deps); err != nil {
		t.Error(err)
		return
	}
}
