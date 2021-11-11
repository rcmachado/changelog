package cmd

import (
	"encoding/json"
	"github.com/cucumber/changelog/parser"
	"github.com/spf13/cobra"
)

func newFmtCmd(iostreams *IOStreams) *cobra.Command {
	var jsonFlag bool

	command := &cobra.Command{
		Use:   "fmt",
		Short: "Reformat the change log file",
		Long:  "Reformats changelog input following keepachangelog.com spec",
		RunE: func(cmd *cobra.Command, args []string) error {
			changelog := parser.Parse(iostreams.In)
			if jsonFlag {
				enc := json.NewEncoder(iostreams.Out)
				enc.SetIndent("", "  ")
				return enc.Encode(changelog)
			} else {
				changelog.Render(iostreams.Out)
				return nil
			}
		},
	}
	command.Flags().BoolVar(&jsonFlag, "json", false, "output JSON")

	return command
}
