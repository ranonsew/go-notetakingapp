package db

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
	fn "github.com/ranonsew/go-notetakingapp/functions"
)

type Note struct {
	Id int `sql:"id"`
	Name string `sql:"name"`
	Content string `sql:"content"`
	LastUpdated time.Time `sql:"last_updated"`
}

// type NoteRow struct {
// 	Note
// 	LastUpdated time.Time `sql:"last_updated"`
// }

var db *sql.DB // pointer to sql db instance

func Open() error {
	var err error // to ensure we use the above "db" (not creating new)
	db, err = sql.Open("sqlite3", "./data/sqlite-notetakingapp.db")
	if err != nil {
		return err
	}

	return db.Ping() // connection verification
}

// creating a new "notes" table if it doesn't already exist
func CreateTable() {
	query := `CREATE TABLE IF NOT EXISTS notes (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		content TEXT,
		last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	statement, err := db.Prepare(query)
	fn.CheckError(err)
	defer statement.Close()

	statement.Exec()
	log.Println("'notes' table created")
}

// dropping the "notes" table if it exists (for testing or smth idk)
func DropTable() {
	query := `DROP TABLE IF EXISTS notes`

	stmt, err := db.Prepare(query)
	fn.CheckError(err)
	defer stmt.Close()

	stmt.Exec()
	log.Println("'notes' table dropped")
}

// list all available notes
func ListNotes() []string {
	titles := []string{} // note titles
	query := `SELECT * FROM notes`

	statement, err := db.Prepare(query)
	fn.CheckError(err)
	defer statement.Close()

	rows, err := statement.Query()
	fn.CheckError(err)
	defer rows.Close()

	for rows.Next() {
		var note Note

		err := rows.Scan(&note.Id, &note.Name, &note.Content, &note.LastUpdated)
		fn.CheckError(err)
		log.Println(note)

		// title == "[id] name (date)"
		titles = append(titles, "[" + strconv.Itoa(note.Id) + "] " + note.Name + " (" + note.LastUpdated.Format("Mon, 02 Jan 2006 15:04:05 +0800") + ")")
	}

	return titles
}

// get a single note and display its contents (to be available for editing)
func ReadNote(id int) Note {
	var note Note
	query := `SELECT * FROM notes WHERE id = ? LIMIT 1`

	statement, err := db.Prepare(query)
	fn.CheckError(err)
	defer statement.Close()

	row := statement.QueryRow(id)
	err = row.Scan(&note.Id, &note.Name, &note.Content, &note.LastUpdated)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal("Error: Note not found in the database")
		}
		log.Fatal(err.Error())
	}

	log.Println(note)
	return note
}

// Insert a new note to the "notes" table
func InsertNote(name string, content string) {
	query := `INSERT INTO notes (name, content, last_updated) VALUES (?, ?, ?)`

	tx, err := db.Begin()
	fn.CheckError(err)

	statement, err := tx.Prepare(query)
	fn.CheckError(err)

	result, err := statement.Exec(name, content, time.Now())
	if err != nil {
		tx.Rollback()
		log.Fatal(err.Error())
	}

	err = tx.Commit()
	fn.CheckError(err)
	log.Println("Note created & saved successfully!", result)

	id, err := result.LastInsertId()
	fn.CheckError(err)
	log.Println(id)
	// here we can run "ReadNote()" to open up the note for us
}

// to update the note that exists, when edit and save the note, this runs (essentially the save button function)
func UpdateNote(id int, name string, content string) {
	if name == "" && content == "" {
		log.Fatal(errors.New("need to update either name, content, or both").Error())
	}

	args := []any{} // arguments allowing for any data type
	query := `UPDATE notes SET `
	if name != "" {
		query += `name = ? `
		args = append(args, name)
	}
	if content != "" {
		query += `content = ? `
		args = append(args, content)
	}
	query += `last_updated = ? WHERE id = ?`
	args = append(args, time.Now(), id)

	tx, err := db.Begin()
	fn.CheckError(err)

	stmt, err := tx.Prepare(query)
	fn.CheckError(err)
	defer stmt.Close()

	result, err := stmt.Exec(args...)
	fn.CheckError(err)
	if err != nil {
		tx.Rollback()
		log.Fatal(err.Error())
	}

	err = tx.Commit()
	fn.CheckError(err)
	log.Println("Note updated successfully!", result)
}

// for deleting a note from the "notes" table
func DeleteNote(id int) {
	query := `DELETE FROM notes WHERE id = ?`

	tx, err := db.Begin()
	fn.CheckError(err)

	stmt, err := tx.Prepare(query)
	fn.CheckError(err)
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		tx.Rollback()
		log.Fatal(err.Error())
	}

	err = tx.Commit()
	fn.CheckError(err)
	log.Println("Note deleted successfully!", result)
}
