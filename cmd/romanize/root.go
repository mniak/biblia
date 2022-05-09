package main

import (
	"log"

	"github.com/mniak/biblia/internal/flagutils"
	"github.com/mniak/biblia/pkg/bible"
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use: "romanize",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		exporter, err = flagutils.Exporter(exporterFlag, outputDirFlag)
		return err
	},
}

func handle(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

var (
	sourceFlag    string
	romanizerFlag string
	exporterFlag  string
	outputDirFlag string

	loader    bible.TestamentLoader
	romanizer bible.Romanizer
	exporter  bible.Exporter
)

func main() {
	rootCmd.AddCommand(&oldTestamentCmd)
	rootCmd.AddCommand(&newTestamentCmd)

	rootCmd.PersistentFlags().StringVar(&exporterFlag, "output", "stdout", "The output format/exporter (options: stdout)")
	rootCmd.PersistentFlags().StringVar(&outputDirFlag, "output-dir", "./export", "Output directory")

	rootCmd.Execute()
}
