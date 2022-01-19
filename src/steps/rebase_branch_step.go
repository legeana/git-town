//nolint:ireturn
package steps

import (
	"github.com/git-town/git-town/v7/src/drivers"
	"github.com/git-town/git-town/v7/src/git"
)

// RebaseBranchStep rebases the current branch
// against the branch with the given name.
type RebaseBranchStep struct {
	NoOpStep
	BranchName string

	previousSha string
}

func (step *RebaseBranchStep) CreateAbortStep() Step {
	return &AbortRebaseBranchStep{}
}

func (step *RebaseBranchStep) CreateContinueStep() Step {
	return &ContinueRebaseBranchStep{}
}

func (step *RebaseBranchStep) CreateUndoStep(repo *git.ProdRepo) (Step, error) {
	return &ResetToShaStep{Hard: true, Sha: step.previousSha}, nil
}

func (step *RebaseBranchStep) Run(repo *git.ProdRepo, driver drivers.CodeHostingDriver) (err error) {
	step.previousSha, err = repo.Silent.CurrentSha()
	if err != nil {
		return err
	}
	err = repo.Logging.Rebase(step.BranchName)
	if err != nil {
		repo.Silent.CurrentBranchCache.Invalidate()
	}
	return err
}
