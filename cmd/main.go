package main

import (
	"github.com/mniak/biblia/cmd/transliterate"
	"github.com/spf13/cobra"
)

func main() {
	cmd := cobra.Command{
		Use: "biblia",
	}
	cmd.AddCommand(
		transliterate.Command(),
		BibleHub(),
	)
	cmd.Execute()
}
