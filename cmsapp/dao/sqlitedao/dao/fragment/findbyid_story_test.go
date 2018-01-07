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
		expectedEntity *entities.Fragment
		entity         *entities.Fragment
	)
	if ok, db, expectedEntity = doInsertStory(t); !ok {
		return false, nil
	}
	s := scope.NewScope("tag")
	finder := FragmentFindByID{}
	finder.deps.DB = db
	if entity, err = finder.Find(s, entities.FragmentAllFields, *expectedEntity.ID); err != nil {
		t.Error(err)
		return false, db
	}
	if *expectedEntity.Name != *entity.Name {
		t.Errorf("Returned field should contains inserted entity value for Name field and it is %v (expeted %v)", entity.Name, expectedEntity.Name)
		return false, db
	}
	if *expectedEntity.Lang != *entity.Lang {
		t.Errorf("Returned field should contains inserted entity value for Lang field and it is %v (expeted %v)", entity.Lang, expectedEntity.Lang)
		return false, db
	}
	if *expectedEntity.Content != *entity.Content {
		t.Errorf("Returned field should contains inserted entity value for Content field and it is %v (expeted %v)", entity.Content, expectedEntity.Content)
		return false, db
	}
	return true, db
}
