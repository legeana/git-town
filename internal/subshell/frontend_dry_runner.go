package subshell

import (
	"fmt"

	"github.com/git-town/git-town/v20/internal/git/gitdomain"
	"github.com/git-town/git-town/v20/internal/gohacks"
	. "github.com/git-town/git-town/v20/pkg/prelude"
)

// FrontendDryRunner prints the given shell commands to the CLI as if they were executed
// but does not execute them.
type FrontendDryRunner struct {
	Backend          gitdomain.Querier
	CommandsCounter  Mutable[gohacks.Counter]
	GetCurrentBranch GetCurrentBranchFunc
	PrintBranchNames bool
	PrintCommands    bool
}

// Run runs the given command in this ShellRunner's directory.
func (self *FrontendDryRunner) Run(executable string, args ...string) error {
	var currentBranch gitdomain.LocalBranchName
	if self.PrintBranchNames {
		var err error
		currentBranch, err = self.GetCurrentBranch(self.Backend)
		if err != nil {
			return err
		}
	}
	if self.PrintCommands {
		PrintCommand(currentBranch, self.PrintBranchNames, executable, args...)
		fmt.Println("(dry run)")
	}
	return nil
}
