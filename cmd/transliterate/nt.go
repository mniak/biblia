package main

import (
	"log"

	"github.com/spf13/cobra"
)

var ntCmd = cobra.Command{
	Use: "nt",
	Aliases: []string{
		"new",
	},
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatalln("new testament/greek not yet implemented")

		// loader := wlc.Loader("sources/UnicodeXML_Westminster_Leningrad_Codex/Tanach")
		// transliterator := hebrew.AcademicTransliterator()
		// exporter := text.StdoutExporter()

		// handle(bible.LoadTransliterateAndExport(
		// 	loader,
		// 	transliterator,
		// 	exporter,
		// ))
	},
}
