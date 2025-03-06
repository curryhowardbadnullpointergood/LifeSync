package main

import (
	"database/sql"
	"embed"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
	"strings"
)

//go:embed views/*
var views embed.FS

var t = template.Must(template.ParseFS(views, "views/*"))

// struct to represent the tasks
type Tasks struct {
	ID   int
	Task string
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
      task TEXT
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
	router.HandleFunc("POST /addtask", func(w http.ResponseWriter, r *http.Request) {

		// log.Println("Hello")
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Unable to parse task form: ", http.StatusInternalServerError)
			return
		}

		task := strings.ToLower(r.FormValue("taskInfo"))

		log.Printf("Task to add: %s", task)

		// insert task into the database
		_, err := DB.Exec("INSERT INTO todos VALUES(NULL,?)", task)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) // Return an HTTP 500 error if insertion fails
			return
		}

		// redirect to the main page after a sucessful creation of the task?
		// http.Redirect(w, r, "/", http.StatusSeeOther)

	})

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	fmt.Println("Listening on Port 3000: ")
	server.ListenAndServe()
}
