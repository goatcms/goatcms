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
		expectedEntity *entities.User
		entity         *entities.User
	)
	if ok, db, expectedEntity = doInsertStory(t); !ok {
		return false, nil
	}
	s := scope.NewScope("tag")
	finder := UserFindByID{}
	finder.deps.DB = db
	if row, err = finder.Find(s, entities.UserMainFields, expectedEntity.ID); err != nil {
		t.Error(err)
		return false, db
	}
	// iterate over each row
	entity = &entities.User{}
	if err = row.StructScan(entity); err != nil {
		t.Error(err)
		return false, db
	}
	if expectedEntity.Email != entity.Email {
		t.Errorf("Returned field should contains inserted entity value for Email field and it is %v (expeted %v)", entity.Email, expectedEntity.Email)
		return false, db
	}
	if expectedEntity.Password != entity.Password {
		t.Errorf("Returned field should contains inserted entity value for Password field and it is %v (expeted %v)", entity.Password, expectedEntity.Password)
		return false, db
	}
	if expectedEntity.Login != entity.Login {
		t.Errorf("Returned field should contains inserted entity value for Login field and it is %v (expeted %v)", entity.Login, expectedEntity.Login)
		return false, db
	}
	if expectedEntity.Firstname != entity.Firstname {
		t.Errorf("Returned field should contains inserted entity value for Firstname field and it is %v (expeted %v)", entity.Firstname, expectedEntity.Firstname)
		return false, db
	}
	return true, db
}
