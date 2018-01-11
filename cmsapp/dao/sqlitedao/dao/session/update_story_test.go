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
		expectedEntity *entities.Session
		entity         *entities.Session
	)
	expectedEntity = NewMockEntity2()
	if ok, db, entity = doInsertWithoutIDStory(t); !ok {
		return false, nil
	}
	entity.Secret = expectedEntity.Secret
	s := scope.NewScope("tag")
	updater := SessionUpdate{}
	updater.deps.DB = db
	if err = updater.Update(s, entity, entities.SessionAllFields); err != nil {
		t.Error(err)
		return false, db
	}
	finder := SessionFindByID{}
	finder.deps.DB = db
	if entity, err = finder.Find(s, entities.SessionAllFields, *entity.ID); err != nil {
		t.Error(err)
		return false, db
	}
	if *expectedEntity.Secret != *entity.Secret {
		t.Errorf("Returned field should contains inserted entity value for Secret field and it is %v (expeted %v)", entity.Secret, expectedEntity.Secret)
		return false, db
	}
	return true, db
}
