package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
    "database/sql"

	"github.com/mattn/go-sqlite3"
)

//go:embed views/*
var views embed.FS

var t = template.Must(template.ParseFS(views, "views/*"))



// struct to represent the tasks
type Tasks struct {
    ID int
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
      title TEXT
     );`

     _, err = DB.Exec(sqlStmt)
     if err != nil {
      log.Fatalf("Error creating table: %q: %s\n", err, sqlStmt) // Log an error if table creation fails
     }

}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if err := t.ExecuteTemplate(w, "index.html", nil); err != nil {
			http.Error(w, "Something went wrong, as usual", http.StatusInternalServerError)
		}
	})

	// Handle add task form submission
	router.HandleFunc("POST /addtask", func(w http.ResponseWriter, r *http.Request) {

		log.Println("Hello")
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Unable to parse task form: ", http.StatusInternalServerError)
			return
		}

		task := strings.ToLower(r.FormValue("taskInfo"))

		log.Printf("Task to add: %s", task)
        // initiates the database 
        initDB()
        // insert task into the database 
        _,err := 



	})

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	fmt.Println("Listening on Port 3000: ")
	server.ListenAndServe()
}
