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
	Short: "Manipulate and validate changelog files",
	Long:  `changelog manipulate and validate markdown changelog files following the keepachangelog.com specification.`,
}

var inputFilename string
var outputFilename string

func init() {
	flags := rootCmd.PersistentFlags()
	flags.StringVarP(&inputFilename, "filename", "f", "CHANGELOG.md", "Changelog file or '-' for stdin")
	rootCmd.MarkFlagFilename("filename")
	flags.StringVarP(&outputFilename, "output", "o", "-", "Output file or '-' for stdout")
	rootCmd.MarkFlagFilename("output")
}

func readChangelog() []byte {
	name := inputFilename
	if name == "-" {
		content, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
			os.Exit(2)
		}
		return content
	}

	var prefixDir string
	if strings.HasPrefix(name, "/") {
		prefixDir = ""
	} else {
		prefixDir = "./"
	}
	filename, err := filepath.Abs(prefixDir + name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(2)
	}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(2)
	}
	return content
}

func writeChangelog(content []byte) {
	if outputFilename == "-" {
		os.Stdout.Write(content)
		return
	}

	var prefixDir string
	if strings.HasPrefix(outputFilename, "/") {
		prefixDir = ""
	} else {
		prefixDir = "./"
	}

	filename, err := filepath.Abs(prefixDir + outputFilename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(2)
	}

	err = ioutil.WriteFile(filename, content, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(2)
	}
}

// Execute the program with command-line args
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
