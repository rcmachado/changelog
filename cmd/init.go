package cmd

import (
	"fmt"

	"github.com/cucumber/changelog/chg"
	"github.com/spf13/cobra"
)

func newInitCmd(iostreams *IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initializes a new changelog",
		Long: `Outputs an empty changelog, with preamble and Unreleased version

You can specify a filename using the --output/-o flag.`,
		Run: func(cmd *cobra.Command, args []string) {
			fs := cmd.Flags()
			compareURL, _ := fs.GetString("compare-url")

			c := chg.NewEmptyChangelog(compareURL)
			c.Render(iostreams.Out)

			destination, _ := fs.GetString("output")
			if destination != "-" {
				out := cmd.OutOrStdout()
				fmt.Fprintf(out, "Changelog file '%s' created.\n", destination)
			}
		},
	}

	cmd.Flags().StringP("compare-url", "c", "", "Set compare URL for Unreleased section")
	cmd.MarkFlagRequired("compare-url")

	return cmd
}
