package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // import sqlite3 adapter as default db driver
)

// Database is global database connection provider
type Database struct {
	instance *sql.DB
	filepath string
}

// NewDatabase create a database instance
func NewDatabase(filepath string) (*Database, error) {
	return &Database{
		instance: nil,
		filepath: filepath,
	}, nil
}

// Open initialize a database connection
func (db *Database) Open() error {
	var err error
	db.instance, err = sql.Open("sqlite3", db.filepath)
	if err != nil {
		return err
	}
	if db.instance == nil {
		return fmt.Errorf("Can not open database conection (pointer is nil)")
	}
	return nil
}

// Close close a connection
func (db *Database) Close() error {
	return db.instance.Close()
}

// Adapter return current golang database interface instance
func (db *Database) Adapter() *sql.DB {
	return db.instance
}

// CreateTables build a database schema
func (db *Database) CreateTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY,
		email TEXT NOT NULL,
		pass_hash TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS articles(
		id INTEGER PRIMARY KEY,
		title TEXT NOT NULL DEFAULT '',
		content TEXT NOT NULL DEFAULT ''
	);
	CREATE TABLE IF NOT EXISTS images(
		id INTEGER PRIMARY KEY,
		article_id INTEGER NOT NULL,
		name TEXT NOT NULL DEFAULT '',
		location TEXT NOT NULL DEFAULT '',
		description TEXT NOT NULL DEFAULT '',
		size int(11) NOT NULL,
		created_at datetime NOT NULL
	);
		`
	_, err := db.instance.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
