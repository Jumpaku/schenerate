package sqlite3

import (
	"database/sql"
	"os"
	"testing"
)

func Setup(t *testing.T, dbPath string, ddls []string) (q queryer, teardown func()) {
	t.Helper()

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		t.Fatalf("failed to open: %v", err)
	}
	defer db.Close()

	for i, ddl := range ddls {
		_, err := db.Exec(ddl)
		if err != nil {
			t.Fatalf("failed to exec %v: %v", i, err)
		}
	}

	dbx, err := Open(dbPath)
	if err != nil {
		t.Fatalf("failed to open: %v", err)
	}
	teardown = func() {
		dbx.Close()
		os.Remove(dbPath)
	}
	return dbx, teardown
}
