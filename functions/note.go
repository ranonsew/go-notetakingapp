package functions

import ()

func OpenNote(id int) {
	// should open up the text editor through this function, to open up the note
	// will probably include
		// pass in "id" from list.go
		// run "db" ReadNote(id) to select the row and return the "db" Note{}
		// take note.Content and write it to a markdown file in either the same dir as the db, or in temp directory
			// the markdown file should be named the same by note.Name
		// note content should then be displayed out
		// If edited and saved, should update using "db" UpdateNote(id, content) based on the new content from the file
}
