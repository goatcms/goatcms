package dao

import (
	"database/sql"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
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
		row            maindef.FragmentRow
		ok             bool
		db             *sql.DB
		err            error
		expectedEntity *entities.Fragment
		entity         *entities.Fragment
	)
	expectedEntity = NewMockEntity2()
	if ok, db, entity = doInsertStory(t); !ok {
		return false, nil
	}
	entity.Content = expectedEntity.Content
	entity.Name = expectedEntity.Name
	entity.Lang = expectedEntity.Lang
	s := scope.NewScope("tag")
	updater := FragmentUpdate{}
	updater.deps.DB = db
	if err = updater.Update(s, entity, entities.FragmentMainFields); err != nil {
		t.Error(err)
		return false, db
	}
	finder := FragmentFindByID{}
	finder.deps.DB = db
	if row, err = finder.Find(s, entities.FragmentMainFields, *entity.ID); err != nil {
		t.Error(err)
		return false, db
	}
	// iterate over each row
	if entity, err = row.Get(); err != nil {
		t.Error(err)
		return false, db
	}
	if *expectedEntity.Content != *entity.Content {
		t.Errorf("Returned field should contains inserted entity value for Content field and it is %v (expeted %v)", entity.Content, expectedEntity.Content)
		return false, db
	}
	if *expectedEntity.Lang != *entity.Lang {
		t.Errorf("Returned field should contains inserted entity value for Lang field and it is %v (expeted %v)", entity.Lang, expectedEntity.Lang)
		return false, db
	}
	if *expectedEntity.Name != *entity.Name {
		t.Errorf("Returned field should contains inserted entity value for Name field and it is %v (expeted %v)", entity.Name, expectedEntity.Name)
		return false, db
	}
	return true, db
}
