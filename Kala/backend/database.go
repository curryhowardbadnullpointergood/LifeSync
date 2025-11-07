package main

import (
	"database/sql"
	"log"
	

	_ "modernc.org/sqlite"
)



// creates all of the tables for the calendar application 
// repeat is a text field and do day, week, year or null  
// this is if the event repeats or not
func CreateTable(db *sql.DB) {
	create := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		details TEXT NOT NULL,
		date TEXT NOT NULL,
		repeat TEXT


	);`

	_, err := db.Exec(create)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	log.Println("Table created")
}

// drops tables, used for testing purposes 
func dropTable(db *sql.DB) {
	_, err := db.Exec(`DROP TABLE IF EXISTS tasks;`)
	if err != nil {
		log.Fatal("Failed to drop table:", err)
	}
	log.Println("Table dropped successfully")
}


