package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
)

// Check executes all the necessary
// checks to guarantee that the
// database exists, if it doesn't
// it will create one and if it can't,
// it will log.Fatal
func Check() {
	checkDatabaseConnection()
	checkDatabaseProvision()
}

// Checks if the database connection
// is OK, otherwise makes the program
// fail
func checkDatabaseConnection() {
	var (
		dbFile           *os.File
		db               *sql.DB
		filePath         string
		err              error
		row              *sql.Row
		testSelectResult string
	)

	if filePath, err = filepath.Abs("./db.sqlite"); err != nil {
		if dbFile, err = os.Create(filePath); err != nil {
			defer dbFile.Close()
			log.Fatal("Database file does not exist. Could not create database file.")
		}
	}

	if db, err = sql.Open("sqlite3", filePath); err != nil {
		defer db.Close()
		log.Fatal("Could not open the database file")
	}

	row = db.QueryRow("SELECT 1")

	if err = row.Scan(&testSelectResult); err != nil {
		log.Fatal("Could not execute test select statement")
	}
}

// Checks that the database exists
// or creates it
func checkDatabaseExists(filePath string) {

}

// Checks if the database needs to be
// provisioned
func checkDatabaseProvision() {

}
