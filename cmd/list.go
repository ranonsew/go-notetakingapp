/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"strings"
	"strconv"

	fn "github.com/ranonsew/go-notetakingapp/functions"
	"github.com/ranonsew/go-notetakingapp/db"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists out all the saved notes",
	Long: `All the notes that are saved within the sqlite database
	are displayed in the form of a list of notes, 
	in which the user can then select.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("list called")

		selectedNote := fn.PromptGetSelect(fn.PromptContent{
			ErrorMsg: "A note must be selected to continue.",
			Label: "Select a note to continue",
		}, db.ListNotes())

		// using selected, grab out the [id] and use the "db" ReadNote() to grab the note to display (this part should take from "functions")
		// now need to get the thing staring from the second index to one before the fourth index
		idx := strings.Index(selectedNote, "]") // "[" is always index 0
		id, err := strconv.Atoi(selectedNote[1:idx]) // from after the "[" to before the "]"
		fn.CheckError(err)
		db.ReadNote(id) // open the note that was selected
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
