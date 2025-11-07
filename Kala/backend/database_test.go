package main

import (
	"database/sql"
	"os"
	"testing"

	_ "modernc.org/sqlite"
)

func setupTestDB(t *testing.T) *sql.DB {
	t.Helper()

	testDBPath := "./test_kala.db"
	_ = os.Remove(testDBPath) // ensure clean slate

	db, err := sql.Open("sqlite", testDBPath)
	if err != nil {
		t.Fatalf("failed to open test DB: %v", err)
	}

	return db
}

func teardownTestDB(db *sql.DB) {
	db.Close()
	os.Remove("./test_kala.db")
}

func TestCreateTable(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(db)

	createTable(db)

	row := db.QueryRow(`
		SELECT name FROM sqlite_master WHERE type='table' AND name='tasks';
	`)

	var tableName string
	err := row.Scan(&tableName)

	if err != nil {
		t.Fatalf("tasks table not created: %v", err)
	}

	if tableName != "tasks" {
		t.Fatalf("expected table 'tasks', got '%s'", tableName)
	}
}

func TestDropTable(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(db)

	createTable(db)
	dropTable(db)

	row := db.QueryRow(`
		SELECT name FROM sqlite_master WHERE type='table' AND name='tasks';
	`)

	var tableName string
	err := row.Scan(&tableName)

	if err == nil {
		t.Fatalf("expected tasks table to be dropped but it still exists")
	}
}
