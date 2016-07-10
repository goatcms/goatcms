package database

import "testing"

func TestShouldOpenConnection(t *testing.T) {
	const dbpath = "test.db"

	db, err := NewDatabase(dbpath)
	if err != nil {
		t.Error("Can not create new database", err)
		return
	}
	if err := db.Open(); err != nil {
		t.Error("Can not open new connection", err)
		return
	}
	defer db.Close()
	// test doesn't create test.db file?
	// should it?
}

func TestShouldCreateSchema(t *testing.T) {
	const dbpath = "test2.db"
	// this test works

	db, err := NewDatabase(dbpath)
	if err != nil {
		t.Error("Can not create new database", err)
		return
	}
	if err := db.Open(); err != nil {
		t.Error("Can not open new connection", err)
		return
	}
	defer db.Close()

	if err := db.CreateTables(); err != nil {
		t.Error("Can not create schema", err)
		return
	}
	// maybe also delete test2.db file?
}
