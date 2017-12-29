package dao

import (
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlite/helpers"
	entities "github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestInsertStory(t *testing.T) {
	t.Parallel()
	doInsertStory(t)
}

func doInsertStory(t *testing.T) (bool, *sqlx.DB, *entities.Translation) {
	var (
		resultID int64
		ok       bool
		db       *sqlx.DB
		err      error
		entity   *entities.Translation
	)
	if ok, db = doCreateTable(t); !ok {
		return false, nil, nil
	}
	entity = NewMockEntity1()
	s := scope.NewScope("tag")
	persister := TranslationInsert{}
	persister.deps.DB = db
	if resultID, err = persister.Insert(s, entity, entities.TranslationMainFields); err != nil {
		t.Error(err)
		return false, db, entity
	}
	// expected set a new entity id
	if entity.ID == 0 {
		t.Errorf("the entity id should contains database ID and it is %v", entity.ID)
		return false, db, entity
	}
	if resultID == 0 {
		t.Errorf("id returned by fuction should contains database id and it is %v", resultID)
		return false, db, entity
	}
	if resultID != entity.ID {
		t.Errorf("id returned by fuction should and entity.ID are the same database id and must be equals. They are (%v != %v)", resultID, entity.ID)
		return false, db, entity
	}
	if _, err = helpers.Commit(s); err != nil {
		t.Error(err)
		return false, db, entity
	}
	return true, db, entity
}
