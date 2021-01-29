package config

import (
	"io/ioutil"
	"log"
	"os"
)

// SetupIfNeeded Checks if the user wants to
// set/update the configuration, if so
// it will proceed to do so
func SetupIfNeeded(config *string) {
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

// Check checks if the configuration is OK
func Check() {

}
