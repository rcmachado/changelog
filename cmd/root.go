package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var inputFile *bufio.Reader
var outputFile *bufio.Writer

var rootCmd = &cobra.Command{
	Use:   "changelog",
	Short: "Manipulate and validate changelog files",
	Long:  `changelog manipulate and validate markdown changelog files following the keepachangelog.com specification.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fs := cmd.Flags()

		fdr := openFileOrExit(fs, "filename", os.O_RDONLY)
		inputFile = bufio.NewReader(fdr)

		fdw := openFileOrExit(fs, "output", os.O_WRONLY|os.O_CREATE)
		outputFile = bufio.NewWriter(fdw)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if outputFile != nil {
			outputFile.Flush()
		}
	},
}

var inputFilename string
var outputFilename string

func openFileOrExit(fs *pflag.FlagSet, option string, flag int) *os.File {
	filename, err := fs.GetString(option)
	if err != nil {
		fmt.Printf("Failed to get option '%s': %s\n", option, err)
		os.Exit(2)
	}
	file, err := os.OpenFile(filename, flag, 0644)
	if err != nil {
		fmt.Printf("Failed to open file '%s': %s\n", filename, err)
		os.Exit(2)
	}
	return file
}

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
