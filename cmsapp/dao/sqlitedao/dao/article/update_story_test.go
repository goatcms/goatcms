package dao

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
		expectedEntity *entities.Article
		entity         *entities.Article
	)
	expectedEntity = NewMockEntity2()
	if ok, db, entity = doInsertStory(t); !ok {
		return false, nil
	}
	entity.Content = expectedEntity.Content
	entity.Title = expectedEntity.Title
	s := scope.NewScope("tag")
	updater := ArticleUpdate{}
	updater.deps.DB = db
	if err = updater.Update(s, entity, entities.ArticleMainFields); err != nil {
		t.Error(err)
		return false, db
	}
	finder := ArticleFindByID{}
	finder.deps.DB = db
	if row, err = finder.Find(s, entities.ArticleMainFields, entity.ID); err != nil {
		t.Error(err)
		return false, db
	}
	// iterate over each row
	entity = &entities.Article{}
	if err = row.StructScan(entity); err != nil {
		t.Error(err)
		return false, db
	}
	if expectedEntity.Content != entity.Content {
		t.Errorf("Returned field should contains inserted entity value for Content field and it is %v (expeted %v)", entity.Content, expectedEntity.Content)
		return false, db
	}
	if expectedEntity.Title != entity.Title {
		t.Errorf("Returned field should contains inserted entity value for Title field and it is %v (expeted %v)", entity.Title, expectedEntity.Title)
		return false, db
	}
	return true, db
}
