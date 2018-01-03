package dao

import (
	"database/sql"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	entities "github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app/scope"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestSearchStory(t *testing.T) {
	t.Parallel()
	doSearchStory(t)
}

func doSearchStory(t *testing.T) (bool, *sql.DB) {
	var (
		rows           maindef.TranslationRows
		ok             bool
		db             *sql.DB
		err            error
		expectedEntity *entities.Translation
	)
	if ok, db, expectedEntity = doInsertStory(t); !ok {
		return false, nil
	}
	s := scope.NewScope("tag")
	searcher := TranslationSearch{}
	searcher.deps.DB = db
	if rows, err = searcher.Search(s, entities.TranslationMainFields, &maindef.TranslationSearchParams{
		Value: *expectedEntity.Value,
		Key:   *expectedEntity.Key,
	}); err != nil {
		t.Error(err)
		return false, db
	}
	// iterate over each row
	count := 0
	for rows.Next() {
		var e *entities.Translation
		count++
		if e, err = rows.Get(); err != nil {
			t.Error(err)
			return false, db
		}
		if *expectedEntity.Value != *e.Value {
			t.Errorf("Returned field should contains inserted entity value for Value field and it is %v (expeted %v)", e.Value, expectedEntity.Value)
			return false, db
		}
		if *expectedEntity.Key != *e.Key {
			t.Errorf("Returned field should contains inserted entity value for Key field and it is %v (expeted %v)", e.Key, expectedEntity.Key)
			return false, db
		}
	}
	if count != 1 {
		t.Errorf("FindAll should return one result and it return %v results", count)
		return false, db
	}
	return true, db
}
