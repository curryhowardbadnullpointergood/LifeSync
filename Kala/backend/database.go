package main

import (
	"database/sql"
	"log"
	"fmt"
	

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
		repeatoccurrences INTEGRER,
        repeatNever TEXT, 
        repeatEndDate TEXT

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


// simple task struct, this is for tasks with just title and detail nothing else
type SimpleTask struct {
	ID      int		`json:"id"`
	Title   string	`json:"title"`
	Details string	`json:"details"`
}





// AddTask inserts a new task into the DB
func AddTask(db *sql.DB, title, details, date, repeat, time, enddate, numrepeat, repeatoccurrences, repeatnever, repeatenddate string) error {

	

	CreateTable(db)

	query := `
	INSERT INTO tasks (title, details, date, repeat, time, enddate, numrepeat, repeatoccurrences, repeatnever, repeatenddate )
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

    // checking if repeat is none, which means repeat isn't like checked, so should be False boolean value 
    //this code is a bit cursed but code fast and break things 
    // plus I think it's a neat quick solution 
    //if repeat == "" {
        //print("repeat is not toggled so it is not set to false.")
        // this works already saving as false text 
      //  repeat = "null"; 
    //}

    _, err := db.Exec(query, title, details, date, repeat, time, enddate, numrepeat, repeatoccurrences, repeatnever, repeatenddate)
	return err
}



// gets tasks so that it can be shown to the ui if that makes sense right 
// so we need like 2 different get tasks 
// one get task gets like the tasks without a time/date
// the other get tasks deal with that 
// going to keep this seperate so like, everytime i am calling a get task im not going to query or have to display 

func GetTaskSimple( db *sql.DB) ([]SimpleTask, error){


	rows, err := db.Query(`
		SELECT id, title, details
		FROM tasks
		WHERE date = 'null'
		  AND time = 'null'
		  AND enddate = 'null'
		  AND repeat = 'null';
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []SimpleTask

	for rows.Next() {
		var t SimpleTask
		if err := rows.Scan(&t.ID, &t.Title, &t.Details); err != nil {
			return nil, err
		}

		fmt.Printf("TASK: ID=%d | Title=%s | Details=%s\n",
		t.ID, t.Title, t.Details)


		tasks = append(tasks, t)
	}

	

	
	return tasks, nil

}












































