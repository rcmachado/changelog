package cmd

import (
	"github.com/rcmachado/changelog/parser"
	"github.com/spf13/cobra"
)

func NewFmtCmd(iostreams *IOStreams) *cobra.Command {
	return &cobra.Command{
		Use:   "fmt",
		Short: "Reformat the change log file",
		Long:  "Reformats changelog input following keepachangelog.com spec",
		Run: func(cmd *cobra.Command, args []string) {
			format(ioStreams)
		},
	}
}

func format(ioStreams *IOStreams) {
	changelog := parser.Parse(ioStreams.In)
	changelog.Render(ioStreams.Out)
}
