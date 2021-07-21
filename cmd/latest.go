package cmd

import (
	"fmt"
	"io"

	"github.com/rcmachado/changelog/chg"
	"github.com/rcmachado/changelog/parser"
	"github.com/spf13/cobra"
)

func newLatestCmd(iostreams *IOStreams) *cobra.Command {
	return &cobra.Command{
		Use:   "latest",
		Short: "Show latest version",
		Long:  `Show version number for the top (released) entry in the changelog`,
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			changelog := parser.Parse(iostreams.In)
			if len(changelog.Versions) == 0 {
				cmd.SilenceUsage = true
				return fmt.Errorf("There are no versions in the changelog yet")
			}
			releasedVersions := releasedVersions(changelog)
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

func releasedVersions(changelog *chg.Changelog) (result []chg.Version) {
	for _, version := range changelog.Versions {
		if version.Name != "Unreleased" {
			result = append(result, *version)
		}
	}
	return
}
