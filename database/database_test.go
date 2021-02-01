package database

import (
	"testing"
)

func TestCheckDatabaseConnection(t *testing.T) {
	db := Check()

	if db == nil {
		t.FailNow()
	}
}
