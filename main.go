package main

import (
	"flag"

	"github.com/domenicomastrangelo/safe/backup"
	"github.com/domenicomastrangelo/safe/config"
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
		configFile   *string = flag.String("config", "", "")
		service      *string = flag.String("service", "", "")
		renew        *bool   = flag.Bool("renew", false, "")
		shouldBackup *bool   = flag.Bool("backup", false, "")
	)

	flag.Parse()

	backup.Everything(shouldBackup)
	database.Check()
	config.SetupIfNeeded(configFile)
	config.Check()
	startStopService(service)
	renewEncryption(renew)
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
