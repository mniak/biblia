package main

import "github.com/spf13/cobra"

func main() {
	root := cobra.Command{
		Use: "biblehub",
	}
	root.AddCommand(DownloadCmd())
	root.Execute()
}

func DownloadCmd() *cobra.Command {
	cmd := cobra.Command{
		Use: "download",
		Run: func(cmd *cobra.Command, args []string) {
			// biblehub.
		},
	}
	return &cmd
}
