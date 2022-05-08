package main

import (
	"github.com/mniak/biblia/pkg/bible"
	"github.com/mniak/biblia/pkg/hebrew"
	"github.com/mniak/biblia/pkg/sources/wlc"
	"github.com/mniak/biblia/pkg/text"
	"github.com/spf13/cobra"
)

var otCmd = cobra.Command{
	Use: "ot",
	Aliases: []string{
		"old",
		"tanach",
	},
	Run: func(cmd *cobra.Command, args []string) {
		loader := wlc.Loader("sources/UnicodeXML_Westminster_Leningrad_Codex/Tanach")
		transliterator := hebrew.AcademicTransliterator()
		exporter := text.StdoutExporter()

		handle(bible.LoadTransliterateAndExport(
			loader,
			transliterator,
			exporter,
		))
	},
}
