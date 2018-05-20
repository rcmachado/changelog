package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "changelog",
	Short: "changelog parses and validate changelogs",
	Long: `changelog parses and validates markdown changelog files
following the keepachangelog.com specification.`,
}

var filename string

func init() {
	flags := rootCmd.PersistentFlags()
	flags.StringVarP(&filename, "filename", "f", "CHANGELOG.md", "Changelog file (use '-' for stdin)")
}

func readChangelog(name string) ([]byte, error) {
	if name == "-" {
		content, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return nil, err
		}
		return content, nil
	}

	var prefixDir string
	if strings.HasPrefix(name, "/") {
		prefixDir = ""
	} else {
		prefixDir = "./"
	}
	filename, err := filepath.Abs(prefixDir + name)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return content, nil
}

// Execute the program with command-line args
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
