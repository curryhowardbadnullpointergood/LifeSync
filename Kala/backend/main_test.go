package main

import (
	"database/sql"
	"os"
	"testing"

	_ "modernc.org/sqlite"
)

/*
func TestAddTask(t *testing.T) {
	dbFile := "test_kala.db"
	defer os.Remove(dbFile) // cleanup after test

	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	CreateTable(db)

	err = AddTask(db,
		"Test Task",
		"Testing details",
		"2026-01-28",
		"",
		"14:00",
	)
	if err != nil {
		t.Fatalf("AddTask failed: %v", err)
	}

	var count int
	err = db.QueryRow(`SELECT COUNT(*) FROM tasks;`).Scan(&count)
	if err != nil {
		t.Fatalf("query failed: %v", err)
	}

	if count != 1 {
		t.Fatalf("expected 1 task, got %d", count)
	}
}

*/

func TestGetTaskSimple(t *testing.T) {
	dbFile := "test_kala_simple.db"
	defer os.Remove(dbFile) // cleanup after test

	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	// create schema
	CreateTable(db)

	// insert a "simple" task (no date, no time, no enddate)
	_, err = db.Exec(`
		INSERT INTO tasks (
			title, details, date, time, enddate, repeat 
		) VALUES (?, ?, ?, ?, ?, ?);
	`,
		"Simple Task",
		"Always visible",
		"null",
		"null",
		"null",
		"null",
	)
	if err != nil {
		t.Fatalf("failed to insert task: %v", err)
	}

	// call the function under test
	tasks, err := GetTaskSimple(db)
	if err != nil {
		t.Fatalf("GetTaskSimple failed: %v", err)
	}

	// basic assertion
	if len(tasks) != 1 {
		t.Fatalf("expected 1 simple task, got %d", len(tasks))
	}

	// sanity check content
	if tasks[0].Title != "Simple Task" {
		t.Fatalf("unexpected task title: %s", tasks[0].Title)
	}

	
}

