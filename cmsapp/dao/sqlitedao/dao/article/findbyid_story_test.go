package dao

import (
	"database/sql"
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
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
		row            maindef.ArticleRow
		ok             bool
		db             *sql.DB
		err            error
		expectedEntity *entities.Article
		entity         *entities.Article
	)
	if ok, db, expectedEntity = doInsertStory(t); !ok {
		return false, nil
	}
	s := scope.NewScope("tag")
	finder := ArticleFindByID{}
	finder.deps.DB = db
	if row, err = finder.Find(s, entities.ArticleMainFields, *expectedEntity.ID); err != nil {
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
	if *expectedEntity.Title != *entity.Title {
		t.Errorf("Returned field should contains inserted entity value for Title field and it is %v (expeted %v)", entity.Title, expectedEntity.Title)
		return false, db
	}
	return true, db
}
