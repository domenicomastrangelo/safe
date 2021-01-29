package service

// ServiceTCP is a constant used
// to specify the TCP service
const ServiceTCP = 1

// ServiceBluetooth is a constant used
// to specify the bluetooth service
const ServiceBluetooth = 2

// StartServices starts the services
// listed in the passed struct
func StartServices(services []int) {
	for service := range services {
		switch service {
		case ServiceTCP:
			startListenTCP()
		case ServiceBluetooth:
			startListenBluetooth()
		}
	}
}

// StopServices starts the services
// listed in the passed struct
func StopServices(services []int) {
	for service := range services {
		switch service {
		case ServiceTCP:
			stopListenTCP()
		case ServiceBluetooth:
			stopListenBluetooth()
		}
	}
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
