package main

import (
	"log"

	"github.com/spf13/cobra"
)

var newTestamentCmd = cobra.Command{
	Use: "new-testament",
	Aliases: []string{
		"nt",
		"new",
	},
	Deprecated: "not yet implemented",
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatalln("new testament/greek not yet implemented")
	},
}

func init() {
	newTestamentCmd.Flags().StringVar(&sourceFlag, "source", "-", "Source text to use (options: -)")
	newTestamentCmd.Flags().StringVar(&romanizerFlag, "romanizer", "-", "Romanizer to use (options: -)")
}
