package transliterate

import (
	"fmt"

	"github.com/mniak/biblia/pkg/bible"
	"github.com/mniak/biblia/pkg/greek"
	"github.com/mniak/biblia/pkg/hebrew"
	"github.com/mniak/biblia/pkg/sources/byzantine"
	"github.com/mniak/biblia/pkg/sources/wlc"
	"github.com/mniak/biblia/pkg/text"
	"github.com/spf13/cobra"
)

var flags struct {
	sourceFlag     string
	transliterator string
	exporterFlag   string
}

func Command() *cobra.Command {

	var loader bible.TestamentLoader
	var transliterator bible.Transliterator
	var exporter bible.Exporter

	cmd := cobra.Command{
		Use: "transliterate",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			switch flags.exporterFlag {
			case "stdout":
				exporter = text.StdoutExporter()
			default:
				return fmt.Errorf("invalid exporter: %s", flags.exporterFlag)
			}

			return nil
		},
	}
	cmd.PersistentFlags().StringVarP(&flags.exporterFlag, "exporter", "e", "stdout", "The exporter to use (options: stdout)")

	otCmd := cobra.Command{
		Use: "old-testament",
		Aliases: []string{
			"ot",
			"old",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			switch flags.sourceFlag {
			case "wlc":
				loader = wlc.Loader("sources/UnicodeXML_Westminster_Leningrad_Codex/Tanach")
			default:
				return fmt.Errorf("invalid source: %s", flags.sourceFlag)
			}

			switch flags.transliterator {
			case "academic":
				transliterator = hebrew.AcademicTransliterator()
			case "simple":
				transliterator = hebrew.SimpleTransliterator()
			default:
				return fmt.Errorf("invalid transliterator: %s", flags.transliterator)
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			err := bible.LoadTransliterateAndExport(loader, transliterator, exporter)
			cobra.CheckErr(err)
		},
	}
	otCmd.Flags().StringVarP(&flags.sourceFlag, "source", "s", "wlc", "Source text to use (options: wlc)")
	otCmd.Flags().StringVarP(&flags.transliterator, "transliterator", "t", "academic", "Transliterator to use (options: academic, simple)")
	cmd.AddCommand(&otCmd)

	ntCmd := cobra.Command{
		Use: "new-testament",
		Aliases: []string{
			"nt",
			"new",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			switch flags.sourceFlag {
			case "byzantine":
				loader = byzantine.Loader("sources/robinson_byzantine-majority-text_b844fff/csv-unicode/ccat/no-variants")
			default:
				return fmt.Errorf("invalid source: %s", flags.sourceFlag)
			}

			switch flags.transliterator {
			case "noop":
				transliterator = greek.NoopTransliterator()
			case "simple":
				transliterator = greek.SimpleTransliterator()
			default:
				return fmt.Errorf("invalid transliterator: %s", flags.transliterator)
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			err := bible.LoadTransliterateAndExport(loader, transliterator, exporter)
			cobra.CheckErr(err)
		},
	}
	ntCmd.Flags().StringVarP(&flags.sourceFlag, "source", "s", "-", "Source text to use (options: byzantine)")
	ntCmd.Flags().StringVarP(&flags.transliterator, "transliterator", "t", "noop", "Transliterator to use (options: noop, simple)")
	cmd.AddCommand(&ntCmd)

	return &cmd
}
