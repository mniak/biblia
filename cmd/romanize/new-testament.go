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
	newTestamentCmd.Flags().StringVarP(&sourceFlag, "source", "s", "-", "Source text to use (options: -)")
	newTestamentCmd.Flags().StringVarP(&romanizerFlag, "transliterator", "t", "-", "Transliterator to use (options: -)")
}
