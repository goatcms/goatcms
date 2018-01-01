package dao

import (
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	entities "github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestFindByIDStory(t *testing.T) {
	t.Parallel()
	doFindByIDStory(t)
}

func doFindByIDStory(t *testing.T) (bool, *sqlx.DB) {
	var (
		row            maindef.Row
		ok             bool
		db             *sqlx.DB
		err            error
		expectedEntity *entities.Translation
		entity         *entities.Translation
	)
	if ok, db, expectedEntity = doInsertStory(t); !ok {
		return false, nil
	}
	s := scope.NewScope("tag")
	finder := TranslationFindByID{}
	finder.deps.DB = db
	if row, err = finder.Find(s, entities.TranslationMainFields, expectedEntity.ID); err != nil {
		t.Error(err)
		return false, db
	}
	// iterate over each row
	entity = &entities.Translation{}
	if err = row.StructScan(entity); err != nil {
		t.Error(err)
		return false, db
	}
	if expectedEntity.Value != entity.Value {
		t.Errorf("Returned field should contains inserted entity value for Value field and it is %v (expeted %v)", entity.Value, expectedEntity.Value)
		return false, db
	}
	if expectedEntity.Key != entity.Key {
		t.Errorf("Returned field should contains inserted entity value for Key field and it is %v (expeted %v)", entity.Key, expectedEntity.Key)
		return false, db
	}
	return true, db
}
