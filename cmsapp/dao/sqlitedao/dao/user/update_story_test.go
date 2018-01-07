package dao

import (
	"database/sql"
	entities "github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app/scope"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestUpdateStory(t *testing.T) {
	t.Parallel()
	doUpdateStory(t)
}

func doUpdateStory(t *testing.T) (bool, *sql.DB) {
	var (
		ok             bool
		db             *sql.DB
		err            error
		expectedEntity *entities.User
		entity         *entities.User
	)
	expectedEntity = NewMockEntity2()
	if ok, db, entity = doInsertStory(t); !ok {
		return false, nil
	}
	entity.Firstname = expectedEntity.Firstname
	entity.Username = expectedEntity.Username
	entity.Roles = expectedEntity.Roles
	entity.Password = expectedEntity.Password
	entity.Lastname = expectedEntity.Lastname
	entity.Email = expectedEntity.Email
	s := scope.NewScope("tag")
	updater := UserUpdate{}
	updater.deps.DB = db
	if err = updater.Update(s, entity, entities.UserAllFields); err != nil {
		t.Error(err)
		return false, db
	}
	finder := UserFindByID{}
	finder.deps.DB = db
	if entity, err = finder.Find(s, entities.UserAllFields, *entity.ID); err != nil {
		t.Error(err)
		return false, db
	}
	if *expectedEntity.Email != *entity.Email {
		t.Errorf("Returned field should contains inserted entity value for Email field and it is %v (expeted %v)", entity.Email, expectedEntity.Email)
		return false, db
	}
	if *expectedEntity.Username != *entity.Username {
		t.Errorf("Returned field should contains inserted entity value for Username field and it is %v (expeted %v)", entity.Username, expectedEntity.Username)
		return false, db
	}
	if *expectedEntity.Roles != *entity.Roles {
		t.Errorf("Returned field should contains inserted entity value for Roles field and it is %v (expeted %v)", entity.Roles, expectedEntity.Roles)
		return false, db
	}
	if *expectedEntity.Password != *entity.Password {
		t.Errorf("Returned field should contains inserted entity value for Password field and it is %v (expeted %v)", entity.Password, expectedEntity.Password)
		return false, db
	}
	if *expectedEntity.Lastname != *entity.Lastname {
		t.Errorf("Returned field should contains inserted entity value for Lastname field and it is %v (expeted %v)", entity.Lastname, expectedEntity.Lastname)
		return false, db
	}
	if *expectedEntity.Firstname != *entity.Firstname {
		t.Errorf("Returned field should contains inserted entity value for Firstname field and it is %v (expeted %v)", entity.Firstname, expectedEntity.Firstname)
		return false, db
	}
	return true, db
}
