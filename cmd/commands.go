package cmd

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/rcmachado/changelog/chg"
	"github.com/rcmachado/changelog/parser"
	"github.com/spf13/cobra"
)

func buildCommands(rootCmd *cobra.Command) {
	cmdTypes := []chg.ChangeType{
		chg.Added, chg.Changed, chg.Deprecated, chg.Fixed, chg.Removed, chg.Security,
	}

	for _, changeType := range cmdTypes {
		cmdType := changeType
		cmd := &cobra.Command{
			Use:   strings.ToLower(cmdType.String()),
			Short: fmt.Sprintf("Add item under '%s' section", cmdType.String()),
			Args:  cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				input := readChangelog()

				changelog := parser.Parse(input)
				changelog.AddItem(cmdType, strings.Join(args, " "))

				var buf bytes.Buffer
				changelog.Render(&buf)
				output := buf.Bytes()

				writeChangelog(output)
			},
		}

		rootCmd.AddCommand(cmd)
	}
}

func init() {
	buildCommands(rootCmd)
}
