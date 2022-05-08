package main

import (
	"log"

	"github.com/mniak/biblia/pkg/bible"
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

var (
	sourceFlag         string
	transliteratorFlag string

	loader         bible.TestamentLoader
	transliterator bible.Transliterator
)

func main() {
	rootCmd.AddCommand(&oldTestamentCmd)
	rootCmd.AddCommand(&newTestamentCmd)

	rootCmd.Execute()
}
