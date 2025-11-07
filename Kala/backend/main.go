package main

import (
    "fmt"
    "net/http"

)

func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello\n")
}


// api function that handles the add tasks button in the 
// tasks page when you click the plus
func AddTaskButton(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	title := req.URL.Query().Get("title")
	details := req.URL.Query().Get("details")
	date := req.URL.Query().Get("date")
	repeat := req.URL.Query().Get("repeat")

	if title == "" || details == "" || date == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	err := AddTask(title, details, date, repeat)
	if err != nil {
		fmt.Println("AddTask error:", err)
		http.Error(w, "Insert Error", http.StatusInternalServerError)
		return 
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Task Added!\n"))
}


func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {

    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)
	http.HandleFunc("/addtask", AddTaskButton)

    http.ListenAndServe(":8090", nil)
}
