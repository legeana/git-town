package cmd

import (
	"fmt"

	"github.com/git-town/git-town/v7/src/cli"
	"github.com/git-town/git-town/v7/src/steps"

	"github.com/spf13/cobra"
)

var undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Undoes the last run git-town command",
	Run: func(cmd *cobra.Command, args []string) {
		runState, err := steps.LoadPreviousRunState(prodRepo)
		if err != nil {
			cli.Exit(fmt.Errorf("cannot load previous run state: %v", err))
		}
		if runState == nil || runState.IsUnfinished() {
			cli.Exit(fmt.Errorf("nothing to undo"))
		}
		undoRunState := runState.CreateUndoRunState()
		err = steps.Run(&undoRunState, prodRepo, nil)
		if err != nil {
			cli.Exit(err)
		}
	},
	Args: cobra.NoArgs,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := ValidateIsRepository(prodRepo); err != nil {
			return err
		}
		return validateIsConfigured(prodRepo)
	},
}

func init() {
	RootCmd.AddCommand(undoCmd)
}
