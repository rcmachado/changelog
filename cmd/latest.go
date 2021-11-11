package cmd

import (
	"fmt"
	"io"

	"github.com/cucumber/changelog/parser"
	"github.com/spf13/cobra"
)

func newLatestCmd(iostreams *IOStreams) *cobra.Command {
	return &cobra.Command{
		Use:   "latest",
		Short: "Show latest released version number",
		Long:  `Show version number for the top (released) entry in the changelog`,
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			changelog := parser.Parse(iostreams.In)
			if len(changelog.Versions) == 0 {
				cmd.SilenceUsage = true
				return fmt.Errorf("There are no versions in the changelog yet")
			}
			releasedVersions := changelog.ReleasedVersions()
			if len(releasedVersions) == 0 {
				cmd.SilenceUsage = true
				return fmt.Errorf("There are no released versions in the changelog yet")
			}
			v := releasedVersions[0]
			io.WriteString(iostreams.Out, v.Name+"\n")
			return nil
		},
	}
}
