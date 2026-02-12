package main

import (
    "fmt"
    "net/http"
	"database/sql"
	"encoding/json"
	"time"

)

func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello\n")
}


// api function that handles the add tasks button in the 
// tasks page when you click the plus
func AddTaskButton(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if req.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	title := req.URL.Query().Get("title")
	details := req.URL.Query().Get("details")
	date := req.URL.Query().Get("date")
	repeat := req.URL.Query().Get("repeat")
	time := req.URL.Query().Get("time")
    enddate := req.URL.Query().Get("enddate")
    numrepeat := req.URL.Query().Get("numrepeat")
    repeatoccurrences := req.URL.Query().Get("repeatoccurrences")
    repeatnever := req.URL.Query().Get("repeatnever")
    repeatenddate := req.URL.Query().Get("repeatenddate")
    

	fmt.Println("Adding this task!" )
	fmt.Println("Title: ", title, 
	"\n details: ", details, 
				"\n date: " , date,
				"\n repeat: " , repeat,
				"\n time: ", time,
				"\n repeat end date: " , enddate, 
				"\n Num of repeats: ", numrepeat, 
				"\n Occurrences: " , repeatoccurrences,
				"\n Repeat Never?: " , repeatnever, 
				"\n Repeat end date?: " , repeatenddate)

	if title == "" || details == "" || date == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}



	dbFile := "kala.db"

	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
        	fmt.Println("failed to open db:", err)
	        http.Error(w, "Database error", http.StatusInternalServerError)
	        return
	}
	defer db.Close()


	err = AddTask( db ,title, details, date, repeat, time, enddate, numrepeat, repeatoccurrences, repeatnever, repeatenddate)
	if err != nil {
		fmt.Println("AddTask error:", err)
		http.Error(w, "Insert Error", http.StatusInternalServerError)
		return 
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Task Added!\n"))
}



func GetSimpleTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    
	db, err := sql.Open("sqlite", "./kala.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	tasks, err := GetTaskSimple(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

func GetRangeTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db, err := sql.Open("sqlite", "./kala.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	dateStr := r.URL.Query().Get("date")
	if dateStr == "" {
		http.Error(w, "missing date param", http.StatusBadRequest)
		return
	}

	uiDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "invalid date format", http.StatusBadRequest)
		return
	}

	tasks, err := GetTaskRange(db, uiDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

func CompleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	// CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// preflight
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var task CompletedTask
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// calculate completed date & time HERE
	now := time.Now()
	task.CompletedDate = now.Format("2006-01-02") // YYYY-MM-DD
	task.CompletedTime = now.Format("15:04")      // HH:MM

	db, err := sql.Open("sqlite", "./kala.db")
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	err = AddTaskToCompleted(db, task)
	if err != nil {
		fmt.Println("Insert failed:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task completed"))
}

func FinishRangeTaskHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var task CompletedTask
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Bad JSON", http.StatusBadRequest)
		return
	}

	db, err := sql.Open("sqlite", "./kala.db")
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	if err := FinishRangeTask(db, task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Range task finished"))
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
	http.HandleFunc("/tasks/simple", GetSimpleTasksHandler)
	http.HandleFunc("/tasks/range", GetRangeTasksHandler)
    http.HandleFunc("/tasks/complete", CompleteTaskHandler)
    http.HandleFunc("/tasks/completerange", FinishRangeTaskHandler)


    http.ListenAndServe(":8090", nil)
}
