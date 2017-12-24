package userdao

import (
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	entities "github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestUpdateStory(t *testing.T) {
	t.Parallel()
	doUpdateStory(t)
}

func doUpdateStory(t *testing.T) (bool, *sqlx.DB) {
	var (
		row            maindef.Row
		ok             bool
		db             *sqlx.DB
		err            error
		expectedEntity *entities.User
		entity         *entities.User
	)
	expectedEntity = NewMockEntity2()
	if ok, db, entity = doInsertStory(t); !ok {
		return false, nil
	}
	entity.Lastname = expectedEntity.Lastname
	entity.Firstname = expectedEntity.Firstname
	s := scope.NewScope("tag")
	updater := UserUpdate{}
	updater.deps.DB = db
	if err = updater.Update(s, entity, entities.UserMainFields); err != nil {
		t.Error(err)
		return false, db
	}
	finder := UserFindByID{}
	finder.deps.DB = db
	if row, err = finder.Find(s, entities.UserMainFields, entity.ID); err != nil {
		t.Error(err)
		return false, db
	}
	// iterate over each row
	entity = &entities.User{}
	if err = row.StructScan(entity); err != nil {
		t.Error(err)
		return false, db
	}
	if expectedEntity.Lastname != entity.Lastname {
		t.Errorf("Returned field should contains inserted entity value for Lastname field and it is %v (expeted %v)", entity.Lastname, expectedEntity.Lastname)
		return false, db
	}
	if expectedEntity.Firstname != entity.Firstname {
		t.Errorf("Returned field should contains inserted entity value for Firstname field and it is %v (expeted %v)", entity.Firstname, expectedEntity.Firstname)
		return false, db
	}
	return true, db
}
