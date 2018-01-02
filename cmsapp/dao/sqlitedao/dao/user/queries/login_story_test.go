package queries

import (
	"bytes"
	"strings"
	"testing"

	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	dao "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/dao/user"
	database "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/database"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	entities "github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/mockupapp"
	"github.com/goatcms/goatcore/app/scope"
	_ "github.com/mattn/go-sqlite3"
)

func TestUserLoginQuery(t *testing.T) {
	t.Parallel()
	doUserLoginQuery(t)
}

func doUserLoginQuery(t *testing.T) bool {
	var (
		mapp           app.App
		err            error
		expectedEntity *entities.User
		entity         *entities.User
		deps           struct {
			TableCreator maindef.CreateTable    `dependency:"UserCreateTable"`
			Inserter     maindef.UserInsert     `dependency:"UserInsert"`
			Query        maindef.UserLoginQuery `dependency:"UserLoginQuery"`
		}
	)
	configScope := scope.NewScope(app.ConfigTagName)
	configScope.Set("database.url", ":memory:")
	if mapp, err = mockupapp.NewApp(mockupapp.MockupOptions{
		Input:       gio.NewInput(strings.NewReader("")),
		Output:      gio.NewOutput(new(bytes.Buffer)),
		ConfigScope: configScope,
	}); err != nil {
		t.Error(err)
		return false
	}
	if err = mapp.DependencyProvider().AddDefaultFactory("db0.engine", database.EngineFactory); err != nil {
		t.Error(err)
		return false
	}
	if err = dao.RegisterDependencies(mapp.DependencyProvider()); err != nil {
		t.Error(err)
		return false
	}
	if err = RegisterDependencies(mapp.DependencyProvider()); err != nil {
		t.Error(err)
		return false
	}
	if err = mapp.DependencyProvider().InjectTo(&deps); err != nil {
		t.Error(err)
		return false
	}
	s := scope.NewScope("testtag")
	if err := deps.TableCreator.CreateTable(s); err != nil {
		t.Error(err)
		return false
	}
	if _, err = helpers.Commit(s); err != nil {
		t.Error(err)
		return false
	}
	expectedEntity = NewMockEntity1()
	if _, err = deps.Inserter.Insert(s, expectedEntity); err != nil {
		t.Error(err)
		return false
	}
	if _, err = helpers.Commit(s); err != nil {
		t.Error(err)
		return false
	}
	params := &maindef.UserLoginQueryParams{
		Email:    *expectedEntity.Email,
		Password: *expectedEntity.Password,
		Login:    *expectedEntity.Login,
	}
	if entity, err = deps.Query.Login(s, entities.UserMainFields, params); err != nil {
		t.Error(err)
		return false
	}
	if *expectedEntity.Firstname != *entity.Firstname {
		t.Errorf("Returned field should contains inserted entity value for Firstname field and it is %v (expeted %v)", entity.Firstname, expectedEntity.Firstname)
		return false
	}
	if *expectedEntity.Login != *entity.Login {
		t.Errorf("Returned field should contains inserted entity value for Login field and it is %v (expeted %v)", entity.Login, expectedEntity.Login)
		return false
	}
	if *expectedEntity.Email != *entity.Email {
		t.Errorf("Returned field should contains inserted entity value for Email field and it is %v (expeted %v)", entity.Email, expectedEntity.Email)
		return false
	}
	if *expectedEntity.Password != *entity.Password {
		t.Errorf("Returned field should contains inserted entity value for Password field and it is %v (expeted %v)", entity.Password, expectedEntity.Password)
		return false
	}
	return true
}
