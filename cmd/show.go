package cmd

import (
	"fmt"
	"os"

	"github.com/rcmachado/changelog/parser"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show [version]",
	Short: "Show changelog for [version]",
	Long:  `Show changelog section and entries for version [version]`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		changelog := parser.Parse(inputFile)

		v := changelog.Version(args[0])
		if v == nil {
			fmt.Printf("Unknown version: '%s'\n", args[0])
			os.Exit(3)
		}

		v.RenderChanges(outputFile)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
