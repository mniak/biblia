package transliterate

import (
	"fmt"

	"github.com/mniak/biblia/pkg/bible"
	"github.com/mniak/biblia/pkg/text"
	"github.com/spf13/cobra"
)

var (
	sourceFlag         string
	transliteratorFlag string
	exporterFlag       string

	loader         bible.TestamentLoader
	transliterator bible.Transliterator
	exporter       bible.Exporter
)

func Command() *cobra.Command {
	cmd := cobra.Command{
		Use: "transliterate",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			switch exporterFlag {
			case "stdout":
				exporter = text.StdoutExporter()
			default:
				return fmt.Errorf("invalid exporter: %s", exporterFlag)
			}

			return nil
		},
	}

	cmd.AddCommand(&oldTestamentCmd)
	cmd.AddCommand(&newTestamentCmd)

	cmd.PersistentFlags().StringVarP(&exporterFlag, "exporter", "e", "stdout", "The exporter to use (options: stdout)")

	return &cmd
}
