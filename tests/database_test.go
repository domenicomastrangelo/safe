package tests

import (
	"testing"

	"github.com/domenicomastrangelo/safe/database"
)

func TestCheckDatabaseConnection(t *testing.T) {
	db := database.Check()

	if db == nil {
		t.FailNow()
	}
}
