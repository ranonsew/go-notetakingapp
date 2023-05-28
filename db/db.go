package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Note struct {
	id int
	content string
}

var db *sql.DB // pointer to sql db instance

func Open() error {
	var err error // to ensure we use the above "db" (not creating new)
	db, err = sql.Open("sqlite3", "./sqlite-notetakingapp-data.db")
	if err != nil {
		return err
	}

	return db.Ping() // connection verification
}

func CreateTable() {
	query := `CREATE TABLE IF NOT EXISTS notes (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		content TEXT
	);`

	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer statement.Close()

	statement.Exec()
	log.Println("mynotes table created")
}

func ListNotes() {
	query := `SELECT * FROM notes`

	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer statement.Close()

	rows, err := statement.Query()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var note Note
		err := rows.Scan(&note.id, &note.content)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Println(note)
	}
}

func ReadNote(id int) {
	query := `SELECT * FROM notes WHERE id = ? LIMIT 1`

	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer statement.Close()

	row := statement.QueryRow(id)

	var note Note
	err = row.Scan(&note.id, &note.content)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal("Error: User not found")
		}
		log.Fatal(err.Error())
	}
	log.Println(note)
	// something about nano and write out SQLite3 data into md for nano to read
}

func CreateNote(content string) {
	query := `INSERT INTO notes (content) VALUES (?)`
	
	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := statement.Exec(content)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Note created successfully!", result)
	// something about nano and open the md for nano to write into, which then saves into SQLite3
}

func UpdateNote(id int) {
	query := `INSERT INTO notes ()`
}

func DeleteNote(id int) {}
