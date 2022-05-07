package main

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use: "transliterate",
}

func handle(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	rootCmd.AddCommand(&otCmd)
	rootCmd.AddCommand(&ntCmd)

	rootCmd.Execute()
}
