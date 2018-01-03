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
		rows           maindef.UserRows
		ok             bool
		db             *sql.DB
		err            error
		expectedEntity *entities.User
	)
	if ok, db, expectedEntity = doInsertStory(t); !ok {
		return false, nil
	}
	s := scope.NewScope("tag")
	searcher := UserSearch{}
	searcher.deps.DB = db
	if rows, err = searcher.Search(s, entities.UserMainFields, &maindef.UserSearchParams{
		Firstname: *expectedEntity.Firstname,
		Password:  *expectedEntity.Password,
		Login:     *expectedEntity.Login,
		Email:     *expectedEntity.Email,
	}); err != nil {
		t.Error(err)
		return false, db
	}
	// iterate over each row
	count := 0
	for rows.Next() {
		var e *entities.User
		count++
		if e, err = rows.Get(); err != nil {
			t.Error(err)
			return false, db
		}
		if *expectedEntity.Firstname != *e.Firstname {
			t.Errorf("Returned field should contains inserted entity value for Firstname field and it is %v (expeted %v)", e.Firstname, expectedEntity.Firstname)
			return false, db
		}
		if *expectedEntity.Password != *e.Password {
			t.Errorf("Returned field should contains inserted entity value for Password field and it is %v (expeted %v)", e.Password, expectedEntity.Password)
			return false, db
		}
		if *expectedEntity.Login != *e.Login {
			t.Errorf("Returned field should contains inserted entity value for Login field and it is %v (expeted %v)", e.Login, expectedEntity.Login)
			return false, db
		}
		if *expectedEntity.Email != *e.Email {
			t.Errorf("Returned field should contains inserted entity value for Email field and it is %v (expeted %v)", e.Email, expectedEntity.Email)
			return false, db
		}
	}
	if count != 1 {
		t.Errorf("FindAll should return one result and it return %v results", count)
		return false, db
	}
	return true, db
}
