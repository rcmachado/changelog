package cmd

import (
	"fmt"

	"github.com/rcmachado/changelog/chg"
	"github.com/spf13/cobra"
)

func newInitCmd(iostreams *IOStreams) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initializes a new changelog",
		Long: `Outputs an empty changelog, with preamble and Unreleased version

You can specify a filename using the --output/-o flag.`,
		Run: func(cmd *cobra.Command, args []string) {
			compareURL := "https://github.com/rcmachado/changelog/compare/abcdef...HEAD"
			c := chg.NewEmptyChangelog(compareURL)
			c.Render(iostreams.Out)

			fs := cmd.Flags()
			destination, _ := fs.GetString("output")
			if destination != "-" {
				out := cmd.OutOrStdout()
				fmt.Fprintf(out, "Changelog file '%s' created.\n", destination)
			}
		},
	}
}
