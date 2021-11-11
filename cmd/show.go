package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/cucumber/changelog/parser"
	"github.com/spf13/cobra"
)

func newShowCmd(iostreams *IOStreams) *cobra.Command {
	var jsonFlag bool

	command := &cobra.Command{
		Use:   "show [version]",
		Short: "Show changelog for [version]",
		Long:  `Show changelog section and entries for version [version]`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			version := args[0]
			changelog := parser.Parse(iostreams.In)

			v := changelog.Version(version)
			if v == nil {
				cmd.SilenceUsage = true
				return fmt.Errorf("Unknown version: '%s'\n", version)
			}
			if jsonFlag {
				enc := json.NewEncoder(iostreams.Out)
				enc.SetIndent("", "  ")
				return enc.Encode(v)
			} else {
				v.RenderChanges(iostreams.Out)
				return nil
			}
		},
	}
	command.Flags().BoolVar(&jsonFlag, "json", false, "output JSON")
	return command
}
