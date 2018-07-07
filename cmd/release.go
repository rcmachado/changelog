package cmd

import (
	"bytes"
	"fmt"
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
	Use:   "release",
	Short: "Bump version on changelog",
	Long:  "Change current Unreleased version into the new version",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		input := readChangelog()

		version := chg.Version{
			Name: args[0],
			Date: releaseDate,
		}

		if compareURL != "" {
			version.Link = compareURL
		}

		changelog := parser.Parse(input)
		_, err := changelog.Release(version)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to bump version to %s: %s\n", args[0], err)
			os.Exit(3)
		}

		var buf bytes.Buffer
		changelog.Render(&buf)
		output := buf.Bytes()

		writeChangelog(output)
	},
}

func init() {
	today := time.Now().Format(dateFormat)
	releaseCmd.Flags().StringVarP(&releaseDate, "release-date", "d", today, "")
	releaseCmd.Flags().StringVarP(&compareURL, "compare-url", "c", "", "Overwrite compare URL for Unreleased section")
	rootCmd.AddCommand(releaseCmd)
}
