package main

import (
	"database/sql"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Check flags
	// -config configuration file
	// -service start starts the service
	// -service stop stops the service
	// -renew
	var (
		config  *string = flag.String("config", "", "")
		service *string = flag.String("service", "", "")
		renew   *bool   = flag.Bool("renew", false, "")
	)

	flag.Parse()

	checkDatabaseConnection()
	checkDatabaseProvision()
	setupConfigIfNeeded(config)
	checkConfigOk()
	startStopService(service)
	renewEncryption(renew)
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

// Checks if the user wants to
// set/update the configuration
func setupConfigIfNeeded(config *string) {
	if len(*config) > 0 {
		if _, err := os.Stat(*config); err == nil {
			content, err := ioutil.ReadFile(*config)

			if err != nil {
				log.Fatal("Could not read file " + *config)
			}

			setupConfig(&content)
		} else {
			log.Fatal("Could not read file " + *config)
		}
	}
}

// Sets up the config based
// on the passed yaml content
func setupConfig(content *[]byte) {

}

// Starts or stops the service based
// on what the passed string says
func startStopService(service *string) {
	switch *service {
	case "start":
		startListenTCP()
		startListenBluetooth()
	case "stop":
		stopListenTCP()
		stopListenBluetooth()
	}
}

// Generates a new encryption key,
// decrypts and re-encrypts all the
// data inside the database
func renewEncryption(renew *bool) {

}

// Checks if the configuration is OK
func checkConfigOk() {

}

// Start listening for incoming
// connections on port 1000
// and dispatches the commands
func startListenTCP() {

}

// Start listening for incoming
// connections via Bluetooth
// and dispatches the commands
func startListenBluetooth() {

}

// Start listening for incoming
// connections on port 1000
// and dispatches the commands
func stopListenTCP() {

}

// Start listening for incoming
// connections via Bluetooth
// and dispatches the commands
func stopListenBluetooth() {

}
