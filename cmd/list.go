/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/ranonsew/go-notetakingapp/functions"
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
		fmt.Println("list called")

		functions.PromptGetSelect(functions.PromptContent{
			ErrorMsg: "A note must be selected to continue.",
			Label: "Select a note to continue",
		}, db.ListNotes())
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
