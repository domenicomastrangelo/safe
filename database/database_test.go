package database

import (
	"testing"
)

func TestCheckDatabaseConnection(t *testing.T) {
	db := checkDatabaseConnection()

	if db == nil {
		t.FailNow()
	}
}

func TestCheckDatabaseProvision(t *testing.T) {

}
