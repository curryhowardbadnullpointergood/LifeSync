package main 


import (
 "database/sql"   // Package for SQL database interactions
 "fmt"            // Package for formatted I/O
 "log"            // Package for logging
 "net/http"       // Package for HTTP client and server
 "text/template"  // Package for HTML templates

 _ "github.com/mattn/go-sqlite3" // SQLite driver
)

// This represents a task in the task list
// Id is a unique identifier 
// Title = title of the task 
// Priority is the priority of the task in question 
// there needs to be a couple more things added to this like starred, task list name (which is like a list for every 
// specific type of task, ie one for work, one for gym/fitness one for cooking etc. 
// points is the number of points you get for completing a task 
type Task struct {
 ID    int
 Title string
 Priority int 
 Starred bool 
 Points int 
}

// struct to represent the contacts
type Contacts struct {
	ID          int
	Name        string
	Description string
}

// DB is a global variable for the SQLite database connection
var DB *sql.DB


// initDB initializes the SQLite database and creates the todos table if it doesn't exist
func initDB() {
 var err error
 DB, err = sql.Open("sqlite3", "./kala.db") // Open a connection to the SQLite database file named app.db
 if err != nil {
  log.Fatal(err) // Log an error and stop the program if the database can't be opened
 }

 // SQL statement to create the todos table if it doesn't exist
 sqlStmt := `
 CREATE TABLE IF NOT EXISTS tasks (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  title TEXT   
  priority INTEGER 
  starred BOOL 
  points INTEGER

 );`

 _, err = DB.Exec(sqlStmt)
 if err != nil {
  log.Fatalf("Error creating table: %q: %s\n", err, sqlStmt) // Log an error if table creation fails
 }
}




