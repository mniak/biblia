package main

import "github.com/spf13/cobra"

var ntCmd = cobra.Command{
	Use: "nt",
	Aliases: []string{
		"new",
	},
	Run: func(cmd *cobra.Command, args []string) {},
}
