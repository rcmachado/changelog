package cmd

import (
	"fmt"
	"github.com/rcmachado/changelog/chg"
	"github.com/spf13/cobra"
)

func newInitCmd(iostreams *IOStreams) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "creates a new changelog file",
		Run: func(cmd *cobra.Command, args []string) {
			compareURL := "https://github.com/rcmachado/changelog/compare/abcdef...HEAD"
			c := chg.NewEmptyChangelog(compareURL)
			c.Render(iostreams.Out)

			out := cmd.OutOrStdout()
			filename, _ := cmd.Flags().GetString("output")
			fmt.Fprintf(out, "Changelog file '%s' created\n", filename)
		},
	}
}
