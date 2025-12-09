package main

import (
	"database/sql"
	"log"
	

	_ "modernc.org/sqlite"
)



// creates all of the tables for the calendar application 
// TASKS: ---------------------------------------------------------------
// in tasks repeat is a text field and do day, week, year or none (in text) 
// this is if the event repeats or not
// tasks, time is the time on the date ie, lunch at 2pm etc 
// end date -- field that is for the end of the task (if range) 
// num_repeat -- num value of time gap between repeats 
// repeat_timeframe -- timeframe between repeats is 1 day, 1 month, etc
// should timeframe have an infinity option? probably not that will break things ugh 
// maybe not, this is a pain, no 1 year max I think.  review what you spend your time 
// on every new year 
func CreateTable(db *sql.DB) {
	create := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
	 	details TEXT NOT NULL,
		date TEXT NOT NULL,
		repeat TEXT,
		time TEXT,
		enddate TEXT, 
		numrepeat INTEGER,
		repeattimeframe TEXT,
		repeatoccureences INTEGRER

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



