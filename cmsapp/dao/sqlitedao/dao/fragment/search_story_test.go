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
		rows           maindef.FragmentRows
		ok             bool
		db             *sql.DB
		err            error
		expectedEntity *entities.Fragment
	)
	if ok, db, expectedEntity = doInsertStory(t); !ok {
		return false, nil
	}
	s := scope.NewScope("tag")
	searcher := FragmentSearch{}
	searcher.deps.DB = db
	if rows, err = searcher.Search(s, entities.FragmentMainFields, &maindef.FragmentSearchParams{
		Name:    *expectedEntity.Name,
		Content: *expectedEntity.Content,
		Lang:    *expectedEntity.Lang,
	}); err != nil {
		t.Error(err)
		return false, db
	}
	// iterate over each row
	count := 0
	for rows.Next() {
		var e *entities.Fragment
		count++
		if e, err = rows.Get(); err != nil {
			t.Error(err)
			return false, db
		}
		if *expectedEntity.Name != *e.Name {
			t.Errorf("Returned field should contains inserted entity value for Name field and it is %v (expeted %v)", e.Name, expectedEntity.Name)
			return false, db
		}
		if *expectedEntity.Content != *e.Content {
			t.Errorf("Returned field should contains inserted entity value for Content field and it is %v (expeted %v)", e.Content, expectedEntity.Content)
			return false, db
		}
		if *expectedEntity.Lang != *e.Lang {
			t.Errorf("Returned field should contains inserted entity value for Lang field and it is %v (expeted %v)", e.Lang, expectedEntity.Lang)
			return false, db
		}
	}
	if count != 1 {
		t.Errorf("FindAll should return one result and it return %v results", count)
		return false, db
	}
	return true, db
}
