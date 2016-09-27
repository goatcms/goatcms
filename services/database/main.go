package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // import sqlite3 adapter as default db driver
)

// Database is global database connection provider
type Database struct {
	instance *sqlx.DB
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
	db.instance, err = sqlx.Open("sqlite3", db.filepath)
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
func (db *Database) Adapter() *sqlx.DB {
	return db.instance
}

/*
// CreateTables build a database schema
func (db *Database) CreateTables() error {

	query := `
	CREATE TABLE IF NOT EXISTS users(
		id INT PRIMARY KEY,
		email TEXT NOT NULL,
		pass_hash TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS articles(
		id INT PRIMARY KEY,
		image_id INT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		//image_id INT FOREIGN KEY REFERENCES images(id)
	);
	CREATE TABLE IF NOT EXISTS images(
		id INT PRIMARY KEY,
		name VARCHAR(25) UNIQUE NOT NULL,
		location VARCHAR(300) UNIQUE NOT NULL,
		description TEXT,
		created_at datetime NOT NULL
	);
		`
	_, err := db.instance.Exec(query)
	if err != nil {
		return err
	}
	return nil
}*/
