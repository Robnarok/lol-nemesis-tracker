package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

type Entry struct {
	Match string
	Date  string
}

var (
	databasepath string
)

func Init(path string) {
	databasepath = path
}

func CreateDatabase() {
	os.Remove(databasepath) // I delete the file to avoid duplicated records. SQLite is a file based database.

	log.Println("Creating sqlite-database.db...")
	file, err := os.Create(databasepath) // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db created")

	sqliteDatabase, _ := sql.Open("sqlite3", "./"+databasepath) // Open the created SQLite File
	defer sqliteDatabase.Close()                                // Defer Closing the database
	createTable(sqliteDatabase)                                 // Create Database Tables
}

func createTable(db *sql.DB) {
	createStudentTableSQL := `CREATE TABLE entry (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"match" TEXT,
		"date" TEXT
		);` // SQL Statement for Create Table

	statement, err := db.Prepare(createStudentTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
}

// We are passing db reference connection from main to our method with other parameters
func AddEntry(match string, date string) {
	sqliteDatabase, _ := sql.Open("sqlite3", "./"+databasepath)                              // Open the created SQLite File
	statement, err := sqliteDatabase.Prepare(`INSERT INTO entry(match, date) VALUES (?, ?)`) // Prepare statement. This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Printf("%s, %s", match, date)
	queryStatus, err := statement.Exec(match, date)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(queryStatus.RowsAffected())
}

func GetAllEntrys() []Entry {
	sqliteDatabase, _ := sql.Open("sqlite3", "./"+databasepath) // Open the created SQLite File
	entries := []Entry{}
	row, err := sqliteDatabase.Query("SELECT * FROM entry")
	if err != nil {
		log.Fatal(err)
	}
	//defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id string
		var match string
		var date string

		err = row.Scan(&id, &match, &date)
		fmt.Println(err)
		entries = append(entries, Entry{match, date})

	}

	return entries

}
