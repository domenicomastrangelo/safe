package database

import (
	"database/sql"
	"testing"
)

func TestCheck(t *testing.T) {
	var db interface{}

	db = Check()

	if _, ok := db.(*sql.DB); !ok {
		t.FailNow()
	}
}

func TestCheckDatabaseConnection(t *testing.T) {
	var db interface{}

	db = Check()

	if _, ok := db.(*sql.DB); !ok {
		t.FailNow()
	}
}

func TestCheckDataTable(t *testing.T) {
	db := Check()

	var check interface{}
	check = checkDataTable(db)

	if _, ok := check.(bool); !ok {
		t.FailNow()
	}
}

func TestCheckTypesTable(t *testing.T) {
	db := Check()

	var check interface{}
	check = checkTypesTable(db)

	if _, ok := check.(bool); !ok {
		t.FailNow()
	}
}

func TestCheckConfigTable(t *testing.T) {
	db := Check()

	var check interface{}
	check = checkConfigTable(db)

	if _, ok := check.(bool); !ok {
		t.FailNow()
	}
}

func TestCheckDatabaseProvision(t *testing.T) {
	db := Check()

	checkDatabaseProvision(db)
}
