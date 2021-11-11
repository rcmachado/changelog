package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var ioStreams *IOStreams

// IOStreams holds input/output streams for commands
type IOStreams struct {
	In  io.Reader
	Out io.Writer
}

var rootCmd = &cobra.Command{
	Use:   "changelog",
	Short: "Manipulate and validate changelog files",
	Long:  `changelog manipulate and validate markdown changelog files following the keepachangelog.com specification.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fs := cmd.Flags()

		fdr := openFileOrExit(fs, "filename", os.O_RDONLY, os.Stdin)
		ioStreams.In = bufio.NewReader(fdr)

		fdw := openFileOrExit(fs, "output", os.O_WRONLY|os.O_CREATE, os.Stdout)
		ioStreams.Out = bufio.NewWriter(fdw)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if ioStreams.Out != nil {
			ioStreams.Out.(*bufio.Writer).Flush()
		}
	},
}

func openFileOrExit(fs *pflag.FlagSet, option string, flag int, defaultIfDash *os.File) *os.File {
	filename, err := fs.GetString(option)
	if err != nil {
		fmt.Printf("Failed to get option '%s': %s\n", option, err)
		os.Exit(2)
	}

	if filename == "-" {
		return defaultIfDash
	}

	file, err := os.OpenFile(filename, flag, 0644)
	if err != nil {
		fmt.Printf("Failed to open file '%s': %s\n", filename, err)
		os.Exit(2)
	}
	return file
}

func init() {
	ioStreams = &IOStreams{}

	rootCmd.AddCommand(
		newInitCmd(ioStreams),
		newFmtCmd(ioStreams),
		newReleaseCmd(ioStreams),
		newShowCmd(ioStreams),
		newLatestCmd(ioStreams),
	)

	manipulationCmds := newChangeTypeCmds(ioStreams)
	for _, cmd := range manipulationCmds {
		rootCmd.AddCommand(cmd)
	}

	flags := rootCmd.PersistentFlags()
	flags.StringP("filename", "f", "CHANGELOG.md", "Changelog file or '-' for stdin")
	rootCmd.MarkFlagFilename("filename")
	flags.StringP("output", "o", "-", "Output file or '-' for stdout")
	rootCmd.MarkFlagFilename("output")
}

// Execute the program with command-line args
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
