package cmd

import (
	"github.com/rcmachado/changelog/parser"
	"github.com/spf13/cobra"
)

func newFmtCmd(iostreams *IOStreams) *cobra.Command {
	return &cobra.Command{
		Use:   "fmt",
		Short: "Reformat the change log file",
		Long:  "Reformats changelog input following keepachangelog.com spec",
		Run: func(cmd *cobra.Command, args []string) {
			changelog := parser.Parse(iostreams.In)
			changelog.Render(iostreams.Out)
		},
	}
}
