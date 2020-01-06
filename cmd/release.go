package cmd

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rcmachado/changelog/chg"
	"github.com/rcmachado/changelog/parser"
	"github.com/spf13/cobra"
)

const dateFormat = "2006-01-02"

var releaseDate string
var compareURL string

var releaseCmd = &cobra.Command{
	Use:   "release [version]",
	Short: "Change Unreleased to [version]",
	Long: `Change Unreleased section to [version], updating the compare links accordingly.
It will normalize the output with the new version.
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		release(inputFile, args, outputFile)
	},
}

func init() {
	today := time.Now().Format(dateFormat)
	releaseCmd.Flags().StringVarP(&releaseDate, "release-date", "d", today, "Release date")
	releaseCmd.Flags().StringVarP(&compareURL, "compare-url", "c", "", "Overwrite compare URL for Unreleased section")
	rootCmd.AddCommand(releaseCmd)
}

func release(input io.Reader, args []string, w io.Writer) {
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

	changelog.Render(w)
}
