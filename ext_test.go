package sqlite3_ext

import (
	"database/sql"
	"testing"
)

func TestOpenReturnsWithoutError(t *testing.T) {
	db, err := sql.Open("sqlite3-ext", ":memory:")

	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	err = db.Ping()

	if err != nil {
		t.Fatalf("%s", err.Error())
	}
}
