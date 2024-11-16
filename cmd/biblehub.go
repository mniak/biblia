package main

import "github.com/spf13/cobra"

func BibleHub() *cobra.Command {
	cmd := &cobra.Command{
		Use: "biblehub",
	}

	cmdDownload := &cobra.Command{
		Use: "download",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	cmd.AddCommand(cmdDownload)
	return cmd
}
