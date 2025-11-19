package main

import (
	"database/sql"
	"log"
	

	_ "modernc.org/sqlite"
)



// creates all of the tables for the calendar application 
// repeat is a text field and do day, week, year or none (in text) 
// this is if the event repeats or not
func CreateTable(db *sql.DB) {
	create := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
	 	details TEXT NOT NULL,
		date TEXT NOT NULL,
		repeat TEXT,
		time TEXT


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


// AddTask inserts a new task into the DB
func AddTask(title, details, date, repeat, time string) error {

	db, err := sql.Open("sqlite", "./kala.db")
	if err != nil {
		return err
	}
	defer db.Close()

	CreateTable(db)

	query := `
	INSERT INTO tasks (title, details, date, repeat, time)
	VALUES (?, ?, ?, ?, ?);`

	_, err = db.Exec(query, title, details, date, repeat, time)
	return err
}



