package db_sqlc

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// Init initializes the database connection pool.
func Init() {
	// Open a database connection
	var err error
	DB, err = sql.Open("sqlite3", "registry.db")
	if err != nil {
		log.Fatal(err)
	}

	// Test the database connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
}

// Init initializes the database connection pool.
func InitTestDB() {
	// Open a database connection
	var err error
	DB, err = sql.Open("sqlite3", "../registry_test.db")
	if err != nil {
		log.Fatal(err)
	}

	// Test the database connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	log.Println("connected to test db...")
}
