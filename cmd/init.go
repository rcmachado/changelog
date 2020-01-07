package cmd

import (
	"fmt"
	"github.com/rcmachado/changelog/chg"
	"github.com/spf13/cobra"
)

func newInitCmd(iostreams *IOStreams) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Creates a new changelog file",
		Long: `Creates CHANGELOG.md file in the current directory

You can specify a different filename using the --output/-o flag.`,
		Run: func(cmd *cobra.Command, args []string) {
			compareURL := "https://github.com/rcmachado/changelog/compare/abcdef...HEAD"
			c := chg.NewEmptyChangelog(compareURL)
			c.Render(iostreams.Out)

			out := cmd.OutOrStdout()
			filename, _ := cmd.Flags().GetString("output")
			fmt.Fprintf(out, "Changelog file '%s' created.\n", filename)
		},
	}
}
