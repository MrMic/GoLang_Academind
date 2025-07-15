package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error // Declare err variable first

	// Use assignment (=) instead of short declaration (:=) to assign to global DB
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	DB.SetMaxOpenConns(10) // * NOTE: Set the maximum number of open connections to the database
	DB.SetMaxIdleConns(5)  // * NOTE: Set the maximum number of idle connections to the database

	createTables() // * NOTE: Call the function to create the necessary tables in the database
}

// ------------------------------------------------------------------------------
func createTables() {
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	)
	`
	_, err := DB.Exec(createEventsTable) // * NOTE: Execute the SQL command to create the events table if it does not exist
	if err != nil {
		panic("Failed to create events table: " + err.Error())
	}
}
