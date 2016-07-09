package database

import "testing"

func TestShouldOpenConncetion(t *testing.T) {
	const dbpath = "test.db"

	db, err := NewDatabase(dbpath)
	if err != nil {
		t.Error("Can not create new database", err)
		return
	}
	if err := db.Open(); err != nil {
		t.Error("Can not open new conncetion", err)
		return
	}
	defer db.Close()
}

func TestShouldScreateSchema(t *testing.T) {
	const dbpath = "test2.db"

	db, err := NewDatabase(dbpath)
	if err != nil {
		t.Error("Can not create new database", err)
		return
	}
	if err := db.Open(); err != nil {
		t.Error("Can not open new conncetion", err)
		return
	}
	defer db.Close()

	if err := db.CreateTables(); err != nil {
		t.Error("Can not create schema", err)
		return
	}
}
