package cmd

import (
	"bytes"
	"fmt"

	"github.com/rcmachado/changelog/parser"
	"github.com/spf13/cobra"
)

var fmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "Reformat the change log file",
	Long:  "Reformats changelog input following keepachangelog.com spec",
	Run: func(cmd *cobra.Command, args []string) {
		input := readChangelog()

		chg := parser.Parse(input)

		var buf bytes.Buffer
		chg.Render(&buf)
		output := buf.Bytes()

		fmt.Printf("%s", output)
	},
}

func init() {
	rootCmd.AddCommand(fmtCmd)
}
