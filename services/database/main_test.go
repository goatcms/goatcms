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
}
