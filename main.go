package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/domenicomastrangelo/safe/backup"
	"github.com/domenicomastrangelo/safe/database"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Check flags
	// -config configuration file
	// -service start starts the service
	// -service stop stops the service
	// -renew
	// -backup
	var (
		config       *string = flag.String("config", "", "")
		service      *string = flag.String("service", "", "")
		renew        *bool   = flag.Bool("renew", false, "")
		shouldBackup *bool   = flag.Bool("backup", false, "")
	)

	flag.Parse()

	backup.Everything(shouldBackup)
	database.Check()
	setupConfigIfNeeded(config)
	checkConfigOk()
	startStopService(service)
	renewEncryption(renew)
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
