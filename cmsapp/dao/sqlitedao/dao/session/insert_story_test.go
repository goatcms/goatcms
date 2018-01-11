package dao

import (
	"database/sql"
	helpers "github.com/goatcms/goatcms/cmsapp/dao/sqlitedao/helpers"
	entities "github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app/scope"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestInsertWithoutIDStory(t *testing.T) {
	t.Parallel()
	doInsertWithoutIDStory(t)
}

func doInsertWithoutIDStory(t *testing.T) (bool, *sql.DB, *entities.Session) {
	var (
		resultID int64
		ok       bool
		db       *sql.DB
		err      error
		entity   *entities.Session
	)
	if ok, db = doCreateTable(t); !ok {
		return false, nil, nil
	}
	entity = NewMockEntity1()
	s := scope.NewScope("tag")
	persister := SessionInsert{}
	persister.deps.DB = db
	if resultID, err = persister.Insert(s, entity); err != nil {
		t.Error(err)
		return false, db, entity
	}
	// expected set a new entity id
	if *entity.ID == 0 {
		t.Errorf("the entity id should contains database ID and it is %v", entity.ID)
		return false, db, entity
	}
	if resultID == 0 {
		t.Errorf("id returned by fuction should contains database id and it is %v", resultID)
		return false, db, entity
	}
	if resultID != *entity.ID {
		t.Errorf("id returned by fuction should and entity.ID are the same database id and must be equals. They are (%v != %v)", resultID, entity.ID)
		return false, db, entity
	}
	if _, err = helpers.Commit(s); err != nil {
		t.Error(err)
		return false, db, entity
	}
	return true, db, entity
}

func TestInsertWithIDStory(t *testing.T) {
	t.Parallel()
	doInsertWithIDStory(t)
}

func doInsertWithIDStory(t *testing.T) (bool, *sql.DB, *entities.Session) {
	var (
		resultID   int64
		ok         bool
		db         *sql.DB
		err        error
		entity     *entities.Session
		expectedID int64 = 666
	)
	if ok, db = doCreateTable(t); !ok {
		return false, nil, nil
	}
	entity = NewMockEntity1()
	entity.ID = &expectedID
	s := scope.NewScope("tag")
	persister := SessionInsert{}
	persister.deps.DB = db
	if resultID, err = persister.Insert(s, entity); err != nil {
		t.Error(err)
		return false, db, entity
	}
	// expected set a new entity id
	if *entity.ID == 0 {
		t.Errorf("the entity id should contains database ID and it is %v", entity.ID)
		return false, db, entity
	}
	if resultID != expectedID {
		t.Errorf("id returned by fuction should contains database id and it is %v", resultID)
		return false, db, entity
	}
	if resultID != *entity.ID {
		t.Errorf("id returned by fuction should and entity.ID are the same database id and must be equals. They are (%v != %v)", resultID, entity.ID)
		return false, db, entity
	}
	if _, err = helpers.Commit(s); err != nil {
		t.Error(err)
		return false, db, entity
	}
	return true, db, entity
}
