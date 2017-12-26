package translationdao

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
		expectedEntity *entities.Translation
		entity         *entities.Translation
	)
	expectedEntity = NewMockEntity2()
	if ok, db, entity = doInsertStory(t); !ok {
		return false, nil
	}
	entity.Key = expectedEntity.Key
	entity.Value = expectedEntity.Value
	s := scope.NewScope("tag")
	updater := TranslationUpdate{}
	updater.deps.DB = db
	if err = updater.Update(s, entity, entities.TranslationMainFields); err != nil {
		t.Error(err)
		return false, db
	}
	finder := TranslationFindByID{}
	finder.deps.DB = db
	if row, err = finder.Find(s, entities.TranslationMainFields, entity.ID); err != nil {
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
