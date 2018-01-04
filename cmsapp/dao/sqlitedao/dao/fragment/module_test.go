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
		FragmentCreateTable maindef.CreateTable      `dependency:"FragmentCreateTable"`
		FragmentDelete      maindef.Delete           `dependency:"FragmentDelete"`
		FragmentDropTable   maindef.DropTable        `dependency:"FragmentDropTable"`
		FragmentFindAll     maindef.FragmentFindAll  `dependency:"FragmentFindAll"`
		FragmentFindByID    maindef.FragmentFindByID `dependency:"FragmentFindByID"`
		FragmentInsert      maindef.FragmentInsert   `dependency:"FragmentInsert"`
		FragmentUpdate      maindef.FragmentUpdate   `dependency:"FragmentUpdate"`
		FragmentSearch      maindef.FragmentSearch   `dependency:"FragmentSearch"`
	}
	if err = mapp.DependencyProvider().InjectTo(&deps); err != nil {
		t.Error(err)
		return
	}
}
