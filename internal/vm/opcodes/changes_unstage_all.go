package opcodes

import "github.com/git-town/git-town/v17/internal/vm/shared"

type ChangesUnstageAll struct {
	undeclaredOpcodeMethods `exhaustruct:"optional"`
}

func (self *ChangesUnstageAll) Run(args shared.RunArgs) error {
	return args.Git.UnstageAll(args.Frontend)
}
