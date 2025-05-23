package undoconfig

import (
	"github.com/git-town/git-town/v20/internal/vm/program"
)

func DetermineUndoConfigProgram(initialConfigSnapshot, finalConfigSnapshot ConfigSnapshot) program.Program {
	configDiff := NewConfigDiffs(initialConfigSnapshot, finalConfigSnapshot)
	return configDiff.UndoProgram()
}
