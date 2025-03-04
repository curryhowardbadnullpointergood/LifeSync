package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

//go:embed views/*
var views embed.FS

var t = template.Must(template.ParseFS(views, "views/*"))

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

	})

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	fmt.Println("Listening on Port 3000: ")
	server.ListenAndServe()
}
