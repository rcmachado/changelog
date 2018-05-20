package cmd

import (
	"fmt"
	"os"

	"github.com/rcmachado/keepachangelog/parser"
	"github.com/spf13/cobra"
)

var fmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "Reformat the change log file",
	Long:  "Reformats changelog input following keepachangelog.com spec",
	Run: func(cmd *cobra.Command, args []string) {
		input, err := readChangelog(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		output := parser.NewReader(input)
		fmt.Printf("%s", output)
	},
}

func init() {
	rootCmd.AddCommand(fmtCmd)
}
