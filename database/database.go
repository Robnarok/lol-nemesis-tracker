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
type TrackedSummoners struct {
	Name    string
	Nemesis string
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
	createEntryTable(sqliteDatabase)
	createSummonerTable(sqliteDatabase)
}
func createSummonerTable(db *sql.DB) {
	createSummonerTableSQL := `CREATE TABLE trackedSummoner (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT,
		"nemesis" TEXT
		);` // SQL Statement for Create Table

	statement, err := db.Prepare(createSummonerTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
}

func createEntryTable(db *sql.DB) {
	createEntryTableSQL := `CREATE TABLE entry (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"match" TEXT,
		"date" TEXT
		);` // SQL Statement for Create Table

	statement, err := db.Prepare(createEntryTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
}

func AddSummoner(name string, nemesis string) {
	sqliteDatabase, _ := sql.Open("sqlite3", "./"+databasepath)                                          // Open the created SQLite File
	statement, err := sqliteDatabase.Prepare(`INSERT INTO trackedSummoner(name, nemesis) VALUES (?, ?)`) // Prepare statement. This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Printf("%s, %s", name, nemesis)
	queryStatus, err := statement.Exec(name, nemesis)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(queryStatus.RowsAffected())
}

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

func GetAllSummoners() []Entry {
	sqliteDatabase, _ := sql.Open("sqlite3", "./"+databasepath) // Open the created SQLite File
	entries := []Entry{}
	row, err := sqliteDatabase.Query("SELECT * FROM trackedsummoners")
	if err != nil {
		log.Fatal(err)
	}
	//defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id string
		var name string
		var nemesis string

		err = row.Scan(&id, &name, &nemesis)
		entries = append(entries, Entry{name, nemesis})

	}

	return entries

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
		entries = append(entries, Entry{match, date})

	}

	return entries

}
