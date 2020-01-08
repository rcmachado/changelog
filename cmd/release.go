package cmd

import (
	"fmt"
	"time"

	"github.com/rcmachado/changelog/chg"
	"github.com/rcmachado/changelog/parser"
	"github.com/spf13/cobra"
)

func newReleaseCmd(iostreams *IOStreams) *cobra.Command {

	const dateFormat = "2006-01-02"

	cmd := &cobra.Command{
		Use:   "release [version]",
		Short: "Change Unreleased to [version]",
		Long: `Change Unreleased section to [version], updating the compare links accordingly.
It will normalize the output with the new version.
`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fs := cmd.Flags()

			releaseDate, _ := fs.GetString("release-date")
			compareURL, _ := fs.GetString("compare-url")

			version := chg.Version{
				Name: args[0],
				Date: releaseDate,
			}

			if compareURL != "" {
				version.Link = compareURL
			}

			changelog := parser.Parse(iostreams.In)

			_, err := changelog.Release(version)
			if err != nil {
				cmd.SilenceUsage = true
				return fmt.Errorf("Failed to create release '%s': %s\n", args[0], err)
			}

			changelog.Render(iostreams.Out)
			return nil
		},
	}

	fs := cmd.Flags()

	today := time.Now().Format(dateFormat)
	fs.StringP("release-date", "d", today, "Release date")
	fs.StringP("compare-url", "c", "", "Overwrite compare URL for Unreleased section")

	return cmd
}
