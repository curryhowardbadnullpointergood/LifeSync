package main

import (
	"database/sql"
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed views/*
var views embed.FS

// template to more like index html
var indexh *template.Template

var t = template.Must(template.ParseFS(views, "views/*"))

// struct to represent the tasks
type Tasks struct {
	ID     int
	Task   string
	Points int
}

// Db is a global var for the sqlite database connection
var DB *sql.DB

func initDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}

	// SQL statement to create the todos table if it doesn't exist
	sqlStmt := `
     CREATE TABLE IF NOT EXISTS todos (
      id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
      task TEXT,
	  points INTEGER
     );`

	_, err = DB.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("Error creating table: %q: %s\n", err, sqlStmt) // Log an error if table creation fails
	}

}

func main() {

	initDB()
	defer DB.Close()
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if err := t.ExecuteTemplate(w, "index.html", nil); err != nil {
			http.Error(w, "Something went wrong, as usual", http.StatusInternalServerError)
		}

	})

	// Handle add task form submission
	router.HandleFunc("POST /addtask", addTask)

	// Handle deleting a task
	// this also needs to initiate a done tasks table, which take the done task and adds it to this table
	// needs to be a button in the future ugh hopefully???
	router.HandleFunc("POST /deletetask", deleteTask)

	// this handles getting the task, really should break this down and make it cleaner
	router.HandleFunc("GET /gettask", func(w http.ResponseWriter, r *http.Request) {

		log.Println("Trying to get tasks from database: ")
		query := "SELECT * FROM todos"
		rows, err := DB.Query(query)

		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var tasks []Tasks
		var responseHTML string

		for rows.Next() {
			var todo Tasks
			rowErr := rows.Scan(&todo.ID, &todo.Task, &todo.Points)
			if rowErr != nil {
				log.Fatal(err)
			}
			tasks = append(tasks, todo)
			responseHTML += fmt.Sprintf("<p>%d: %s: %d</p>", todo.ID, todo.Task, todo.Points)
		}
		if err = rows.Err(); err != nil {
			log.Fatal(err)
		}
		log.Println(tasks)

		for _, t := range tasks {
			fmt.Println(t.ID, t.Task, t.Points)
		}

		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(responseHTML))

	})

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	fmt.Println("Listening on Port 3000: ")
	server.ListenAndServe()
}

func addTask(w http.ResponseWriter, r *http.Request) {

	//fmt.Println("hello")

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse task form: ", http.StatusInternalServerError)
		return
	}

	task := strings.ToLower(r.FormValue("taskInfo"))
	pointsStr := r.FormValue("points")     // This gets the value as a string
	points, err := strconv.Atoi(pointsStr) // Convert to integer
	if err != nil {
		http.Error(w, "Invalid points value", http.StatusBadRequest)
		return
	}

	log.Printf("Task to add: %s : And points: %d", task, points)

	// insert task into the database
	_, err = DB.Exec("INSERT INTO todos VALUES(NULL,?,?)", task, points)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Return an HTTP 500 error if insertion fails
		return
	}

	w.Header().Set("HX-Refresh", "true")
	w.WriteHeader(http.StatusOK)

}

func deleteTask(w http.ResponseWriter, r *http.Request) {

	//fmt.Println("hello")

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse task form: ", http.StatusInternalServerError)
		return
	}

	idStr := (r.FormValue("taskID"))
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid points value", http.StatusBadRequest)
		return
	}

	// so this gets the task and poin values
	var taskD string
	var pointsD int

	err = DB.QueryRow("SELECT task, points FROM todos WHERE id = ?", id).Scan(&taskD, &pointsD)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No task found with the given ID")
		} else {
			log.Fatal(err)
		}
		return
	}

	// Print the retrieved values
	fmt.Printf("Task: %s, Points: %d\n", taskD, pointsD)

	log.Printf("ID to delete: %d", id)

	// delete task from the database
	_, err = DB.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// making a table for done tasks

	// SQL statement to create the todos table if it doesn't exist
	sqlStmt := `
     CREATE TABLE IF NOT EXISTS done (
      id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
      task TEXT,
	  points INTEGER
     );`

	_, err = DB.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("Error creating table: %q: %s\n", err, sqlStmt) // Log an error if table creation fails
	}

	_, err = DB.Exec("INSERT INTO done VALUES(NULL,?,?)", taskD, pointsD)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Return an HTTP 500 error if insertion fails
		return
	}

	// now add the deleted task to  this table, so to do this I need to query for the task and point value using the id before deleteing it and storing that information in 2 variables then adding that to this table.

	w.Header().Set("HX-Refresh", "true")
	w.WriteHeader(http.StatusOK)

}
