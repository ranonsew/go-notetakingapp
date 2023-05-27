package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB // pointer to sql db instance

func Open() error {
	db, err := sql.Open("sqlite3", "./sqlite-notetakingapp-data.db")
	if err != nil {
		return err
	}

	return db.Ping() // connection verification
}

func CreateTable() {
	sql := `CREATE TABLE IF NOT EXISTS mynotes (
		"idNote" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"word" TEXT,
		"definition" TEXT,
		"category" TEXT
	);`

	statement, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("mynotes table created")
}
