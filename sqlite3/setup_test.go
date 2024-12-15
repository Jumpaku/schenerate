package sqlite3

import (
	"os"
	"testing"
)

func Setup(t *testing.T, dbPath string, ddls []string) (queryer queryer, teardown func()) {
	t.Helper()

	q, err := Open(dbPath)
	if err != nil {
		t.Fatalf("failed to open: %v", err)
	}

	for i, ddl := range ddls {
		_, err := q.db.Exec(ddl)
		if err != nil {
			t.Fatalf("failed to exec %v: %v", i, err)
		}
	}

	teardown = func() {
		q.Close()
		os.Remove(dbPath)
	}
	return q, teardown
}
