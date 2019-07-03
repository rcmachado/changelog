package cmd

import (
	"bytes"
	"strings"

	"github.com/rcmachado/changelog/chg"
	"github.com/rcmachado/changelog/parser"
	"github.com/spf13/cobra"
)

var addedCmd = &cobra.Command{
	Use:   "added",
	Short: "Add item under 'Added' section",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		input := readChangelog()

		changelog := parser.Parse(input)
		changelog.AddItem(chg.Added, strings.Join(args, " "))

		var buf bytes.Buffer
		changelog.Render(&buf)
		output := buf.Bytes()

		writeChangelog(output)
	},
}

func init() {
	rootCmd.AddCommand(addedCmd)
}
