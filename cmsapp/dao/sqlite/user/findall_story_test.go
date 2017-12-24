package userdao

import (
	maindef "github.com/goatcms/goatcms/cmsapp/dao"
	entities "github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcore/app/scope"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestFindAllStory(t *testing.T) {
	t.Parallel()
	doFindAllStory(t)
}

func doFindAllStory(t *testing.T) (bool, *sqlx.DB) {
	var (
		rows maindef.Rows
		ok   bool
		db   *sqlx.DB
		err  error
	)
	if ok, db, _ = doInsertStory(t); !ok {
		return false, nil
	}
	s := scope.NewScope("tag")
	finder := UserFindAll{}
	finder.deps.DB = db
	if rows, err = finder.Find(s, entities.UserMainFields); err != nil {
		t.Error(err)
		return false, db
	}
	// expected list witn exaclly one row
	// TODO: check this

	// iterate over each row
	count := 0
	expectedEntity := NewMockEntity1()
	for rows.Next() {
		var e entities.User
		count++
		if err = rows.StructScan(&e); err != nil {
			t.Error(err)
			return false, db
		}
		if expectedEntity.Lastname != e.Lastname {
			t.Errorf("Returned field should contains inserted entity value for Lastname field and it is %v (expeted %v)", e.Lastname, expectedEntity.Lastname)
			return false, db
		}
		if expectedEntity.Firstname != e.Firstname {
			t.Errorf("Returned field should contains inserted entity value for Firstname field and it is %v (expeted %v)", e.Firstname, expectedEntity.Firstname)
			return false, db
		}
	}
	if count != 1 {
		t.Errorf("FindAll should return one result and it return %v results", count)
		return false, db
	}
	return true, db
}
