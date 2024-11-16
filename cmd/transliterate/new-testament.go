package transliterate

import (
	"errors"
	"fmt"

	"github.com/mniak/biblia/pkg/bible"
	"github.com/mniak/biblia/pkg/greek"
	"github.com/mniak/biblia/pkg/sources/byzantine"
	"github.com/spf13/cobra"
)

var newTestamentCmd = cobra.Command{
	Use: "new-testament",
	Aliases: []string{
		"nt",
		"new",
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		switch sourceFlag {
		case "byzantine":
			loader = byzantine.Loader("sources/robinson_byzantine-majority-text_b844fff/csv-unicode/ccat/no-variants")
		default:
			return fmt.Errorf("invalid source: %s", sourceFlag)
		}

		switch transliteratorFlag {
		case "noop":
			transliterator = greek.NoopTransliterator()
		case "simple":
			return errors.New("simple transliterator: not yet implemented")
			transliterator = greek.SimpleTransliterator()
		default:
			return fmt.Errorf("invalid transliterator: %s", transliteratorFlag)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := bible.LoadTransliterateAndExport(loader, transliterator, exporter)
		cobra.CheckErr(err)
	},
}

func init() {
	newTestamentCmd.Flags().StringVarP(&sourceFlag, "source", "s", "-", "Source text to use (options: byzantine)")
	newTestamentCmd.Flags().StringVarP(&transliteratorFlag, "transliterator", "t", "noop", "Transliterator to use (options: noop, simple)")
}
