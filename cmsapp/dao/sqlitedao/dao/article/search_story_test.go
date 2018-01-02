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
		rows           maindef.ArticleRows
		ok             bool
		db             *sql.DB
		err            error
		expectedEntity *entities.Article
	)
	if ok, db, expectedEntity = doInsertStory(t); !ok {
		return false, nil
	}
	s := scope.NewScope("tag")
	searcher := ArticleSearch{}
	searcher.deps.DB = db
	if rows, err = searcher.Search(s, entities.ArticleMainFields, &maindef.ArticleSearchParams{
		Content: *expectedEntity.Content,
		Title:   *expectedEntity.Title,
	}); err != nil {
		t.Error(err)
		return false, db
	}
	// iterate over each row
	count := 0
	for rows.Next() {
		var e *entities.Article
		count++
		if e, err = rows.Get(); err != nil {
			t.Error(err)
			return false, db
		}
		if *expectedEntity.Content != *e.Content {
			t.Errorf("Returned field should contains inserted entity value for Content field and it is %v (expeted %v)", e.Content, expectedEntity.Content)
			return false, db
		}
		if *expectedEntity.Title != *e.Title {
			t.Errorf("Returned field should contains inserted entity value for Title field and it is %v (expeted %v)", e.Title, expectedEntity.Title)
			return false, db
		}
	}
	if count != 1 {
		t.Errorf("FindAll should return one result and it return %v results", count)
		return false, db
	}
	return true, db
}
