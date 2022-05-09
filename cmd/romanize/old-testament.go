package main

import (
	"fmt"

	"github.com/mniak/biblia/internal/flagutils"
	"github.com/mniak/biblia/pkg/bible"
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

		var err error
		romanizer, err = flagutils.Romanizer(bible.LangHebrew, romanizerFlag)
		if err != nil {
			return err
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := bible.LoadRomanizeAndExport(loader, romanizer, exporter)
		handle(err)
	},
}

func init() {
	oldTestamentCmd.Flags().StringVar(&sourceFlag, "source", "wlc", "Source text to use (options: wlc)")
	oldTestamentCmd.Flags().StringVar(&romanizerFlag, "romanizer", "academic-hebrew", "Romanizer to use (options: academic)")
}
