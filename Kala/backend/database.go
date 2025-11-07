package main

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)



// creates all of the tables for the calendar application 
func createTable(db *sql.DB) {
	create := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
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


func main() {
	log.Println("Creating kala.db...")
	file, err := os.Create("./kala.db")
	if err != nil {
		log.Fatal("Failed to create DB file:", err)
	}
	file.Close()
	log.Println("kala.db created")

	db, err := sql.Open("sqlite", "./kala.db")
	if err != nil {
		log.Fatal("Failed to open DB:", err)
	}
	defer db.Close()

	createTable(db)

	log.Println("Database initialized successfully!")
}

