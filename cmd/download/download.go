package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/mniak/biblia/pkg/bible"
	"github.com/mniak/biblia/pkg/sources/biblehub"
	"github.com/mniak/biblia/pkg/text"
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:  "download <testament>",
	Args: cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		testament := strings.ToLower(args[0])
		if testament != "ot" && testament != "nt" {
			return fmt.Errorf("invalid testament '%s'", args[0])
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		testament := strings.ToLower(args[0])
		exporter := text.TomlExporter("./sources/BibleHub/Interlinear/OldTestament")

		switch testament {
		case "ot":
			loader := biblehub.NewInterlinearOldTestamentLoader()
			err := bible.LoadAndExport(loader, exporter)
			handle(err)
		case "nt":
			log.Fatalln("New Testament download not yet implemented")
		}
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
	rootCmd.Execute()
}
