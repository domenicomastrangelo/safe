package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	// Importing sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

// ConfigTable is the name of the
// database table for the configuration
const ConfigTable = "config"

// DBFileName is the SQLite database filename
const DBFileName = "db.sqlite"

// Check executes all the necessary
// checks to guarantee that the
// database exists, if it doesn't
// it will create one and if it can't,
// it will log.Fatal
func Check() *sql.DB {
	db := checkDatabaseConnection()
	checkDatabaseProvision(db)

	return db
}

// Checks if the database connection
// is OK, otherwise makes the program
// fail
func checkDatabaseConnection() *sql.DB {
	var (
		dbFile           *os.File
		db               *sql.DB
		filePath         string
		err              error
		row              *sql.Row
		testSelectResult string
	)

	if filePath, err = filepath.Abs("../" + DBFileName); err != nil {
		if dbFile, err = os.Create(filePath); err != nil {
			defer dbFile.Close()
			log.Fatalln("Database file does not exist. Could not create database file.")
		}
	}

	if db, err = sql.Open("sqlite3", filePath); err != nil {
		defer db.Close()
		log.Fatalln("Could not open the database file")
	}

	row = db.QueryRow("SELECT 1")

	if err = row.Scan(&testSelectResult); err != nil {
		log.Fatalln("Could not execute test select statement")
	}

	return db
}

// Checks if the database needs to be
// provisioned
func checkDatabaseProvision(db *sql.DB) {
	if !checkDataTable(db) || !checkTypesTable(db) || !checkConfigTable(db) {
		if !provisionDatabase(db) {
			log.Fatalln("Could not provision the database. Try running with -cleanprovision")
		}
	}
}

func provisionDatabase(db *sql.DB) bool {
	if !provisionTypesTable(db) || !provisionDataTable(db) || !provisionConfigTable(db) {
		return false
	}

	return true
}

func provisionTypesTable(db *sql.DB) bool {
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS 'types' (
				id integer auto_increment primary key, 
				name varchar(255)
			)`); err != nil {
		return false
	}

	return true
}

func provisionDataTable(db *sql.DB) bool {
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS 'data' (
				id integer auto_increment primary key, 
				type_id integer, 
				value longtext, 
				created_at datetime, 
				updated_at datetime,
				FOREIGN KEY(type_id) REFERENCES types(id)
			)`); err != nil {
		return false
	}

	return true
}

func provisionConfigTable(db *sql.DB) bool {
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS 'config' (
				id integer auto_increment primary key, 
				name varchar(255), 
				value longtext, 
				created_at datetime, 
				updated_at datetime,
				FOREIGN KEY(type_id) REFERENCES types(id)
			)`); err != nil {
		return false
	}

	return true
}

func checkDataTable(db *sql.DB) bool {
	var result string
	row := db.QueryRow("SELECT id, type_id, value, created_at, updated_at FROM data limit 1")

	if err := row.Scan(&result); err != nil {
		return false
	}

	return true
}

func checkTypesTable(db *sql.DB) bool {
	var result string
	row := db.QueryRow("SELECT id, name FROM types limit 1")

	if err := row.Scan(&result); err != nil {
		return false
	}

	return true
}

func checkConfigTable(db *sql.DB) bool {
	var result string
	row := db.QueryRow("SELECT id, name, value, created_at, updated_at FROM config limit 1")

	if err := row.Scan(&result); err != nil {
		return false
	}

	return true
}
