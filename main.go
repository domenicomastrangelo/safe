package main

import (
	"flag"

	"./backup"
	"./config"
	"./database"
	"./encryption"
	"./service"
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
		configFile    *string = flag.String("config", "", "")
		serviceToggle *string = flag.String("service", "", "")
		renew         *bool   = flag.Bool("renew", false, "")
		shouldBackup  *bool   = flag.Bool("backup", false, "")
	)

	flag.Parse()

	backup.Everything(shouldBackup)
	database.Check()
	config.SetupIfNeeded(configFile)
	config.Check()

	switch *serviceToggle {
	case "start":
		service.StartServices([]int{service.ServiceTCP, service.ServiceBluetooth})
	case "stop":
		service.StartServices([]int{service.ServiceTCP, service.ServiceBluetooth})
	}

	encryption.RenewEncryption(renew)
}
