package main

import (
	"database/sql"
	"log"
	"fmt"
	"time"
	

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

// table for completed tasks
func CreateCompletedTasksTable(db *sql.DB) {
	create := `
	CREATE TABLE IF NOT EXISTS completed_tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
        task_id INTEGER NOT NULL,
		title TEXT NOT NULL,
		details TEXT NOT NULL,
		date TEXT NOT NULL,
		time TEXT,
		enddate TEXT,
		completeddate TEXT NOT NULL,
        completedtime TEXT NOT NULL,
        instancedate TEXT NOT NULL
	);`

	_, err := db.Exec(create)
	if err != nil {
		log.Fatal("Failed to create completed_tasks table:", err)
	}

	log.Println("Completed tasks table created")
}

// date here means date when task info is opened/pulled up
// ah guess we also need session date, don't we hmm 
func CreateTaskNotes(db *sql.DB) {
	create := `
	CREATE TABLE IF NOT EXISTS taskNotes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task_id INTEGER NOT NULL,
		taskType TEXT NOT NULL,
		opendate TEXT NOT NULL,
		instancedate TEXT NOT NULL, 
		status TEXT, 
		durationSeconds INTEGER,
		notes_path TEXT,
		UNIQUE(task_id, instancedate)

	);`

	_, err := db.Exec(create)
	if err != nil {
		log.Fatal("Failed to create task notes table:", err)
	}

	log.Println("task notes table created")
}




// drops tables, used for testing purposes 
func dropTable(db *sql.DB) {
	_, err := db.Exec(`DROP TABLE IF EXISTS tasks;`)
	if err != nil {
		log.Fatal("Failed to drop table:", err)
	}
	_, err2 := db.Exec(`DROP TABLE IF EXISTS completed_tasks;`)
	if err2 != nil {
		log.Fatal("Failed to drop table:", err2)
	}
	log.Println("Table dropped successfully")
}


// simple task struct, this is for tasks with just title and detail nothing else
type SimpleTask struct {
	ID      int		`json:"id"`
	Title   string	`json:"title"`
	Details string	`json:"details"`
}

// range task struct 
// thinking we need and end time attribute? here probs makes sense 
// end time 
type RangeTask struct{
	ID	int 		`json:"id"`
	Title string 	`json:"title"`
	Details string 	`json:"details"`
	Date string 	`json:"date"`
	Time string 	`json:"time"`
	EndDate string 	`json:"enddate"`

}

// repeat task struct 
type RepeatTask struct{
	ID	int 					`json:"id"`
	Title string 				`json:"title"`
	Details string 				`json:"details"`
	Date string 				`json:"date"`
	Time string 				`json:"time"`
	Repeat string 				`json:"Repeat"`
	NumRepeat string 			`json:"NumRepeat"`
	RepeatOccurrences string 	`json:"RepeatOccurrences"`
	RepeatNever string 			`json:"RepeatNever"`
	RepeatEndDate string 		`json:"RepeatEndDate"`


}

// this is for completed tasks
type CompletedTask struct {
	ID        int    `json:"id"`
	Title         string `json:"title"`
	Details       string `json:"details"`
	Time          string `json:"time"`
    Date          string `json:"date"`
	EndDate       string `json:"enddate"`
	CompletedDate string `json:"completed_date"`
	CompletedTime string `json:"completed_time"`
    InstanceDate string `json:"instancedate"`
}


// this is for task notes 
type TaskSession struct {
	ID           int    `json:"sessionId"`
	TaskID       int    `json:"taskId"`
	TaskType     string `json:"taskType"`
	InstanceDate string `json:"instanceDate"`
	OpenDate     string `json:"openDate"`
	Status       string `json:"status"`
	Duration     int    `json:"duration"`
	NotesPath    string `json:"notesPath"`
}


// AddTask inserts a new task into the DB
func AddTask(db *sql.DB, title, details, date, repeat, time, enddate, numrepeat, repeatoccurrences, repeatnever, repeatenddate string) error {

	

	CreateTable(db)
    CreateCompletedTasksTable(db)
    CreateTaskNotes(db)
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

    // to stop the error/crash when loading the page

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


// this is for getting tasks with a range
// this will be generated prodecurally dynamically 
// not sure if this will cause some serious fuck ups later down the line 
// i dont see this, the brances are either high, or my brain is cooked, so i Cant visualise how i would deal with some of the 
// problems with this approach and comapre it to my original approach or storing it in data with limited time frame getting around
// this inifinite loop issue with occurences calculated this feels mroe wrong, ugh 
// i genuinely think generating more rows in the db makes sense for range but then id need to add more rows to the db 
// which should honestle be fine but increases complexity 
// idk 

// this filters based on date
func GetTaskRange(db *sql.DB, uiDate time.Time) ([]RangeTask, error) {

    // to stop the error/crash when loading the page

	tasks, err := GetTaskRangehelper(db, uiDate)
	if err != nil {
		return nil, err
	}

	var result []RangeTask

	for _, task := range tasks {

		start, err := time.Parse("2006-01-02", task.Date)
		if err != nil {
			continue
		}

		end, err := time.Parse("2006-01-02", task.EndDate)
		if err != nil {
			continue
		}

		if !uiDate.Before(start) && !uiDate.After(end) {
			result = append(result, task)
		}
	}

	return result, nil
}

func GetTaskRangehelper(db *sql.DB, uiDate time.Time) ([]RangeTask, error) {

	targetDate := uiDate.Format("2006-01-02")

	rows, err := db.Query(`
		SELECT t.id, t.title, t.details, t.date, t.time, t.enddate
		FROM tasks t
		WHERE t.numrepeat = 'null'
		  AND t.repeat = 'null'
		  AND t.enddate != 'null'
		  AND t.date != 'null'
		  AND NOT EXISTS (
		      SELECT 1
		      FROM completed_tasks c
		      WHERE c.task_id = t.id
		        AND c.instancedate = ?
		  );
	`, targetDate)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []RangeTask

	for rows.Next() {
		var t RangeTask

		if err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Details,
			&t.Date,
			&t.Time,
			&t.EndDate,
		); err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}

	return tasks, nil
}


// range seems to have a neat solution but im not sure about 
// repeats this is a pain 

func GetRepeatTasksFilter(db *sql.DB) ([]RepeatTask, error) {

	rows, err := db.Query(`
		SELECT
			id,
			title,
			details,
			date,
			time,
			repeat,
			numrepeat,
			repeatoccurrences,
			repeatNever,
			repeatEndDate
		FROM tasks
		WHERE repeat != 'null'
		  AND date != 'null';
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []RepeatTask

	for rows.Next() {
		var t RepeatTask
		if err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Details,
			&t.Date,
			&t.Time,
			&t.Repeat,
			&t.NumRepeat,
			&t.RepeatOccurrences,
			&t.RepeatNever,
			&t.RepeatEndDate,
		); err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}

	return tasks, nil
}


func occursOnDate(task RepeatTask, uiDate time.Time) bool {

	start, err := time.Parse("2006-01-01", task.Date)
	if err != nil {
		return false
	}

	if uiDate.Before(start) {
		return false
	}

	if task.RepeatEndDate != "null" {
		end, err := time.Parse("2006-01-01", task.RepeatEndDate)
		if err == nil && uiDate.After(end) {
			return false
		}
	}

	step := 1
	if task.NumRepeat != "null" {
		fmt.Sscan(task.NumRepeat, &step)
	}

	switch task.Repeat {

	case "day":
		days := int(uiDate.Sub(start).Hours() / 24)
		return days%step == 0

	case "week":
		weeks := int(uiDate.Sub(start).Hours() / (24 * 7))
		return weeks%step == 0

	case "month":
		diff :=
			(uiDate.Year()-start.Year())*12 +
				int(uiDate.Month()-start.Month())
		return diff%step == 0

	case "year":
		diff := uiDate.Year() - start.Year()
		return diff%step == 0
	}

	return false
}



func GetRepeatTasks(db *sql.DB, uiDate time.Time) ([]RepeatTask, error) {

	tasks, err := GetRepeatTasksFilter(db)
	if err != nil {
		return nil, err
	}

	var result []RepeatTask

	for _, task := range tasks {
		if occursOnDate(task, uiDate) {
			result = append(result, task)
		}
	}

	return result, nil
}



// allows me to not have to make multiple add complete task functions 
// for different types of tasks 
func normalize(val string) string {
	if val == "" {
		return "null"
	}
	return val
}


func AddTaskToCompleted(db *sql.DB, t CompletedTask) error {

    CreateCompletedTasksTable(db)

	_, err := db.Exec(`
		INSERT INTO completed_tasks
		(task_id, title, details, time, date, enddate, completeddate, completedtime, instancedate)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`,
		t.ID,
		t.Title,
		t.Details,
		t.Time,
        t.Date,
		t.EndDate,
		t.CompletedDate,
		t.CompletedTime,
        t.InstanceDate,
	)

    if err != nil {
		return err
	}

    return removeCompletedTask( db, t)

}

func removeCompletedTask(db *sql.DB, t CompletedTask) error {
	//today := time.Now().Format("2006-01-02")

	_, err := db.Exec(`
		DELETE FROM tasks
		WHERE id = ?
		  AND repeat = 'null'
		  AND (
		        enddate = 'null'
		     OR enddate = ?
		  );
	`, t.ID, t.InstanceDate)

	return err
}


func FinishRangeTask(db *sql.DB, t CompletedTask) error {

	// ensure table exists
	CreateCompletedTasksTable(db)

	// timestamp
	//now := time.Now()
	//t.CompletedDate = now.Format("2006-01-02")
	//t.CompletedTime = now.Format("15:04")

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		INSERT INTO completed_tasks
		(task_id, title, details, time, date, enddate, completeddate, completedtime, instancedate)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`,
		t.ID,
		t.Title,
		t.Details,
		t.Time,
		t.Date,
		t.EndDate,
		t.CompletedDate,
		t.CompletedTime,
        t.InstanceDate,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

    fmt.Println("FinishRangeTask called for ID:", t.ID, "title: ",t.Title)


	_, err = tx.Exec(`
		DELETE FROM tasks
		WHERE id = ?
	`, t.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}


// this is for id tasks functions that deal with this, i probably need a new schema
// this is going to get super fucked super quickly if the strucutre is off
// this can be a pain or it can work properly ughhhhhhh 

// public method that should get the tasks information based on the id? 
// oh i need to pass along the "instance date" to this don't I?? probably should 
func GetTaskInfo(db *sql.DB, id int ) (*CompletedTask, error){
    
    fmt.Print("I get to get task info")
    // think this should work but really i should be using a new schema 
    // for the sake of speed and simplicity im going to keep using this though
    var t CompletedTask
    
    	err := db.QueryRow(`
		SELECT id, title, details, enddate
		FROM tasks
		WHERE id = ?;
	`, id).Scan(
		&t.ID,
		&t.Title,
		&t.Details,
		&t.EndDate,
	)

    // check if enddate is null, 
    // which means no deadline for now, so we can do some unique cases with this
    // such bad desing, or maybe this is just getting complex 

	if err != nil {
		return nil, err
	}

	fmt.Println("Task:", t)

	return &t, nil
}





// detect task type 
func detectTaskType(date, repeat, enddate string) string {
	if repeat != "null" {
		return "repeat"
	}
	if enddate != "null" {
		return "range"
	}
	return "simple"
}











