package sqlite3

import (
	"database/sql"
	"os"
	"testing"
)

func Setup(t *testing.T, dbPath string, ddls []string) (sqlite3 *sql.DB, teardown func()) {
	t.Helper()

	sqlite3, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		t.Fatalf("failed to open: %v", err)
	}

	for i, ddl := range ddls {
		_, err := sqlite3.Exec(ddl)
		if err != nil {
			t.Fatalf("failed to exec %v: %v", i, err)
		}
	}

	teardown = func() {
		sqlite3.Close()
		os.Remove(dbPath)
	}
	return sqlite3, teardown
}
