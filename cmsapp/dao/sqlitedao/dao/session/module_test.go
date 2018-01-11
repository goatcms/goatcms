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
		SessionCreateTable maindef.CreateTable     `dependency:"SessionCreateTable"`
		SessionDelete      maindef.Delete          `dependency:"SessionDelete"`
		SessionDropTable   maindef.DropTable       `dependency:"SessionDropTable"`
		SessionFindAll     maindef.SessionFindAll  `dependency:"SessionFindAll"`
		SessionFindByID    maindef.SessionFindByID `dependency:"SessionFindByID"`
		SessionInsert      maindef.SessionInsert   `dependency:"SessionInsert"`
		SessionUpdate      maindef.SessionUpdate   `dependency:"SessionUpdate"`
		SessionSearch      maindef.SessionSearch   `dependency:"SessionSearch"`
	}
	if err = mapp.DependencyProvider().InjectTo(&deps); err != nil {
		t.Error(err)
		return
	}
}
