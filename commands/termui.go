package commands

import (
	"github.com/MichaelMure/git-bug/termui"
	"github.com/spf13/cobra"
)

func runTermUI(cmd *cobra.Command, args []string) error {
	return termui.Run(repo)
}

var termUICmd = &cobra.Command{
	Use:   "termui",
	Short: "Launch the terminal UI",
	RunE:  runTermUI,
}

func init() {
	RootCmd.AddCommand(termUICmd)
}
