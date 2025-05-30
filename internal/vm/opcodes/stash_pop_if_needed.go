package opcodes

import (
	"github.com/git-town/git-town/v20/internal/vm/shared"
)

// StashPopIfNeeded restores stashed away changes into the workspace.
type StashPopIfNeeded struct {
	undeclaredOpcodeMethods `exhaustruct:"optional"`
}

func (self *StashPopIfNeeded) Run(args shared.RunArgs) error {
	stashSize, err := args.Git.StashSize(args.Backend)
	if err != nil {
		return err
	}
	if stashSize == 0 && !args.Config.Value.NormalConfig.DryRun {
		return nil
	}
	args.PrependOpcodes(
		&StashPop{},
		&ChangesUnstageAll{},
	)
	return nil
}
