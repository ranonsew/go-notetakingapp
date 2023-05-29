package db

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Note struct {
	Id int `sql:"id"`
	Name string `sql:"name"`
	Content string `sql:"content"`
	LastUpdated time.Time `sql:"last_updated"`
}

var db *sql.DB // pointer to sql db instance

func Open() error {
	var err error // to ensure we use the above "db" (not creating new)
	db, err = sql.Open("sqlite3", "./data/sqlite-notetakingapp.db")
	if err != nil {
		return err
	}

	return db.Ping() // connection verification
}

func CreateTable() {
	query := `CREATE TABLE IF NOT EXISTS notes (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		content TEXT,
		last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer statement.Close()

	statement.Exec()
	log.Println("'notes' table created")
}

func ListNotes() []string {
	titles := []string{} // note titles
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
		err := rows.Scan(&note.Id, &note.Name, &note.Content, &note.LastUpdated)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Println(note)

		// title == "[id] name (date)"
		titles = append(titles, "[" + strconv.Itoa(note.Id) + "] " + note.Name + " (" + note.LastUpdated.Format("Mon, 02 Jan 2006 15:04:05 +0800") + ")")
	}

	return titles
}

func ReadNote(id int) Note {
	var note Note
	query := `SELECT * FROM notes WHERE id = ? LIMIT 1`

	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer statement.Close()

	row := statement.QueryRow(id)
	err = row.Scan(&note.Id, &note.Name, &note.Content, &note.LastUpdated)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal("Error: User not found")
		}
		log.Fatal(err.Error())
	}

	log.Println(note)
	return note
}

func InsertNote(name string, content string) {
	query := `INSERT INTO notes (name, content, last_updated) VALUES (?, ?, ?)`

	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := statement.Exec(name, content, time.Now())
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Note created & saved successfully!", result)
	// something about nano and open the md for nano to write into, which then saves into SQLite3
}

func UpdateNote(id int, name string, content string) {
	query := `UPDATE notes SET name = ?, content = ?, last_updated = ? WHERE id = ?`

	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := statement.Exec(name, content, time.Now(), id)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Note updated successfully!", result)
}

func DeleteNote(id int) {
	query := `DELETE FROM notes WHERE id = ?`

	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := statement.Exec(id)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Note deleted successfully!", result)
}
