package dao

import (
	"database/sql"
	entities "github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app/scope"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestFindByIDStory(t *testing.T) {
	t.Parallel()
	doFindByIDStory(t)
}

func doFindByIDStory(t *testing.T) (bool, *sql.DB) {
	var (
		ok             bool
		db             *sql.DB
		err            error
		expectedEntity *entities.Session
		entity         *entities.Session
	)
	if ok, db, expectedEntity = doInsertWithoutIDStory(t); !ok {
		return false, nil
	}
	s := scope.NewScope("tag")
	finder := SessionFindByID{}
	finder.deps.DB = db
	if entity, err = finder.Find(s, entities.SessionAllFields, *expectedEntity.ID); err != nil {
		t.Error(err)
		return false, db
	}
	if *expectedEntity.Secret != *entity.Secret {
		t.Errorf("Returned field should contains inserted entity value for Secret field and it is %v (expeted %v)", entity.Secret, expectedEntity.Secret)
		return false, db
	}
	return true, db
}
