package cmd

import (
	"bytes"
	"fmt"
	"os"

	"github.com/rcmachado/changelog/parser"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show information about version",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		input, err := readChangelog(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		chg := parser.Parse(input)

		v := chg.Version(args[0])
		if v == nil {
			fmt.Printf("Unknown version %s\n", args[0])
			os.Exit(3)
		}

		var buf bytes.Buffer
		v.Render(&buf)
		output := buf.Bytes()

		fmt.Printf("%s", output)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
