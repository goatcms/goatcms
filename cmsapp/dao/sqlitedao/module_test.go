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
		DB                  *sql.DB                  `dependency:"db0.engine"`
		FragmentCreateTable maindef.CreateTable      `dependency:"FragmentCreateTable"`
		FragmentDelete      maindef.Delete           `dependency:"FragmentDelete"`
		FragmentDropTable   maindef.DropTable        `dependency:"FragmentDropTable"`
		FragmentFindAll     maindef.FragmentFindAll  `dependency:"FragmentFindAll"`
		FragmentFindByID    maindef.FragmentFindByID `dependency:"FragmentFindByID"`
		FragmentInsert      maindef.FragmentInsert   `dependency:"FragmentInsert"`
		FragmentUpdate      maindef.FragmentUpdate   `dependency:"FragmentUpdate"`
		SessionCreateTable  maindef.CreateTable      `dependency:"SessionCreateTable"`
		SessionDelete       maindef.Delete           `dependency:"SessionDelete"`
		SessionDropTable    maindef.DropTable        `dependency:"SessionDropTable"`
		SessionFindAll      maindef.SessionFindAll   `dependency:"SessionFindAll"`
		SessionFindByID     maindef.SessionFindByID  `dependency:"SessionFindByID"`
		SessionInsert       maindef.SessionInsert    `dependency:"SessionInsert"`
		SessionUpdate       maindef.SessionUpdate    `dependency:"SessionUpdate"`
		UserCreateTable     maindef.CreateTable      `dependency:"UserCreateTable"`
		UserDelete          maindef.Delete           `dependency:"UserDelete"`
		UserDropTable       maindef.DropTable        `dependency:"UserDropTable"`
		UserFindAll         maindef.UserFindAll      `dependency:"UserFindAll"`
		UserFindByID        maindef.UserFindByID     `dependency:"UserFindByID"`
		UserInsert          maindef.UserInsert       `dependency:"UserInsert"`
		UserUpdate          maindef.UserUpdate       `dependency:"UserUpdate"`
	}
	if err = mapp.DependencyProvider().InjectTo(&deps); err != nil {
		t.Error(err)
		return
	}
}
