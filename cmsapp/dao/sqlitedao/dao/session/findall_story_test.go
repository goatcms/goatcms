package dao

import (
	"database/sql"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	entities "github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app/scope"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestFindAllStory(t *testing.T) {
	t.Parallel()
	doFindAllStory(t)
}

func doFindAllStory(t *testing.T) (bool, *sql.DB) {
	var (
		rows maindef.SessionRows
		ok   bool
		db   *sql.DB
		err  error
	)
	if ok, db, _ = doInsertWithoutIDStory(t); !ok {
		return false, nil
	}
	s := scope.NewScope("tag")
	finder := SessionFindAll{}
	finder.deps.DB = db
	if rows, err = finder.Find(s, entities.SessionAllFields); err != nil {
		t.Error(err)
		return false, db
	}
	// iterate over each row
	count := 0
	expectedEntity := NewMockEntity1()
	for rows.Next() {
		var e *entities.Session
		count++
		if e, err = rows.Get(); err != nil {
			t.Error(err)
			return false, db
		}
		if *expectedEntity.Secret != *e.Secret {
			t.Errorf("Returned field should contains inserted entity value for Secret field and it is %v (expeted %v)", e.Secret, expectedEntity.Secret)
			return false, db
		}
	}
	if count != 1 {
		t.Errorf("FindAll should return one result and it return %v results", count)
		return false, db
	}
	return true, db
}
