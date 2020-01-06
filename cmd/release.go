package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/rcmachado/changelog/chg"
	"github.com/rcmachado/changelog/parser"
	"github.com/spf13/cobra"
)

const dateFormat = "2006-01-02"

var releaseCmd = &cobra.Command{
	Use:   "release [version]",
	Short: "Change Unreleased to [version]",
	Long: `Change Unreleased section to [version], updating the compare links accordingly.
It will normalize the output with the new version.
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
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

		changelog := parser.Parse(inputFile)

		_, err := changelog.Release(version)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create release '%s': %s\n", args[0], err)
			os.Exit(3)
		}

		changelog.Render(outputFile)
	},
}

func init() {
	today := time.Now().Format(dateFormat)
	fs := releaseCmd.Flags()
	fs.StringP("release-date", "d", today, "Release date")
	fs.StringP("compare-url", "c", "", "Overwrite compare URL for Unreleased section")
	rootCmd.AddCommand(releaseCmd)
}
