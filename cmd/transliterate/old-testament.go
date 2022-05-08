package main

import (
	"fmt"

	"github.com/mniak/biblia/pkg/bible"
	"github.com/mniak/biblia/pkg/hebrew"
	"github.com/mniak/biblia/pkg/wlc"
	"github.com/spf13/cobra"
)

var oldTestamentCmd = cobra.Command{
	Use: "old-testament",
	Aliases: []string{
		"ot",
		"old",
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		switch sourceFlag {
		case "wlc":
			loader = wlc.Loader("sources/UnicodeXML_Westminster_Leningrad_Codex/Tanach")
		default:
			return fmt.Errorf("invalid source: %s", sourceFlag)
		}

		switch transliteratorFlag {
		case "academic":
			transliterator = hebrew.AcademicTransliterator()
		default:
			return fmt.Errorf("invalid transliterator: %s", transliteratorFlag)
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := bible.LoadTransliterateAndExport(loader, transliterator, exporter)
		handle(err)
	},
}

func init() {
	oldTestamentCmd.Flags().StringVarP(&sourceFlag, "source", "s", "wlc", "Source text to use (options: wlc)")
	oldTestamentCmd.Flags().StringVarP(&transliteratorFlag, "transliterator", "t", "academic", "Transliterator to use (options: academic)")
}
