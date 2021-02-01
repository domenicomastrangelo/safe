package database

import (
	"database/sql"
	"testing"
)

func Test_Check(t *testing.T) {
	var db interface{}

	db = Check()

	if _, ok := db.(*sql.DB); !ok {
		t.FailNow()
	}
}

func Test_checkDatabaseConnection(t *testing.T) {
	var db interface{}

	db = Check()

	if _, ok := db.(*sql.DB); !ok {
		t.FailNow()
	}
}

func Test_checkDataTable(t *testing.T) {
	db := Check()

	var check interface{}
	check = checkDataTable(db)

	if _, ok := check.(bool); !ok {
		t.FailNow()
	}
}

func Test_checkTypesTable(t *testing.T) {
	db := Check()

	var check interface{}
	check = checkTypesTable(db)

	if _, ok := check.(bool); !ok {
		t.FailNow()
	}
}

func Test_checkConfigTable(t *testing.T) {
	db := Check()

	var check interface{}
	check = checkConfigTable(db)

	if _, ok := check.(bool); !ok {
		t.FailNow()
	}
}

func Test_checkDatabaseProvision(t *testing.T) {
	db := Check()

	checkDatabaseProvision(db)
}
