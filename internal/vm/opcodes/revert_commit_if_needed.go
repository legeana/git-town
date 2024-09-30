package opcodes

import (
	"fmt"

	"github.com/git-town/git-town/v16/internal/git/gitdomain"
	"github.com/git-town/git-town/v16/internal/messages"
	"github.com/git-town/git-town/v16/internal/vm/shared"
)

// RevertCommitIfNeeded adds a commit to the current branch
// that reverts the commit with the given SHA.
type RevertCommitIfNeeded struct {
	SHA                     gitdomain.SHA
	undeclaredOpcodeMethods `exhaustruct:"optional"`
}

func (self *RevertCommitIfNeeded) Run(args shared.RunArgs) error {
	currentBranch, err := args.Git.CurrentBranch(args.Backend)
	if err != nil {
		return err
	}
	parent := args.Config.Config.Lineage.Parent(currentBranch)
	commitsInCurrentBranch, err := args.Git.CommitsInBranch(args.Backend, currentBranch, parent)
	if err != nil {
		return err
	}
	if !commitsInCurrentBranch.ContainsSHA(self.SHA) {
		return fmt.Errorf(messages.BranchDoesntContainCommit, currentBranch, self.SHA, commitsInCurrentBranch.SHAs().Join("|"))
	}
	args.PrependOpcodes(&RevertCommit{SHA: self.SHA})
	return nil
}
