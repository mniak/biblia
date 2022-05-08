package main

import (
	"github.com/mniak/biblia/pkg/exporters"
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
		transliterator := hebrew.NewAcademicTransliterator()
		loader := wlc.Loader("sources/UnicodeXML_Westminster_Leningrad_Codex/Tanach")

		testament, err := loader.Load()
		transliteratedTestament := testament.Transliterate(transliterator)
		handle(err)

		exporter := exporters.WriterExporter{
			Writer: text.NewIndentedStdout(),
		}
		exporter.Export(transliteratedTestament)
	},
}
