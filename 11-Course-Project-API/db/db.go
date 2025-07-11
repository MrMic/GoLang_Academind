package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	DB, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	DB.SetMaxOpenConns(10) // * NOTE: Set the maximum number of open connections to the database
	DB.SetMaxIdleConns(5)  // * NOTE: Set the maximum number of idle connections to the database
}
