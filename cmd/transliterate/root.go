package main

import (
	"fmt"
	"log"

	"github.com/mniak/biblia/pkg/bible"
	"github.com/mniak/biblia/pkg/text"
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use: "transliterate",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		switch exporterFlag {
		case "stdout":
			exporter = text.StdoutExporter()
		case "yaml":
			exporter = text.NewYamlExporter(outputDirFlag)
		default:
			return fmt.Errorf("invalid exporter: %s", exporterFlag)
		}

		return nil
	},
}

func handle(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

var (
	sourceFlag         string
	transliteratorFlag string
	exporterFlag       string
	outputDirFlag      string

	loader         bible.TestamentLoader
	transliterator bible.Transliterator
	exporter       bible.Exporter
)

func main() {
	rootCmd.AddCommand(&oldTestamentCmd)
	rootCmd.AddCommand(&newTestamentCmd)

	rootCmd.PersistentFlags().StringVarP(&exporterFlag, "output", "o", "stdout", "The output format/exporter (options: stdout)")
	rootCmd.PersistentFlags().StringVarP(&outputDirFlag, "dir", "d", "./export", "Output directory")

	rootCmd.Execute()
}
