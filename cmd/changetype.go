package cmd

import (
	"fmt"
	"strings"

	"github.com/cucumber/changelog/chg"
	"github.com/cucumber/changelog/parser"
	"github.com/spf13/cobra"
)

func newChangeTypeCmd(iostreams *IOStreams, ct chg.ChangeType) *cobra.Command {
	sectionName := ct.String()

	return &cobra.Command{
		Use:   strings.ToLower(sectionName),
		Short: fmt.Sprintf("Add item under '%s' section", sectionName),
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			changelog := parser.Parse(iostreams.In)
			changelog.AddItem(ct, strings.Join(args, " "))
			changelog.Render(iostreams.Out)
		},
	}

}

func newChangeTypeCmds(iostreams *IOStreams) []*cobra.Command {
	cmdTypes := []chg.ChangeType{
		chg.Added, chg.Changed, chg.Deprecated, chg.Fixed, chg.Removed, chg.Security,
	}

	allCmds := make([]*cobra.Command, len(cmdTypes))

	for idx, changeType := range cmdTypes {
		cmdType := changeType
		cmd := newChangeTypeCmd(iostreams, cmdType)
		allCmds[idx] = cmd
	}

	return allCmds
}
